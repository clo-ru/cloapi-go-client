package cloapi

import (
	"log/slog"
	"net/http"
	"time"
)

// LoggingRoundTripper logs every HTTP request/response via slog. Written once,
// applies to every endpoint.
type LoggingRoundTripper struct {
	Proxied http.RoundTripper
	Logger  *slog.Logger
}

func (l *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := l.Proxied.RoundTrip(req)
	duration := time.Since(start)

	attrs := []slog.Attr{
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()),
		slog.Duration("duration", duration),
	}

	if err != nil {
		attrs = append(attrs, slog.String("error", err.Error()))
		l.Logger.LogAttrs(req.Context(), slog.LevelError, "HTTP request failed", attrs...)
		return nil, err
	}

	attrs = append(attrs, slog.Int("status", resp.StatusCode))

	level := slog.LevelInfo
	switch {
	case resp.StatusCode >= 500:
		level = slog.LevelError
	case resp.StatusCode >= 400:
		level = slog.LevelWarn
	}
	l.Logger.LogAttrs(req.Context(), level, "HTTP request processed", attrs...)

	return resp, nil
}

// RetryConfig controls the retry RoundTripper. Cross-cutting: configured once on
// the client, applies to every endpoint.
type RetryConfig struct {
	// Count is the number of additional attempts after the first (0 disables retry).
	Count int
	// Backoff is the fixed delay between attempts.
	Backoff time.Duration
}

// retryRoundTripper retries transient failures (5xx + 429) for idempotent methods.
// Non-idempotent writes (POST/PATCH) are not retried by default, since the server
// may have already applied them.
type retryRoundTripper struct {
	Proxied http.RoundTripper
	Config  RetryConfig
}

func isIdempotent(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPut, http.MethodDelete, http.MethodOptions:
		return true
	default:
		return false
	}
}

func isRetryableStatus(code int) bool {
	return code == http.StatusTooManyRequests || code >= 500
}

func (rt *retryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.Config.Count <= 0 || !isIdempotent(req.Method) {
		return rt.Proxied.RoundTrip(req)
	}

	var resp *http.Response
	var err error
	for attempt := 0; attempt <= rt.Config.Count; attempt++ {
		if attempt > 0 {
			// Rewind the body for replay; bail out if it can't be rewound.
			if req.Body != nil {
				if req.GetBody == nil {
					return resp, err
				}
				body, gerr := req.GetBody()
				if gerr != nil {
					return resp, err
				}
				req.Body = body
			}
			if rt.Config.Backoff > 0 {
				select {
				case <-time.After(rt.Config.Backoff):
				case <-req.Context().Done():
					return nil, req.Context().Err()
				}
			}
		}

		resp, err = rt.Proxied.RoundTrip(req)
		if err != nil {
			continue // transport error — retry
		}
		if !isRetryableStatus(resp.StatusCode) {
			return resp, nil
		}
		// Retryable status: drain+close before the next attempt (unless it's the last).
		if attempt < rt.Config.Count {
			drainAndClose(resp)
		}
	}
	return resp, err
}

func drainAndClose(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	_ = resp.Body.Close()
}
