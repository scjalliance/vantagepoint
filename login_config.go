package vantagepoint

import (
	"context"
	"fmt"
)

// LoginLanguage represents a supported language for the login UI.
type LoginLanguage struct {
	Culture string `json:"Culture,omitempty"`
	Caption string `json:"Caption,omitempty"`
}

// LoginConfig represents the login configuration for a Vantagepoint instance,
// including supported languages and authentication settings.
type LoginConfig struct {
	Database       string          `json:"database,omitempty"`
	WinAuthDefault string          `json:"winAuthDefault,omitempty"`
	EmailConfig    bool            `json:"emailConfig,omitempty"`
	WAADEnabled    bool            `json:"WAADEnabled,omitempty"`
	Languages      []LoginLanguage `json:"Languages,omitempty"`
	WAADSettings   *WAADSettings   `json:"WAADSettings,omitempty"`
}

// WAADSettings represents Windows Azure Active Directory settings for SSO.
type WAADSettings struct {
	WAADTenant         string `json:"WAADTenant,omitempty"`
	WAADServerClientID string `json:"WAADServerClientID,omitempty"`
}

// GetLoginConfig retrieves the general login configuration including
// supported languages. Uses GET /LoginConfig.
func (c *Client) GetLoginConfig(ctx context.Context) (*LoginConfig, error) {
	var result LoginConfig
	if err := c.get(ctx, "LoginConfig", nil, &result); err != nil {
		return nil, fmt.Errorf("getting login config: %w", err)
	}
	return &result, nil
}

// GetLoginConfigForDatabase retrieves login configuration for a specific
// database, including authentication method and WAAD settings.
// Uses GET /LoginConfig/{database}.
func (c *Client) GetLoginConfigForDatabase(ctx context.Context, database string) (*LoginConfig, error) {
	var result LoginConfig
	if err := c.get(ctx, "LoginConfig/"+database, nil, &result); err != nil {
		return nil, fmt.Errorf("getting login config for %s: %w", database, err)
	}
	return &result, nil
}
