package vantagepoint

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Client is a Vantagepoint API client. It manages authentication, token
// refresh, and provides methods for each API resource. It is safe for
// concurrent use.
type Client struct {
	baseURL      string
	database     string
	clientID     string
	clientSecret string
	httpClient   *http.Client
	autoRefresh  bool
	tokenInfo    *tokenInfo
	mu           sync.Mutex
}

// NewClient creates a new Vantagepoint API client. The baseURL should include
// the full path to the API root (e.g., "https://host/Instance/api"). The
// database, clientID, and clientSecret are used for OAuth 2.0 authentication.
func NewClient(baseURL, database, clientID, clientSecret string, opts ...Option) *Client {
	c := &Client{
		baseURL:      strings.TrimRight(baseURL, "/"),
		database:     database,
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		autoRefresh: true,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// do executes an HTTP request against the Vantagepoint API. It handles
// authentication headers, JSON encoding/decoding, query parameter application,
// and error mapping. If result is non-nil, the response body is decoded into it.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, result any) error {
	reqURL := c.baseURL + "/" + strings.TrimLeft(path, "/")

	if len(query) > 0 {
		reqURL += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	tok, err := c.token(ctx)
	if err != nil {
		return fmt.Errorf("getting auth token: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+tok)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseErrorResponse(resp, tok)
	}

	if result != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
	}

	return nil
}

// doPostAsGet executes a POST request with the dvp-postAsGet header set to
// "true". This allows sending query parameters in the request body when they
// would exceed URL length limits.
func (c *Client) doPostAsGet(ctx context.Context, path string, query url.Values, result any) error {
	reqURL := c.baseURL + "/" + strings.TrimLeft(path, "/")

	var bodyReader io.Reader
	if len(query) > 0 {
		bodyReader = strings.NewReader(query.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("creating post-as-get request: %w", err)
	}

	tok, err := c.token(ctx)
	if err != nil {
		return fmt.Errorf("getting auth token: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("dvp-postAsGet", "true")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending post-as-get request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseErrorResponse(resp, tok)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
	}

	return nil
}

// get performs a GET request. If the query is configured for POST-as-GET mode,
// it uses doPostAsGet instead.
func (c *Client) get(ctx context.Context, path string, q *Query, result any) error {
	if q != nil && q.NeedsPostAsGet() {
		return c.doPostAsGet(ctx, path, q.Values(), result)
	}
	var vals url.Values
	if q != nil {
		vals = q.Values()
	}
	return c.do(ctx, http.MethodGet, path, vals, nil, result)
}

// post performs a POST request with the given body, decoding the response into result.
func (c *Client) post(ctx context.Context, path string, body any, result any) error {
	return c.do(ctx, http.MethodPost, path, nil, body, result)
}

// put performs a PUT request with the given body, decoding the response into result.
func (c *Client) put(ctx context.Context, path string, body any, result any) error {
	return c.do(ctx, http.MethodPut, path, nil, body, result)
}

// delete performs a DELETE request against the given path.
func (c *Client) delete(ctx context.Context, path string) error {
	return c.do(ctx, http.MethodDelete, path, nil, nil, nil)
}

// RawGet performs a GET request against the given path and decodes the JSON
// response into result. This is useful for debugging or accessing endpoints
// not yet wrapped by typed methods.
func (c *Client) RawGet(ctx context.Context, path string, q *Query, result any) error {
	return c.get(ctx, path, q, result)
}

// maxErrorBodyRead bounds how much of an error response is read before it is
// parsed. An error body is a sentence or a small JSON object; anything larger
// is a misrouted page, and reading it in full would cost memory for no
// diagnostic gain.
const maxErrorBodyRead = 64 << 10

// parseErrorResponse reads an error response body and returns an appropriate
// APIError. The secrets are values the request carried that must not come back
// out in the error, such as the bearer token in the Authorization header.
func parseErrorResponse(resp *http.Response, secrets ...string) error {
	return parseErrorResponseBody(resp, true, secrets)
}

// parseCredentialErrorResponse is parseErrorResponse for a request whose body
// carried credentials.
//
// It reads both of Vantagepoint's error shapes as usual, but will not quote an
// unrecognized body. A gateway that answers a POST with a page echoing the
// submitted form would otherwise copy the password and client secret into an
// error that reaches logs and the weekly report.
func parseCredentialErrorResponse(resp *http.Response, secrets ...string) error {
	return parseErrorResponseBody(resp, false, secrets)
}

func parseErrorResponseBody(resp *http.Response, quoteBody bool, secrets []string) error {
	// Read rather than stream-decode, so parseAPIError can try both of
	// Vantagepoint's error shapes and still quote the raw body if it is
	// neither.
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxErrorBodyRead))
	if err != nil {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("request failed with status %d", resp.StatusCode),
			Detail:     fmt.Sprintf("the error body could not be read: %v", err),
		}
	}
	apiErr := parseAPIError(resp.StatusCode, body, quoteBody)

	// A short secret is worth matching exactly where the request carried it.
	// quoteBody is false on that path and only there, so it selects the floor
	// as well: an eight-character password is realistic and would otherwise
	// reach the error through error_description untouched.
	minLen := minDetectableSecret
	if !quoteBody {
		minLen = noSecretLengthFloor
	}

	// Judged on the raw body and on what came out of it. A JSON encoder is free
	// to escape characters it need not, writing & for "&", which the raw
	// bytes do not match but the decoded message does; checking only the parsed
	// fields would in turn miss an echo in a body shape the parser ignored.
	if containsSecret(body, secrets, minLen) ||
		containsSecret([]byte(apiErr.Message+"\x00"+apiErr.Detail), secrets, minLen) {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("request failed with status %d", resp.StatusCode),
			Detail:     echoedCredentialsDetail,
		}
	}
	return apiErr
}
