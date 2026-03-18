package vantagepoint

import (
	"context"
	"fmt"
)

// Campaign represents a marketing campaign record in Vantagepoint CRM.
type Campaign struct {
	CampaignID                     string  `json:"CampaignID,omitempty"`
	Name                           string  `json:"Name,omitempty"`
	Description                    string  `json:"Description,omitempty"`
	Number                         string  `json:"Number,omitempty"`
	WBS1                           string  `json:"WBS1,omitempty"`
	WBS1Name                       string  `json:"WBS1Name,omitempty"`
	PotentialResponses             int     `json:"PotentialResponses,omitempty"`
	CustomCurrencyCode             string  `json:"CustomCurrencyCode,omitempty"`
	Org                            string  `json:"Org,omitempty"`
	OrgName                        string  `json:"OrgName,omitempty"`
	CampaignMgr                    string  `json:"CampaignMgr,omitempty"`
	EmCampaignMgrFL                string  `json:"emCampaignMgrFL,omitempty"`
	EmCampaignMgrPhone             string  `json:"emCampaignMgrPhone,omitempty"`
	EmCampaignMgrEmail             string  `json:"emCampaignMgrEmail,omitempty"`
	EmCampaignMgrTitle             string  `json:"emCampaignMgrTitle,omitempty"`
	EmCampaignMgrLocation          string  `json:"emCampaignMgrLocation,omitempty"`
	MktgMgr                        string  `json:"MktgMgr,omitempty"`
	EmMarketingMgrFL               string  `json:"emMarketingMgrFL,omitempty"`
	EmMarketingMgrPhone            string  `json:"emMarketingMgrPhone,omitempty"`
	EmMarketingMgrEmail            string  `json:"emMarketingMgrEmail,omitempty"`
	EmMarketingMgrTitle            string  `json:"emMarketingMgrTitle,omitempty"`
	EmMarketingMgrLocation         string  `json:"emMarketingMgrLocation,omitempty"`
	Manager3                       string  `json:"Manager3,omitempty"`
	EmMarketingCoordinatorFL       string  `json:"emMarketingCoordinatorFL,omitempty"`
	EmMarketingCoordinatorPhone    string  `json:"emMarketingCoordinatorPhone,omitempty"`
	EmMarketingCoordinatorEmail    string  `json:"emMarketingCoordinatorEmail,omitempty"`
	EmMarketingCoordinatorTitle    string  `json:"emMarketingCoordinatorTitle,omitempty"`
	EmMarketingCoordinatorLocation string  `json:"emMarketingCoordinatorLocation,omitempty"`
	FirstAction                    string  `json:"FirstAction,omitempty"`
	CurrentAction                  string  `json:"CurrentAction,omitempty"`
	NextAction                     string  `json:"NextAction,omitempty"`
	HasPhoto                       int     `json:"HasPhoto,omitempty"`
	PhotoModDate                   string  `json:"PhotoModDate,omitempty"`
	LaunchDate                     string  `json:"LaunchDate,omitempty"`
	EndDate                        string  `json:"EndDate,omitempty"`
	RecordStatus                   string  `json:"RecordStatus,omitempty"`
	Status                         string  `json:"Status,omitempty"`
	Type                           string  `json:"Type,omitempty"`
	Audience                       string  `json:"Audience,omitempty"`
	Objective                      string  `json:"Objective,omitempty"`
	Revenue                        float64 `json:"Revenue,omitempty"`
	Budget                         float64 `json:"Budget,omitempty"`
	ActualCost                     float64 `json:"ActualCost,omitempty"`
	Responses                      int     `json:"Responses,omitempty"`
	PctResponses                   float64 `json:"PctResponses,omitempty"`
	Projects                       int     `json:"Projects,omitempty"`
	QualifiedContacts              int     `json:"QualifiedContacts,omitempty"`
	QualifiedClients               int     `json:"QualifiedClients,omitempty"`
	CreateUser                     string  `json:"CreateUser,omitempty"`
	CreateDate                     string  `json:"CreateDate,omitempty"`
	ModUser                        string  `json:"ModUser,omitempty"`
	ModDate                        string  `json:"ModDate,omitempty"`
	ExchangeRateDate               string  `json:"ExchangeRateDate,omitempty"`
	ActualRevenue                  float64 `json:"ActualRevenue,omitempty"`
	PotentialRevenue               float64 `json:"PotentialRevenue,omitempty"`
}

// CampaignProject represents a project linked to a marketing campaign.
type CampaignProject struct {
	CampaignID string `json:"CampaignID,omitempty"`
	WBS1       string `json:"WBS1,omitempty"`
	Name       string `json:"Name,omitempty"`
	CreateUser string `json:"CreateUser,omitempty"`
	CreateDate string `json:"CreateDate,omitempty"`
	ModUser    string `json:"ModUser,omitempty"`
	ModDate    string `json:"ModDate,omitempty"`
}

