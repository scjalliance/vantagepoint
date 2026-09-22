package vantagepoint

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"
)

// The OAuth /token endpoint answers with error/error_description, not the
// message/detail shape the REST endpoints use. Decoding one into the other
// succeeds and leaves every field empty, which reported a failed login as
// "vantagepoint API error 400: " with nothing after the colon.
func TestParseAPIErrorReadsOAuthShape(t *testing.T) {
	body := `{"error":"invalid_grant","error_description":"Your login has been disabled.  Please contact your system administrator."}`

	err := parseAPIError(http.StatusBadRequest, []byte(body), true)

	if !strings.Contains(err.Message, "Your login has been disabled") {
		t.Errorf("Message = %q, want the error_description sentence", err.Message)
	}
	if !strings.Contains(err.Detail, "invalid_grant") {
		t.Errorf("Detail = %q, want the machine-readable error code", err.Detail)
	}
	if !strings.Contains(err.Error(), "Your login has been disabled") {
		t.Errorf("Error() = %q, want the reason in the rendered message", err.Error())
	}
}

// Some OAuth errors carry only the code.
func TestParseAPIErrorFallsBackToErrorCode(t *testing.T) {
	err := parseAPIError(http.StatusBadRequest, []byte(`{"error":"unsupported_grant_type"}`), true)

	if err.Message != "unsupported_grant_type" {
		t.Errorf("Message = %q, want the error code", err.Message)
	}
}

// The REST shape must keep working exactly as before.
func TestParseAPIErrorReadsRESTShape(t *testing.T) {
	err := parseAPIError(http.StatusNotFound, []byte(`{"message":"Record not found","detail":"WBS1 99-999999"}`), true)

	if err.Message != "Record not found" {
		t.Errorf("Message = %q", err.Message)
	}
	if err.Detail != "WBS1 99-999999" {
		t.Errorf("Detail = %q", err.Detail)
	}
}

// A proxy or gateway answers with HTML, which is neither shape. The status
// alone is not much, but a quote of what actually arrived usually identifies
// the hop that produced it.
func TestParseAPIErrorQuotesUnrecognizedBody(t *testing.T) {
	err := parseAPIError(http.StatusBadGateway, []byte("<html><body>502 Bad Gateway</body></html>"), true)

	if !strings.Contains(err.Message, "502") {
		t.Errorf("Message = %q, want the status", err.Message)
	}
	if !strings.Contains(err.Detail, "Bad Gateway") {
		t.Errorf("Detail = %q, want a quote of the body", err.Detail)
	}
}

// An unbounded quote would paste a whole error page into a log line and an
// email.
func TestParseAPIErrorTruncatesLongBody(t *testing.T) {
	err := parseAPIError(http.StatusBadGateway, []byte(strings.Repeat("x", 4000)), true)

	if len(err.Detail) > maxErrorBodySnippet+len(snippetEllipsis) {
		t.Errorf("Detail is %d bytes, want it bounded to %d", len(err.Detail), maxErrorBodySnippet)
	}
	if !strings.HasSuffix(err.Detail, snippetEllipsis) {
		t.Errorf("Detail = %q, want a truncation marker", err.Detail)
	}
}

// Truncating at a fixed byte offset splits a multi-byte rune. The result goes
// into a log line, an HTML report, and an email, and a JSON encoder would
// silently rewrite the broken rune, so the snippet must already be well-formed.
func TestParseAPIErrorTruncatesOnRuneBoundary(t *testing.T) {
	// Lands the cut inside the first é.
	body := strings.Repeat("x", maxErrorBodySnippet-1) + strings.Repeat("é", 20)

	err := parseAPIError(http.StatusBadGateway, []byte(body), true)

	if !utf8.ValidString(err.Detail) {
		t.Errorf("Detail is not valid UTF-8: %q", err.Detail)
	}
	if !strings.HasSuffix(err.Detail, snippetEllipsis) {
		t.Errorf("Detail = %q, want a truncation marker", err.Detail)
	}
}

// A gateway can answer with something that was never UTF-8 at all.
func TestParseAPIErrorSanitizesBinaryBody(t *testing.T) {
	err := parseAPIError(http.StatusBadGateway, []byte{0xff, 0xfe, 0x00, 0xff}, true)

	if !utf8.ValidString(err.Detail) {
		t.Errorf("Detail is not valid UTF-8: %q", err.Detail)
	}
}

