package cloapi

import (
	"errors"
	"fmt"
	"net/http"
)

// Error makes the generated ApiError satisfy the error interface. The generated
// Parse*Response functions return *ApiError as the error on any non-2xx response,
// so callers get an idiomatic `resp, err := ...WithResponse(...)` flow.
func (e ApiError) Error() string {
	if e.Code != 0 {
		if e.Description != "" {
			return fmt.Sprintf("API Error [%d]: %s (%s)", e.Code, e.Message, e.Description)
		}
		return fmt.Sprintf("API Error [%d]: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("API Error: %s", e.Message)
}

// AsApiError extracts an *ApiError from err, if the chain contains one.
func AsApiError(err error) (*ApiError, bool) {
	var apiErr *ApiError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}

// IsNotFound reports whether err is an API error with HTTP status 404. This is the
// primary classifier Terraform-style consumers use to drop a resource from state.
func IsNotFound(err error) bool {
	return HasStatus(err, http.StatusNotFound)
}

// IsClientError reports whether err is an API error with a 4xx status.
func IsClientError(err error) bool {
	if apiErr, ok := AsApiError(err); ok {
		return apiErr.Code >= 400 && apiErr.Code < 500
	}
	return false
}

// IsServerError reports whether err is an API error with a 5xx status.
func IsServerError(err error) bool {
	if apiErr, ok := AsApiError(err); ok {
		return apiErr.Code >= 500 && apiErr.Code < 600
	}
	return false
}

// HasStatus reports whether err is an API error whose code equals status.
func HasStatus(err error, status int) bool {
	if apiErr, ok := AsApiError(err); ok {
		return apiErr.Code == status
	}
	return false
}
