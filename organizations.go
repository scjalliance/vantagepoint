package vantagepoint

import (
	"context"
	"fmt"
)

// Organization represents a Vantagepoint organization record.
// Organizations are read-only in the API — only GET is supported.
//
// The Org field is a hierarchical key formatted according to the firm's
// CFGFormatOrg key conversion (e.g., "AP:00:SF:TR" where colons separate
// levels like Company, Office, Discipline).
type Organization struct {
	Org    string `json:"Org,omitempty"`
	Name   string `json:"Name,omitempty"`
	Status string `json:"Status,omitempty"`
}

// ListOrganizations retrieves a list of organizations from the Vantagepoint API.
// Pass a *Query to filter, paginate, or sort results; nil retrieves defaults.
func (c *Client) ListOrganizations(ctx context.Context, q *Query) ([]Organization, error) {
	var orgs []Organization
	if err := c.get(ctx, "organization", q, &orgs); err != nil {
		return nil, fmt.Errorf("listing organizations: %w", err)
	}
	return orgs, nil
}

// GetOrganization retrieves a single organization by its Org key. The
// Vantagepoint API returns an array even for single-resource lookups, so the
// response is decoded as a slice and the first element is returned.
func (c *Client) GetOrganization(ctx context.Context, org string) (*Organization, error) {
	var orgs []Organization
	if err := c.get(ctx, "organization/"+org, nil, &orgs); err != nil {
		return nil, fmt.Errorf("getting organization %s: %w", org, err)
	}
	if len(orgs) == 0 {
		return nil, fmt.Errorf("getting organization %s: %w", org, ErrNotFound)
	}
	return &orgs[0], nil
}