// The /token request posts the password and the client secret. A gateway that
// answers with a page echoing the submitted form would otherwise copy them into
// an error that reaches logs and the emailed weekly report.
func TestTokenErrorDoesNotQuoteUnrecognizedBody(t *testing.T) {
	const password = "hunter2-should-never-appear"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadRequest)
		// A WAF echoing the rejected request back at the caller.
		fmt.Fprintf(w, "<html><body>Request blocked: %s</body></html>", body)
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "db", "id", "client-secret-should-never-appear")
	err := c.Authenticate(context.Background(), "user", password)
	if err == nil {
		t.Fatal("expected an error")
	}

	if strings.Contains(err.Error(), password) {
		t.Errorf("the password reached the error: %q", err.Error())
	}
	if strings.Contains(err.Error(), "client-secret-should-never-appear") {
		t.Errorf("the client secret reached the error: %q", err.Error())
	}
	// It must still say something, or this is the blank message all over again.
	if !strings.Contains(err.Error(), "400") {
		t.Errorf("error = %q, want the status", err.Error())
	}
}

// Withholding the body applies only to the credential path. A REST endpoint
// still quotes what arrived, which is usually what identifies the bad hop.
func TestRESTErrorStillQuotesUnrecognizedBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/token") {
			w.Write([]byte(`{"access_token":"t","token_type":"bearer","expires_in":3600}`))
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("<html><body>upstream-proxy-7 refused</body></html>"))
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "db", "id", "secret")
	if err := c.Authenticate(context.Background(), "user", "pass"); err != nil {
		t.Fatalf("Authenticate: %v", err)
	}

	var out any
	err := c.RawGet(context.Background(), "/Projects", nil, &out)
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "upstream-proxy-7") {
		t.Errorf("RawGet error = %q, want the body quoted", err.Error())
	}
}

// An empty body must still produce a reason rather than a bare colon.
func TestParseAPIErrorHandlesEmptyBody(t *testing.T) {
	err := parseAPIError(http.StatusBadRequest, nil, true)

	if err.Message == "" {
		t.Fatal("Message is empty; the caller would render a blank reason")
	}
	if strings.HasSuffix(err.Error(), ": ") {
		t.Errorf("Error() = %q, want no trailing empty reason", err.Error())
	}
}

// Status-code classification must survive the rewrite, since callers branch on
// the sentinel errors rather than on the message.
func TestParseAPIErrorStillUnwraps(t *testing.T) {
	for _, tt := range []struct {
		status int
		want   error
	}{
		{http.StatusBadRequest, ErrBadRequest},
		{http.StatusUnauthorized, ErrNotAuthenticated},
		{http.StatusForbidden, ErrForbidden},
		{http.StatusNotFound, ErrNotFound},
		{http.StatusTooManyRequests, ErrRateLimited},
		{http.StatusInternalServerError, ErrServerError},
	} {
		err := parseAPIError(tt.status, []byte(`{"error":"nope"}`), true)
		if !errors.Is(err, tt.want) {
			t.Errorf("status %d does not unwrap to %v", tt.status, tt.want)
		}
	}
}

// End to end through Authenticate, which is the path that produced a blank
// failure email two weeks running.
func TestAuthenticateReportsServerReason(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid_grant","error_description":"Please update the password for the API's User record in Vantagepoint."}`))
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "db", "id", "secret")
	err := c.Authenticate(context.Background(), "user", "pass")
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "Please update the password") {
		t.Errorf("Authenticate error = %q, want the server's reason", err.Error())
	}
	if !errors.Is(err, ErrBadRequest) {
		t.Errorf("Authenticate error does not unwrap to ErrBadRequest")
	}
}

// The same applies to every other endpoint: parseErrorResponse had the same
// blank-message hole.
func TestRequestReportsServerReason(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/token") {
			w.Write([]byte(`{"access_token":"t","token_type":"bearer","expires_in":3600}`))
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"error":"insufficient_scope","error_description":"The API user lacks access to Projects."}`))
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "db", "id", "secret")
	if err := c.Authenticate(context.Background(), "user", "pass"); err != nil {
		t.Fatalf("Authenticate: %v", err)
	}

	var out any
	err := c.RawGet(context.Background(), "/Projects", nil, &out)
	if err == nil {
		t.Fatal("expected an error")
	}
	if !strings.Contains(err.Error(), "lacks access to Projects") {
		t.Errorf("RawGet error = %q, want the server's reason", err.Error())
	}
}
