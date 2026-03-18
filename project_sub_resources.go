package vantagepoint

import (
	"context"
	"fmt"
)

// ProjectAward represents an award record on a project.
type ProjectAward struct {
	WBS1        string `json:"WBS1,omitempty"`
	WBS2        string `json:"WBS2,omitempty"`
	WBS3        string `json:"WBS3,omitempty"`
	RecordID    string `json:"RecordID,omitempty"`
	Description string `json:"Description,omitempty"`
	Institution string `json:"Institution,omitempty"`
	AwardDate   string `json:"AwardDate,omitempty"`
	ModUser     string `json:"ModUser,omitempty"`
	CreateUser  string `json:"CreateUser,omitempty"`
	ModDate     string `json:"ModDate,omitempty"`
	CreateDate  string `json:"CreateDate,omitempty"`
}

// ProjectContract represents a contract record on a project.
type ProjectContract struct {
	WBS1                string  `json:"WBS1,omitempty"`
	ContractNumber      string  `json:"ContractNumber,omitempty"`
	ContractStatus      string  `json:"ContractStatus,omitempty"`
	FeeIncludeInd       string  `json:"FeeIncludeInd,omitempty"`
	RequestDate         string  `json:"RequestDate,omitempty"`
	ApprovedDate        string  `json:"ApprovedDate,omitempty"`
	Period              int     `json:"Period,omitempty"`
	ContractType        string  `json:"ContractType,omitempty"`
	Notes               string  `json:"Notes,omitempty"`
	ContractDescription string  `json:"ContractDescription,omitempty"`
	FeeDirLab           float64 `json:"FeeDirLab,omitempty"`
	FeeDirExp           float64 `json:"FeeDirExp,omitempty"`
	Fee                 float64 `json:"Fee,omitempty"`
	ConsultFee          float64 `json:"ConsultFee,omitempty"`
	ReimbAllowExp       float64 `json:"ReimbAllowExp,omitempty"`
	ReimbAllowCons      float64 `json:"ReimbAllowCons,omitempty"`
	ReimbAllow          float64 `json:"ReimbAllow,omitempty"`
	Total               float64 `json:"Total,omitempty"`
}

// ProjectProposal represents a proposal record on a project.
type ProjectProposal struct {
	CustomPropID    string  `json:"CustomPropID,omitempty"`
	Name            string  `json:"Name,omitempty"`
	Number          string  `json:"Number,omitempty"`
	Type            string  `json:"Type,omitempty"`
	WBS1            string  `json:"WBS1,omitempty"`
	Org             string  `json:"Org,omitempty"`
	OrgName         string  `json:"OrgName,omitempty"`
	ProposalManager string  `json:"ProposalManager,omitempty"`
	Status          string  `json:"Status,omitempty"`
	Source          string  `json:"Source,omitempty"`
	DateAdvertised  string  `json:"DateAdvertised,omitempty"`
	SubmittalDate   string  `json:"SubmittalDate,omitempty"`
	DueDate         string  `json:"DueDate,omitempty"`
	AwardDate       string  `json:"AwardDate,omitempty"`
	Fee             float64 `json:"Fee,omitempty"`
	Notes           string  `json:"Notes,omitempty"`
	CreateUser      string  `json:"CreateUser,omitempty"`
	CreateDate      string  `json:"CreateDate,omitempty"`
	ModUser         string  `json:"ModUser,omitempty"`
	ModDate         string  `json:"ModDate,omitempty"`
}

// ProjectMilestone represents a milestone record on a project.
type ProjectMilestone struct {
	RecordID    string `json:"RecordID,omitempty"`
	WBS1        string `json:"WBS1,omitempty"`
	WBS2        string `json:"WBS2,omitempty"`
	WBS3        string `json:"WBS3,omitempty"`
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
	EndDate     string `json:"EndDate,omitempty"`
	Notes       string `json:"Notes,omitempty"`
	SystemInd   string `json:"SystemInd,omitempty"`
	WBSName     string `json:"WBSName,omitempty"`
}

