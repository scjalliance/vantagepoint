package vantagepoint

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
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
		// Only fill an empty detail. A body like {"detail":"account locked"}
		// has no message but does carry the reason, and overwriting it with
		// boilerplate would throw away the only useful thing in the response.
		if detail == "" {
			switch {
			case quoteBody:
				detail = bodySnippet(body)
			case len(bytes.TrimSpace(body)) > 0:
				detail = withheldBodyDetail
			}
		}
	}

	return &APIError{StatusCode: statusCode, Message: message, Detail: detail}
}

// echoedCredentialsDetail stands in for a response that gave back something the
// request had sent in confidence.
const echoedCredentialsDetail = "the response repeated credentials from the request, so it is withheld in full"

// minDetectableSecret is the shortest value containsSecret will look for on a
// request that did not itself carry the secret in its body.
//
// Substring matching cannot tell a secret from a coincidence, and a short
// value is nearly all coincidence: a one-character token matches every body
// containing that letter, which would withhold the reason from every error.
// The only secret on that path is the bearer token, which is far longer than
// this, so the floor costs nothing and stops a degenerate value from
// suppressing every message.
//
// The credential path uses noSecretLengthFloor instead. There the request
// posted the password itself, so a short one is exactly the case that needs
// catching, and a false positive costs one withheld error body against a
// leaked password.
const (
	minDetectableSecret = 8
	noSecretLengthFloor = 1
)

// containsSecret reports whether body carries any of the given secrets.
//
// Three spellings of the body are searched, for two spellings of each secret.
// A token request goes out as an encoded form, so an echo of it comes back
// encoded, but not necessarily the way Go would have written it: the hex case
// may differ, a space may be "+" or "%20", and a body can mix the two. The
// combinations are enumerated in the function rather than reasoned about.
//
// This replaced an earlier attempt that cut the secrets out of the message with
// string replacement. Substring surgery mangled innocent text that happened to
// contain a short secret, turning "Please update the password" into "Please
// update the [redacted]word". Deciding once whether the response is tainted,
// and dropping the whole thing when it is, keeps every untainted reason intact.
func containsSecret(body []byte, secrets []string, minLen int) bool {
	s := string(body)
	// Three spellings, because an echo can come back encoded any way the
	// gateway likes rather than the way the form sent it: as-is, percent-decoded
	// (which covers %20 for a space and either hex case), and percent-decoded
	// with '+' read as a space, which is what catches a body mixing the two
	// conventions, such as "secret+pass%2fword".
	//
	// The '+' variant risks calling an innocent body an echo. That costs one
	// withheld error body; missing a real echo costs a credential, so the
	// asymmetry decides it. The caller's minLen floor keeps the false-positive
	// rate negligible in practice.
	//
	// The '+' replacement happens before the unescape, not after, so a secret
	// containing a literal '+' (which a form encodes as %2B) is not mangled
	// into a space. Reversing the two would break that case.
	haystacks := [3]string{s, tolerantUnescape(s), tolerantUnescape(strings.ReplaceAll(s, "+", " "))}

	for _, secret := range secrets {
		if len(secret) < minLen {
			continue
		}
		encoded := url.QueryEscape(secret)
		for _, haystack := range haystacks {
			if strings.Contains(haystack, secret) || strings.Contains(haystack, encoded) {
				return true
			}
		}
	}
	return false
}

// tolerantUnescape decodes the %XX sequences in s and leaves everything else
// exactly as it was, including a % that begins no valid escape.
//
// url.QueryUnescape cannot be used here because it is all or nothing: one
// stray percent sign, as in "quota 100% used", makes it fail and return no
// string at all. Treating that as "nothing to check" disabled the decoded pass
// for the whole response, and a body reading "quota 100% used; rejected
// password=correct%20horse%20battery%20staple" walked the password straight
// into the error.
//
// '+' is not treated as a space here, because this function does not know
// whether it is looking at form-encoded text. Reading '+' as a space is the
// caller's job, and containsSecret does exactly that for one of its three
// haystacks. Do not fold that into this function, and do not drop it from the
// caller as redundant: an echo spelled "secret+pass%2fword" matches nothing
// without it, which is a leak this code has already had once.
func tolerantUnescape(s string) string {
	if !strings.ContainsRune(s, '%') {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); {
		if s[i] == '%' && i+2 < len(s) && isHexDigit(s[i+1]) && isHexDigit(s[i+2]) {
			b.WriteByte(unhexDigit(s[i+1])<<4 | unhexDigit(s[i+2]))
			i += 3
			continue
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}

func isHexDigit(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}

func unhexDigit(c byte) byte {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	default:
		return c - 'A' + 10
	}
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
	// Sanitized before measuring, not after. ToValidUTF8 replaces each run of
	// invalid bytes with a single three-byte U+FFFD, so a body whose bad bytes
	// are separated by good ones grows by up to three times, and capping first
	// would let the result run past the bound it was just trimmed to. A solid
	// run of invalid bytes shrinks instead, which is why a test for this has to
	// alternate.
	s = strings.ToValidUTF8(s, "�")
	if len(s) > maxErrorBodySnippet {
		// Back up to a rune boundary: cutting at a fixed byte offset would
		// split a multi-byte rune and leave the tail invalid. s is valid UTF-8
		// by now, so this walks at most three bytes.
		cut := maxErrorBodySnippet
		for cut > 0 && !utf8.RuneStart(s[cut]) {
			cut--
		}
		s = s[:cut] + snippetEllipsis
	}
	return s
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
