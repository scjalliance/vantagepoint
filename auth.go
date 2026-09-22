package vantagepoint

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// tokenInfo holds the OAuth 2.0 token data returned by the Vantagepoint API.
type tokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	// ExpiresAt is the calculated absolute expiry time based on ExpiresIn.
	ExpiresAt time.Time `json:"-"`
}

// tokenExpiryMargin is subtracted from the token expiry time to ensure
// the token is refreshed before it actually expires.
const tokenExpiryMargin = 60 * time.Second

// Authenticate obtains an access token using the OAuth 2.0 resource owner
// password grant. The username and password are sent as form data to the
// /token endpoint along with the client credentials and database name.
func (c *Client) Authenticate(ctx context.Context, username, password string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	form := url.Values{
		"grant_type":    {"password"},
		"username":      {username},
		"password":      {password},
		"database":      {c.database},
		"client_id":     {c.clientID},
		"client_secret": {c.clientSecret},
	}

	tok, err := c.requestToken(ctx, form)
	if err != nil {
		return fmt.Errorf("authenticating: %w", err)
	}

	c.tokenInfo = tok
	return nil
}

// RefreshToken obtains a new access token using the stored refresh token.
// This requires a prior successful call to Authenticate.
func (c *Client) RefreshToken(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.tokenInfo == nil || c.tokenInfo.RefreshToken == "" {
		return fmt.Errorf("refreshing token: %w", ErrNotAuthenticated)
	}

	form := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {c.tokenInfo.RefreshToken},
		"database":      {c.database},
		"client_id":     {c.clientID},
		"client_secret": {c.clientSecret},
	}

	tok, err := c.requestToken(ctx, form)
	if err != nil {
		return fmt.Errorf("refreshing token: %w", err)
	}

	c.tokenInfo = tok
	return nil
}

// requestToken sends a token request to the /token endpoint and parses the response.
func (c *Client) requestToken(ctx context.Context, form url.Values) (*tokenInfo, error) {
	tokenURL := strings.TrimRight(c.baseURL, "/") + "/token"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("creating token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// The /token endpoint answers with the OAuth error shape rather than
		// the one the REST endpoints use, so this goes through the shared
		// parser. Decoding it straight into an APIError used to succeed with
		// every field empty, and a rejected login carried no reason at all.
		//
		// The credential variant: this request posted the password, the client
		// secret, and on a refresh the refresh token. An unrecognized body is
		// never quoted back, and a response that repeats any of those values is
		// withheld in full. Read from the form so every grant type is covered by
		// the values it actually sent.
		return nil, parseCredentialErrorResponse(resp,
			form.Get("password"),
			form.Get("client_secret"),
			form.Get("refresh_token"),
		)
	}

	var tok tokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tok); err != nil {
		return nil, fmt.Errorf("decoding token response: %w", err)
	}

	tok.ExpiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
	return &tok, nil
}

// token returns the current access token, refreshing it if necessary and
// auto-refresh is enabled. This method is safe for concurrent use.
func (c *Client) token(ctx context.Context) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.tokenInfo == nil {
		return "", ErrNotAuthenticated
	}

	if time.Now().After(c.tokenInfo.ExpiresAt.Add(-tokenExpiryMargin)) {
		if !c.autoRefresh || c.tokenInfo.RefreshToken == "" {
			return "", ErrTokenExpired
		}
		// Unlock for the refresh call, which acquires its own lock.
		// Instead, call requestToken directly since we already hold the lock.
		form := url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {c.tokenInfo.RefreshToken},
			"database":      {c.database},
			"client_id":     {c.clientID},
			"client_secret": {c.clientSecret},
		}
		tok, err := c.requestToken(ctx, form)
		if err != nil {
			return "", fmt.Errorf("auto-refreshing token: %w", err)
		}
		c.tokenInfo = tok
	}

	return c.tokenInfo.AccessToken, nil
}