// CampaignResponse represents a response to a marketing campaign.
type CampaignResponse struct {
	CampaignID string `json:"CampaignID,omitempty"`
	ContactID  string `json:"ContactID,omitempty"`
	Name       string `json:"Name,omitempty"`
	CreateUser string `json:"CreateUser,omitempty"`
	CreateDate string `json:"CreateDate,omitempty"`
	ModUser    string `json:"ModUser,omitempty"`
	ModDate    string `json:"ModDate,omitempty"`
}

// ListCampaigns retrieves a list of marketing campaigns.
// Uses GET /campaign.
func (c *Client) ListCampaigns(ctx context.Context, q *Query) ([]Campaign, error) {
	var results []Campaign
	if err := c.get(ctx, "campaign", q, &results); err != nil {
		return nil, fmt.Errorf("listing campaigns: %w", err)
	}
	return results, nil
}

// GetCampaign retrieves a single marketing campaign by its CampaignID.
// Uses GET /campaign/{campaignID}.
func (c *Client) GetCampaign(ctx context.Context, campaignID string) (*Campaign, error) {
	var results []Campaign
	if err := c.get(ctx, "campaign/"+campaignID, nil, &results); err != nil {
		return nil, fmt.Errorf("getting campaign %s: %w", campaignID, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("getting campaign %s: %w", campaignID, ErrNotFound)
	}
	return &results[0], nil
}

// CreateCampaign creates a new marketing campaign.
// Uses POST /campaign.
func (c *Client) CreateCampaign(ctx context.Context, campaign *Campaign) (*Campaign, error) {
	var results []Campaign
	if err := c.post(ctx, "campaign", campaign, &results); err != nil {
		return nil, fmt.Errorf("creating campaign: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("creating campaign: %w", ErrNotFound)
	}
	return &results[0], nil
}

// UpdateCampaign updates an existing marketing campaign.
// Uses PUT /campaign/{campaignID}.
func (c *Client) UpdateCampaign(ctx context.Context, campaignID string, campaign *Campaign) (*Campaign, error) {
	var results []Campaign
	if err := c.put(ctx, "campaign/"+campaignID, campaign, &results); err != nil {
		return nil, fmt.Errorf("updating campaign %s: %w", campaignID, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("updating campaign %s: %w", campaignID, ErrNotFound)
	}
	return &results[0], nil
}

// DeleteCampaign deletes a marketing campaign by its CampaignID.
// Uses DELETE /campaign/{campaignID}.
func (c *Client) DeleteCampaign(ctx context.Context, campaignID string) error {
	if err := c.delete(ctx, "campaign/"+campaignID); err != nil {
		return fmt.Errorf("deleting campaign %s: %w", campaignID, err)
	}
	return nil
}

// ListCampaignProjects retrieves projects linked to a campaign.
// Uses GET /campaign/{campaignID}/project.
func (c *Client) ListCampaignProjects(ctx context.Context, campaignID string, q *Query) ([]CampaignProject, error) {
	var results []CampaignProject
	if err := c.get(ctx, "campaign/"+campaignID+"/project", q, &results); err != nil {
		return nil, fmt.Errorf("listing campaign projects for %s: %w", campaignID, err)
	}
	return results, nil
}

// ListCampaignResponses retrieves responses to a campaign.
// Uses GET /campaign/{campaignID}/responses.
func (c *Client) ListCampaignResponses(ctx context.Context, campaignID string, q *Query) ([]CampaignResponse, error) {
	var results []CampaignResponse
	if err := c.get(ctx, "campaign/"+campaignID+"/responses", q, &results); err != nil {
		return nil, fmt.Errorf("listing campaign responses for %s: %w", campaignID, err)
	}
	return results, nil
}

// CampaignLink represents a link associated with a campaign.
type CampaignLink struct {
	Description string `json:"Description,omitempty"`
	FilePath    string `json:"FilePath,omitempty"`
	Graphic     string `json:"Graphic,omitempty"`
}

// ListCampaignLinks retrieves links associated with a campaign.
// Uses GET /campaign/{campaignID}/links.
func (c *Client) ListCampaignLinks(ctx context.Context, campaignID string, q *Query) ([]CampaignLink, error) {
	var results []CampaignLink
	if err := c.get(ctx, "campaign/"+campaignID+"/links", q, &results); err != nil {
		return nil, fmt.Errorf("listing campaign links for %s: %w", campaignID, err)
	}
	return results, nil
}

// ListCampaignFiles retrieves supporting documents for a campaign.
// Uses GET /campaign/{campaignID}/files.
func (c *Client) ListCampaignFiles(ctx context.Context, campaignID string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "campaign/"+campaignID+"/files", q, &results); err != nil {
		return nil, fmt.Errorf("listing campaign files for %s: %w", campaignID, err)
	}
	return results, nil
}

// ListCampaignCustomTable retrieves custom table data for a campaign.
// Uses GET /campaign/{campaignID}/customTable/{customTable}.
func (c *Client) ListCampaignCustomTable(ctx context.Context, campaignID, customTable string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "campaign/"+campaignID+"/customTable/"+customTable, q, &results); err != nil {
		return nil, fmt.Errorf("listing campaign custom table %s for %s: %w", customTable, campaignID, err)
	}
	return results, nil
}
