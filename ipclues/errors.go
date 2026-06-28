package ipclues

import (
	"errors"
	"fmt"
)

// APIError represents an error response from the IPClues API.
// It implements the error interface and can be matched with errors.As.
//
//	var apiErr *ipclues.APIError
//	if errors.As(err, &apiErr) {
//	    fmt.Println(apiErr.Code, apiErr.StatusCode)
//	}
type APIError struct {
	// StatusCode is the HTTP status code returned by the server.
	StatusCode int
	// Code is the machine-readable error code (e.g. "rate_limit_exceeded").
	Code string
	// Message is a human-readable description of the error.
	Message string
}

// Error implements the error interface.
// Format: "ipclues: api error 429 (rate_limit_exceeded): too many requests"
func (e *APIError) Error() string {
	return fmt.Sprintf("ipclues: api error %d (%s): %s", e.StatusCode, e.Code, e.Message)
}

// Is reports whether this APIError matches the target error.
// Two APIErrors are considered equal when their Code fields match,
// allowing errors.Is to work correctly regardless of pointer identity
// or whether the error has been wrapped.
//
//	errors.Is(err, ipclues.ErrRateLimit) // true when Code == "rate_limit_exceeded"
func (e *APIError) Is(target error) bool {
	var t *APIError
	if !errors.As(target, &t) {
		return false
	}
	return e.Code == t.Code
}

var (
	// ErrInvalidIP is returned when the provided string is not a valid IP address.
	ErrInvalidIP = errors.New("ipclues: invalid IP address")

	// ErrUnauthorized is returned when the API key is missing or invalid (HTTP 401).
	ErrUnauthorized = &APIError{StatusCode: 401, Code: "unauthorized", Message: "invalid or missing API key"}

	// ErrRateLimit is returned when the request rate limit has been exceeded (HTTP 429).
	ErrRateLimit = &APIError{StatusCode: 429, Code: "rate_limit_exceeded", Message: "rate limit exceeded"}

	// ErrNotFound is returned when the IP address has no data record (HTTP 404).
	ErrNotFound = &APIError{StatusCode: 404, Code: "not_found", Message: "no data found for this IP address"}

	// ErrServerError is returned on 5xx responses from the IPClues API.
	ErrServerError = &APIError{StatusCode: 500, Code: "server_error", Message: "internal server error"}
)

// NewAPIError constructs an APIError from a server response.
// Use this in your HTTP transport layer rather than constructing
// APIError literals directly, so the Message is always populated.
func NewAPIError(statusCode int, code, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}
