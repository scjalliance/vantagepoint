package vantagepoint

import "net/http"

// Option configures a Client.
type Option func(*Client)

// WithHTTPClient sets a custom http.Client for the Vantagepoint client to use.
// This is useful for configuring timeouts, transport-level settings, or
// injecting a mock client for testing.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// WithAutoRefresh enables or disables automatic token refresh when the access
// token expires. When enabled, the client will attempt to refresh the token
// using the stored refresh token before returning an authentication error.
func WithAutoRefresh(enabled bool) Option {
	return func(c *Client) {
		c.autoRefresh = enabled
	}
}