// ProjectCompetition represents a competition record on a project.
type ProjectCompetition struct {
	RecordID  string `json:"RecordID,omitempty"`
	ClientID  string `json:"ClientID,omitempty"`
	ClName    string `json:"clName,omitempty"`
	Strengths string `json:"Strengths,omitempty"`
	Weakness  string `json:"Weakness,omitempty"`
	Notes     string `json:"Notes,omitempty"`
	Incumbent string `json:"Incumbent,omitempty"`
	WBS1      string `json:"WBS1,omitempty"`
}

// ProjectLink represents a link associated with a project.
type ProjectLink struct {
	Description string `json:"Description,omitempty"`
	FilePath    string `json:"FilePath,omitempty"`
	Graphic     string `json:"Graphic,omitempty"`
}

// ListProjectAwards retrieves award records for a project.
// Uses GET /project/{wbsKey}/award.
func (c *Client) ListProjectAwards(ctx context.Context, wbsKey string, q *Query) ([]ProjectAward, error) {
	var results []ProjectAward
	if err := c.get(ctx, "project/"+wbsKey+"/award", q, &results); err != nil {
		return nil, fmt.Errorf("listing project awards for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectContracts retrieves contract records for a project.
// Uses GET /project/{wbsKey}/contract.
func (c *Client) ListProjectContracts(ctx context.Context, wbsKey string, q *Query) ([]ProjectContract, error) {
	var results []ProjectContract
	if err := c.get(ctx, "project/"+wbsKey+"/contract", q, &results); err != nil {
		return nil, fmt.Errorf("listing project contracts for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectContractDetails retrieves contract detail records for a project.
// Uses GET /project/{wbsKey}/contractdetail.
func (c *Client) ListProjectContractDetails(ctx context.Context, wbsKey string, q *Query) ([]ProjectContract, error) {
	var results []ProjectContract
	if err := c.get(ctx, "project/"+wbsKey+"/contractdetail", q, &results); err != nil {
		return nil, fmt.Errorf("listing project contract details for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectProposals retrieves proposal records for a project.
// Uses GET /project/{wbsKey}/proposal.
func (c *Client) ListProjectProposals(ctx context.Context, wbsKey string, q *Query) ([]ProjectProposal, error) {
	var results []ProjectProposal
	if err := c.get(ctx, "project/"+wbsKey+"/proposal", q, &results); err != nil {
		return nil, fmt.Errorf("listing project proposals for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectMilestones retrieves milestone records for a project.
// Uses GET /project/{wbsKey}/milestone.
func (c *Client) ListProjectMilestones(ctx context.Context, wbsKey string, q *Query) ([]ProjectMilestone, error) {
	var results []ProjectMilestone
	if err := c.get(ctx, "project/"+wbsKey+"/milestone", q, &results); err != nil {
		return nil, fmt.Errorf("listing project milestones for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectCompetition retrieves competition records for a project.
// Uses GET /project/{wbsKey}/competition.
func (c *Client) ListProjectCompetition(ctx context.Context, wbsKey string, q *Query) ([]ProjectCompetition, error) {
	var results []ProjectCompetition
	if err := c.get(ctx, "project/"+wbsKey+"/competition", q, &results); err != nil {
		return nil, fmt.Errorf("listing project competition for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectLinks retrieves links associated with a project.
// Uses GET /project/{wbsKey}/links.
func (c *Client) ListProjectLinks(ctx context.Context, wbsKey string, q *Query) ([]ProjectLink, error) {
	var results []ProjectLink
	if err := c.get(ctx, "project/"+wbsKey+"/links", q, &results); err != nil {
		return nil, fmt.Errorf("listing project links for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectFiles retrieves supporting documents for a project.
// Uses GET /project/{wbsKey}/files.
func (c *Client) ListProjectFiles(ctx context.Context, wbsKey string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "project/"+wbsKey+"/files", q, &results); err != nil {
		return nil, fmt.Errorf("listing project files for %s: %w", wbsKey, err)
	}
	return results, nil
}

// ListProjectCustomTable retrieves custom table data for a project.
// Uses GET /project/{wbsKey}/customTable/{customTable}.
func (c *Client) ListProjectCustomTable(ctx context.Context, wbsKey, customTable string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "project/"+wbsKey+"/customTable/"+customTable, q, &results); err != nil {
		return nil, fmt.Errorf("listing project custom table %s for %s: %w", customTable, wbsKey, err)
	}
	return results, nil
}
