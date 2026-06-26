package cloapi

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestApiErrorError(t *testing.T) {
	tests := []struct {
		name string
		err  ApiError
		want string
	}{
		{"code+desc", ApiError{Code: 404, Message: "not found", Description: "no such server"}, "API Error [404]: not found (no such server)"},
		{"code only", ApiError{Code: 500, Message: "boom"}, "API Error [500]: boom"},
		{"no code", ApiError{Message: "weird"}, "API Error: weird"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestClassification(t *testing.T) {
	notFound := &ApiError{Code: http.StatusNotFound, Message: "nope"}
	badReq := &ApiError{Code: http.StatusBadRequest}
	serverErr := &ApiError{Code: http.StatusBadGateway}
	wrapped := fmt.Errorf("call failed: %w", notFound)
	plain := errors.New("not an api error")

	tests := []struct {
		name                           string
		err                            error
		isNotFound, isClient, isServer bool
	}{
		{"404", notFound, true, true, false},
		{"wrapped 404", wrapped, true, true, false},
		{"400", badReq, false, true, false},
		{"502", serverErr, false, false, true},
		{"plain error", plain, false, false, false},
		{"nil", nil, false, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotFound(tt.err); got != tt.isNotFound {
				t.Errorf("IsNotFound = %v, want %v", got, tt.isNotFound)
			}
			if got := IsClientError(tt.err); got != tt.isClient {
				t.Errorf("IsClientError = %v, want %v", got, tt.isClient)
			}
			if got := IsServerError(tt.err); got != tt.isServer {
				t.Errorf("IsServerError = %v, want %v", got, tt.isServer)
			}
		})
	}
}

func TestAsApiErrorAndHasStatus(t *testing.T) {
	src := &ApiError{Code: 418, Message: "teapot"}
	wrapped := fmt.Errorf("layer: %w", src)

	got, ok := AsApiError(wrapped)
	if !ok || got.Code != 418 {
		t.Fatalf("AsApiError = %+v, %v", got, ok)
	}
	if !HasStatus(wrapped, 418) {
		t.Errorf("HasStatus(418) = false, want true")
	}
	if HasStatus(wrapped, 404) {
		t.Errorf("HasStatus(404) = true, want false")
	}
	if _, ok := AsApiError(errors.New("plain")); ok {
		t.Errorf("AsApiError(plain) ok = true, want false")
	}
}
