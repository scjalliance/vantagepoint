package vantagepoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Sentinel errors for common API error conditions.
var (
	// ErrNotAuthenticated is returned when an API call is made without valid credentials.
	ErrNotAuthenticated = errors.New("not authenticated")

	// ErrTokenExpired is returned when the access token has expired and could not be refreshed.
	ErrTokenExpired = errors.New("token expired")

	// ErrNotFound is returned when the requested resource does not exist.
	ErrNotFound = errors.New("not found")

	// ErrForbidden is returned when the authenticated user lacks permission for the request.
	ErrForbidden = errors.New("forbidden")

	// ErrBadRequest is returned when the API rejects the request due to invalid parameters.
	ErrBadRequest = errors.New("bad request")

	// ErrRateLimited is returned when the API rate limit has been exceeded.
	ErrRateLimited = errors.New("rate limited")

	// ErrServerError is returned when the API returns a 5xx status code.
	ErrServerError = errors.New("server error")
)

// APIError represents an error response from the Vantagepoint API.
type APIError struct {
	// StatusCode is the HTTP status code returned by the API.
	StatusCode int `json:"statusCode"`
	// Message is the error message from the API.
	Message string `json:"message"`
	// Detail provides additional error context if available.
	Detail string `json:"detail,omitempty"`
}

// maxErrorBodySnippet bounds how much of an unrecognized error body is quoted
// into the message, and snippetEllipsis marks a quote that was cut short. An
// unbounded quote would paste a whole HTML error page into a log line and into
// the failure emails the scheduled report sends.
const (
	maxErrorBodySnippet = 512
	snippetEllipsis     = "..."
)

// withheldBodyDetail stands in for an unrecognized body that must not be
// quoted. It says a body arrived without repeating it.
const withheldBodyDetail = "the response was not a recognized error shape; its body is withheld because the request carried credentials"

// parseAPIError builds an APIError from an error response body.
//
// Vantagepoint answers with two different error shapes. The REST endpoints
// return {"message":...,"detail":...}; the OAuth /token endpoint returns
// {"error":"invalid_grant","error_description":"..."}. Decoding the second into
// the first succeeds and leaves every field empty, which is how a rejected
// login used to report itself as "vantagepoint API error 400: " with nothing
// after the colon.
//
// A body that is neither shape, such as an HTML page from a proxy or an empty
// response, falls back to the status line so a caller is never handed a blank
// reason. quoteBody decides whether that fallback also quotes what arrived:
// true for the REST endpoints, where the quote usually identifies the hop that
// produced the page, and false for /token, whose request body carries the
// password and client secret that such a page may echo back.
func parseAPIError(statusCode int, body []byte, quoteBody bool) *APIError {
	var parsed struct {
		Message          string `json:"message"`
		Detail           string `json:"detail"`
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
	// A decode failure is not fatal: the fallback below still produces a
	// reason, and for an HTML body it is the more useful one.
	_ = json.Unmarshal(body, &parsed)

	message, detail := parsed.Message, parsed.Detail
	if message == "" {
		// OAuth shape. error_description is the sentence meant for a human;
		// error is the machine-readable code, worth keeping alongside it.
		message = parsed.ErrorDescription
		if message == "" {
			message = parsed.Error
		} else if detail == "" {
			detail = parsed.Error
		}
	}

	if message == "" {
		message = fmt.Sprintf("request failed with status %d", statusCode)
		switch {
		case quoteBody:
			if snippet := bodySnippet(body); snippet != "" {
				detail = snippet
			}
		case len(bytes.TrimSpace(body)) > 0:
			detail = withheldBodyDetail
		}
	}

	return &APIError{StatusCode: statusCode, Message: message, Detail: detail}
}

// bodySnippet returns a bounded, single-line quote of an unrecognized response
// body, or "" when there is nothing worth quoting.
//
// The result is always valid UTF-8. It reaches a log line, an HTML report, and
// an email, and a JSON encoder would silently rewrite a broken rune anyway.
func bodySnippet(body []byte) string {
	s := strings.TrimSpace(string(body))
	if s == "" {
		return ""
	}
	// Collapsed to one line: this ends up in a log line and in an email, and a
	// multi-line HTML dump would bury the surrounding context.
	s = strings.Join(strings.Fields(s), " ")
	if len(s) > maxErrorBodySnippet {
		// Back up to a rune boundary first: cutting at a fixed byte offset
		// splits a multi-byte rune and leaves the tail invalid.
		cut := maxErrorBodySnippet
		for cut > 0 && !utf8.RuneStart(s[cut]) {
			cut--
		}
		s = s[:cut] + snippetEllipsis
	}
	// A body that was never UTF-8 to begin with, such as a binary payload from
	// a confused gateway, still has to come out well-formed.
	return strings.ToValidUTF8(s, "�")
}

// Error implements the error interface for APIError.
func (e *APIError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("vantagepoint API error %d: %s (%s)", e.StatusCode, e.Message, e.Detail)
	}
	return fmt.Sprintf("vantagepoint API error %d: %s", e.StatusCode, e.Message)
}

// Unwrap returns the corresponding sentinel error for the HTTP status code.
func (e *APIError) Unwrap() error {
	switch e.StatusCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrNotAuthenticated
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 429:
		return ErrRateLimited
	default:
		if e.StatusCode >= 500 {
			return ErrServerError
		}
		return nil
	}
}
