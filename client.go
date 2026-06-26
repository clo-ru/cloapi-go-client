package cloapi

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

const defaultBaseURL = "https://api.clo.ru"

// New builds a ClientWithResponses for the CLO API. It wires bearer auth, a logging
// transport, and an optional retry transport, then returns the generated high-level
// client. All ergonomics live here at the transport layer, so they apply to every
// endpoint for free.
func New(token string, opts ...Option) (*ClientWithResponses, error) {
	cfg := &clientConfig{
		baseURL: defaultBaseURL,
		timeout: 30 * time.Second,
		logger:  slog.Default(),
	}
	for _, opt := range opts {
		opt(cfg)
	}

	httpClient := cfg.httpClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: cfg.timeout}
	}

	// Transport chain (outermost first): retry -> logging -> base.
	base := httpClient.Transport
	if base == nil {
		base = http.DefaultTransport
	}
	var transport http.RoundTripper = &LoggingRoundTripper{Proxied: base, Logger: cfg.logger}
	if cfg.retry.Count > 0 {
		transport = &retryRoundTripper{Proxied: transport, Config: cfg.retry}
	}
	httpClient.Transport = transport

	return NewClientWithResponses(
		cfg.baseURL,
		WithHTTPClientDoer(httpClient),
		WithRequestEditorFn(withAuthToken(token)),
	)
}

func withAuthToken(token string) RequestEditorFn {
	return func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
}

type clientConfig struct {
	baseURL    string
	timeout    time.Duration
	logger     *slog.Logger
	httpClient *http.Client
	retry      RetryConfig
}

// Option configures the client. Functional options keep the constructor stable as
// new knobs are added.
type Option func(*clientConfig)

// WithBaseURL overrides the API base URL (default https://api.clo.ru).
func WithBaseURL(u string) Option {
	return func(c *clientConfig) { c.baseURL = u }
}

// WithTimeout sets the HTTP client timeout (ignored if WithHTTPClient is used).
func WithTimeout(t time.Duration) Option {
	return func(c *clientConfig) { c.timeout = t }
}

// WithLogger sets the slog.Logger used by the logging transport.
func WithLogger(l *slog.Logger) Option {
	return func(c *clientConfig) { c.logger = l }
}

// WithHTTPClient supplies a custom *http.Client. Its Transport is wrapped with the
// logging (and optional retry) RoundTrippers.
func WithHTTPClient(client *http.Client) Option {
	return func(c *clientConfig) { c.httpClient = client }
}

// WithRetry enables retrying transient failures (5xx + 429) for idempotent methods,
// count additional attempts spaced by backoff.
func WithRetry(count int, backoff time.Duration) Option {
	return func(c *clientConfig) { c.retry = RetryConfig{Count: count, Backoff: backoff} }
}
