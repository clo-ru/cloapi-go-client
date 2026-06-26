package cloapi

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func newRetryClient(count int, backoff time.Duration) *retryRoundTripper {
	return &retryRoundTripper{Proxied: http.DefaultTransport, Config: RetryConfig{Count: count, Backoff: backoff}}
}

func TestRetryRetriesThenSucceeds(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&hits, 1) < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	rt := newRetryClient(3, 0)
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&hits); got != 3 {
		t.Errorf("hits = %d, want 3", got)
	}
}

func TestRetryDoesNotRetryPOST(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	rt := newRetryClient(3, 0)
	req, _ := http.NewRequest(http.MethodPost, srv.URL, strings.NewReader("{}"))
	resp, _ := rt.RoundTrip(req)
	if resp != nil {
		_ = resp.Body.Close()
	}
	if got := atomic.LoadInt32(&hits); got != 1 {
		t.Errorf("POST hits = %d, want 1 (non-idempotent, no retry)", got)
	}
}

func TestRetryOn429(t *testing.T) {
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&hits, 1) < 2 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	rt := newRetryClient(3, 0)
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	resp, err := rt.RoundTrip(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("resp=%v err=%v", resp, err)
	}
	if got := atomic.LoadInt32(&hits); got != 2 {
		t.Errorf("hits = %d, want 2", got)
	}
}

func TestRetryRewindsBody(t *testing.T) {
	var lastBody string
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		lastBody = string(b)
		if atomic.AddInt32(&hits, 1) < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	rt := newRetryClient(3, 0)
	// PUT is idempotent, so it is retried; the body must be replayed intact.
	req, _ := http.NewRequest(http.MethodPut, srv.URL, strings.NewReader(`{"k":"v"}`))
	resp, err := rt.RoundTrip(req)
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
	if lastBody != `{"k":"v"}` {
		t.Errorf("replayed body = %q, want intact", lastBody)
	}
}

func TestRetryRespectsContextCancel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	rt := newRetryClient(5, 100*time.Millisecond) // backoff outlasts ctx
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	if _, err := rt.RoundTrip(req); err == nil {
		t.Fatal("expected context error, got nil")
	}
}

// End-to-end: New() wires auth + base URL, and a non-2xx is returned as *ApiError
// so IsNotFound works against a real generated call.
func TestNewClientErrorIsNotFound(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"code":404,"message":"not found","description":"no balance"}`))
	}))
	defer srv.Close()

	cli, err := New("tok123", WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.AccountBalanceWithResponse(context.Background())
	if err == nil {
		t.Fatal("expected error on 404")
	}
	if !IsNotFound(err) {
		t.Errorf("IsNotFound = false, want true (err=%v)", err)
	}
	if resp.Error == nil || resp.Error.Code != 404 {
		t.Errorf("resp.Error = %+v, want code 404", resp.Error)
	}
	if gotAuth != "Bearer tok123" {
		t.Errorf("auth header = %q, want Bearer tok123", gotAuth)
	}
}

func TestNewClientSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	cli, err := New("tok", WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := cli.AccountBalanceWithResponse(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode())
	}
}
