package vantagepoint

import (
	"context"
	"fmt"
)

// ContactLink represents a link associated with a contact.
type ContactLink struct {
	Description string `json:"Description,omitempty"`
	FilePath    string `json:"FilePath,omitempty"`
	Graphic     string `json:"Graphic,omitempty"`
}

// ContactCampaign represents a marketing campaign association for a contact.
type ContactCampaign struct {
	ContactID     string `json:"ContactID,omitempty"`
	Name          string `json:"Name,omitempty"`
	CampaignMgrID string `json:"CampaignMgrID,omitempty"`
	CampaignMgr   string `json:"CampaignMgr,omitempty"`
	Type          string `json:"Type,omitempty"`
	Stage         string `json:"Stage,omitempty"`
	Status        string `json:"Status,omitempty"`
	CampaignID    string `json:"CampaignID,omitempty"`
	CreateUser    string `json:"CreateUser,omitempty"`
	CreateDate    string `json:"CreateDate,omitempty"`
	ModUser       string `json:"ModUser,omitempty"`
	ModDate       string `json:"ModDate,omitempty"`
	LaunchDate    string `json:"LaunchDate,omitempty"`
}

// ContactCategory represents a category assignment for a contact.
type ContactCategory struct {
	ContactID    string `json:"ContactID,omitempty"`
	Category     string `json:"Category,omitempty"`
	CategoryDesc string `json:"CategoryDesc,omitempty"`
	Info         string `json:"Info,omitempty"`
	CreateUser   string `json:"CreateUser,omitempty"`
	CreateDate   string `json:"CreateDate,omitempty"`
	ModUser      string `json:"ModUser,omitempty"`
	ModDate      string `json:"ModDate,omitempty"`
}

// ListContactLinks retrieves links associated with a contact.
// Uses GET /contact/{contactID}/links.
func (c *Client) ListContactLinks(ctx context.Context, contactID string, q *Query) ([]ContactLink, error) {
	var results []ContactLink
	if err := c.get(ctx, "contact/"+contactID+"/links", q, &results); err != nil {
		return nil, fmt.Errorf("listing contact links for %s: %w", contactID, err)
	}
	return results, nil
}

// ListContactFiles retrieves supporting documents for a contact.
// Uses GET /contact/{contactID}/files.
func (c *Client) ListContactFiles(ctx context.Context, contactID string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "contact/"+contactID+"/files", q, &results); err != nil {
		return nil, fmt.Errorf("listing contact files for %s: %w", contactID, err)
	}
	return results, nil
}

// ListContactCampaigns retrieves marketing campaigns associated with a contact.
// Uses GET /contact/{contactID}/campaign.
func (c *Client) ListContactCampaigns(ctx context.Context, contactID string, q *Query) ([]ContactCampaign, error) {
	var results []ContactCampaign
	if err := c.get(ctx, "contact/"+contactID+"/campaign", q, &results); err != nil {
		return nil, fmt.Errorf("listing contact campaigns for %s: %w", contactID, err)
	}
	return results, nil
}

// ListContactCategories retrieves category assignments for a contact.
// Uses GET /contact/{contactID}/category.
func (c *Client) ListContactCategories(ctx context.Context, contactID string, q *Query) ([]ContactCategory, error) {
	var results []ContactCategory
	if err := c.get(ctx, "contact/"+contactID+"/category", q, &results); err != nil {
		return nil, fmt.Errorf("listing contact categories for %s: %w", contactID, err)
	}
	return results, nil
}

// ListContactCustomTable retrieves custom table data for a contact.
// Uses GET /contact/{contactID}/customTable/{customTable}.
func (c *Client) ListContactCustomTable(ctx context.Context, contactID, customTable string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "contact/"+contactID+"/customTable/"+customTable, q, &results); err != nil {
		return nil, fmt.Errorf("listing contact custom table %s for %s: %w", customTable, contactID, err)
	}
	return results, nil
}
