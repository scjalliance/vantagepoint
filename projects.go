package vantagepoint

import (
	"context"
	"fmt"
)

// Project represents a Vantagepoint project, phase, or task record.
// The API uses a single unified structure for all WBS levels (WBS1, WBS2, WBS3).
// Which fields are populated depends on the SubLevel of the record.
type Project struct {
	WBS1                      string  `json:"WBS1,omitempty"`
	WBS2                      string  `json:"WBS2,omitempty"`
	WBS3                      string  `json:"WBS3,omitempty"`
	WBSNumber                 string  `json:"WBSNumber,omitempty"`
	Name                      string  `json:"Name,omitempty"`
	LongName                  string  `json:"LongName,omitempty"`
	ChargeType                string  `json:"ChargeType,omitempty"`
	SubLevel                  string  `json:"SubLevel,omitempty"`
	Principal                 string  `json:"Principal,omitempty"`
	ProjMgr                   string  `json:"ProjMgr,omitempty"`
	Supervisor                string  `json:"Supervisor,omitempty"`
	ClientID                  string  `json:"ClientID,omitempty"`
	CLAddress                 string  `json:"CLAddress,omitempty"`
	Fee                       float64 `json:"Fee,omitempty"`
	ReimbAllow                float64 `json:"ReimbAllow,omitempty"`
	ConsultFee                float64 `json:"ConsultFee,omitempty"`
	BudOHRate                 float64 `json:"BudOHRate,omitempty"`
	Status                    string  `json:"Status,omitempty"`
	RevType                   string  `json:"RevType,omitempty"`
	MultAmt                   float64 `json:"MultAmt,omitempty"`
	Org                       string  `json:"Org,omitempty"`
	UnitTable                 string  `json:"UnitTable,omitempty"`
	StartDate                 string  `json:"StartDate,omitempty"`
	EndDate                   string  `json:"EndDate,omitempty"`
	PctComp                   float64 `json:"PctComp,omitempty"`
	LabPctComp                float64 `json:"LabPctComp,omitempty"`
	ExpPctComp                float64 `json:"ExpPctComp,omitempty"`
	BillByDefault             string  `json:"BillByDefault,omitempty"`
	BillableWarning           string  `json:"BillableWarning,omitempty"`
	Memo                      string  `json:"Memo,omitempty"`
	BudgetedFlag              string  `json:"BudgetedFlag,omitempty"`
	BudgetedLevels            string  `json:"BudgetedLevels,omitempty"`
	BillWBS1                  string  `json:"BillWBS1,omitempty"`
	BillWBS2                  string  `json:"BillWBS2,omitempty"`
	BillWBS3                  string  `json:"BillWBS3,omitempty"`
	XCharge                   string  `json:"XCharge,omitempty"`
	XChargeMethod             int     `json:"XChargeMethod,omitempty"`
	XChargeMult               float64 `json:"XChargeMult,omitempty"`
	Description               string  `json:"Description,omitempty"`
	Closed                    int     `json:"Closed,omitempty"`
	ReadOnly                  int     `json:"ReadOnly,omitempty"`
	DefaultEffortDriven       int     `json:"DefaultEffortDriven,omitempty"`
	DefaultTaskType           int     `json:"DefaultTaskType,omitempty"`
	VersionID                 int     `json:"VersionID,omitempty"`
	ContactID                 string  `json:"ContactID,omitempty"`
	CLBillingAddr             string  `json:"CLBillingAddr,omitempty"`
	Address1                  string  `json:"Address1,omitempty"`
	Address2                  string  `json:"Address2,omitempty"`
	Address3                  string  `json:"Address3,omitempty"`
	City                      string  `json:"City,omitempty"`
	State                     string  `json:"State,omitempty"`
	Zip                       string  `json:"Zip,omitempty"`
	County                    string  `json:"County,omitempty"`
	Country                   string  `json:"Country,omitempty"`
	FederalInd                string  `json:"FederalInd,omitempty"`
	ProjectType               string  `json:"ProjectType,omitempty"`
	Responsibility            string  `json:"Responsibility,omitempty"`
	Referable                 string  `json:"Referable,omitempty"`
	EstCompletionDate         string  `json:"EstCompletionDate,omitempty"`
	ActCompletionDate         string  `json:"ActCompletionDate,omitempty"`
	ContractDate              string  `json:"ContractDate,omitempty"`
	BidDate                   string  `json:"BidDate,omitempty"`
	ComplDateComment          string  `json:"ComplDateComment,omitempty"`
	FirmCost                  float64 `json:"FirmCost,omitempty"`
	FirmCostComment           string  `json:"FirmCostComment,omitempty"`
	TotalProjectCost          float64 `json:"TotalProjectCost,omitempty"`
	TotalCostComment          string  `json:"TotalCostComment,omitempty"`
	ClientConfidential        string  `json:"ClientConfidential,omitempty"`
	ClientAlias               string  `json:"ClientAlias,omitempty"`
	AvailableForCRM           string  `json:"AvailableForCRM,omitempty"`
	ReadyForApproval          string  `json:"ReadyForApproval,omitempty"`
	ReadyForProcessing        string  `json:"ReadyForProcessing,omitempty"`
	BillingClientID           string  `json:"BillingClientID,omitempty"`
	BillingContactID          string  `json:"BillingContactID,omitempty"`
	Phone                     string  `json:"Phone,omitempty"`
	Fax                       string  `json:"Fax,omitempty"`
	EMail                     string  `json:"EMail,omitempty"`
	ProposalWBS1              string  `json:"ProposalWBS1,omitempty"`
	CostRateMeth              int     `json:"CostRateMeth,omitempty"`
	CostRateTableNo           int     `json:"CostRateTableNo,omitempty"`
	PayRateMeth               int     `json:"PayRateMeth,omitempty"`
	PayRateTableNo            int     `json:"PayRateTableNo,omitempty"`
	Locale                    string  `json:"Locale,omitempty"`
	LineItemApproval          string  `json:"LineItemApproval,omitempty"`
	LineItemApprovalEK        string  `json:"LineItemApprovalEK,omitempty"`
	BudgetSource              string  `json:"BudgetSource,omitempty"`
	BudgetLevel               string  `json:"BudgetLevel,omitempty"`
	ProfServicesComplDate     string  `json:"ProfServicesComplDate,omitempty"`
	ConstComplDate            string  `json:"ConstComplDate,omitempty"`
	ProjectCurrencyCode       string  `json:"ProjectCurrencyCode,omitempty"`
	ProjectExchangeRate       float64 `json:"ProjectExchangeRate,omitempty"`
	BillingCurrencyCode       string  `json:"BillingCurrencyCode,omitempty"`
	BillingExchangeRate       float64 `json:"BillingExchangeRate,omitempty"`
	RestrictChargeCompanies   string  `json:"RestrictChargeCompanies,omitempty"`
	FeeBillingCurrency        float64 `json:"FeeBillingCurrency,omitempty"`
	ReimbAllowBillingCurrency float64 `json:"ReimbAllowBillingCurrency,omitempty"`
	ConsultFeeBillingCurrency float64 `json:"ConsultFeeBillingCurrency,omitempty"`
	RevUpsetLimits            string  `json:"RevUpsetLimits,omitempty"`
	RevUpsetWBS2              string  `json:"RevUpsetWBS2,omitempty"`
	RevUpsetWBS3              string  `json:"RevUpsetWBS3,omitempty"`
	RevUpsetIncludeComp       string  `json:"RevUpsetIncludeComp,omitempty"`
	RevUpsetIncludeCons       string  `json:"RevUpsetIncludeCons,omitempty"`
	RevUpsetIncludeReimb      string  `json:"RevUpsetIncludeReimb,omitempty"`
	PORMBRate                 float64 `json:"PORMBRate,omitempty"`
	POCNSRate                 float64 `json:"POCNSRate,omitempty"`
	PlanID                    string  `json:"PlanID,omitempty"`
	TKCheckRPDate             string  `json:"TKCheckRPDate,omitempty"`
	CreateUser                string  `json:"CreateUser,omitempty"`
	CreateDate                string  `json:"CreateDate,omitempty"`
	ModUser                   string  `json:"ModUser,omitempty"`
	ModDate                   string  `json:"ModDate,omitempty"`
}

// ListProjects retrieves a list of projects from the Vantagepoint API.
// Pass a *Query to filter, paginate, or sort results; nil retrieves defaults.
func (c *Client) ListProjects(ctx context.Context, q *Query) ([]Project, error) {
	var projects []Project
	if err := c.get(ctx, "project", q, &projects); err != nil {
		return nil, fmt.Errorf("listing projects: %w", err)
	}
	return projects, nil
}

// GetProject retrieves a single project by its WBS1 code. The Vantagepoint API
// returns an array even for single-resource lookups, so the response is decoded
// as a slice and the first element is returned.
func (c *Client) GetProject(ctx context.Context, wbs1 string) (*Project, error) {
	var projects []Project
	if err := c.get(ctx, "project/"+wbs1, nil, &projects); err != nil {
		return nil, fmt.Errorf("getting project %s: %w", wbs1, err)
	}
	if len(projects) == 0 {
		return nil, fmt.Errorf("getting project %s: %w", wbs1, ErrNotFound)
	}
	return &projects[0], nil
}

// CreateProject creates a new project record in Vantagepoint.
func (c *Client) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	var result Project
	if err := c.post(ctx, "project", project, &result); err != nil {
		return nil, fmt.Errorf("creating project: %w", err)
	}
	return &result, nil
}

// UpdateProject updates an existing project record identified by WBS1.
func (c *Client) UpdateProject(ctx context.Context, wbs1 string, project *Project) (*Project, error) {
	var result Project
	if err := c.put(ctx, "project/"+wbs1, project, &result); err != nil {
		return nil, fmt.Errorf("updating project %s: %w", wbs1, err)
	}
	return &result, nil
}

// DeleteProject deletes a project record by its WBS1 code.
func (c *Client) DeleteProject(ctx context.Context, wbs1 string) error {
	if err := c.delete(ctx, "project/"+wbs1); err != nil {
		return fmt.Errorf("deleting project %s: %w", wbs1, err)
	}
	return nil
}

// GetProjectByHierarchy retrieves projects by their hierarchy parent ID.
// This returns child records under the specified parent in the project hierarchy.
func (c *Client) GetProjectByHierarchy(ctx context.Context, parentID string, q *Query) ([]Project, error) {
	var projects []Project
	if err := c.get(ctx, "project/hierarchy/"+parentID, q, &projects); err != nil {
		return nil, fmt.Errorf("getting project hierarchy for %s: %w", parentID, err)
	}
	return projects, nil
}

// ListPhases retrieves a list of phases (WBS2) for a given project.
// Phases are returned as Project records with SubLevel set to the WBS2 level.
func (c *Client) ListPhases(ctx context.Context, wbs1 string, q *Query) ([]Project, error) {
	if q == nil {
		q = NewQuery()
	}
	q.WBSType("wbs2")
	q.Filter("WBS1", "eq", wbs1)
	return c.ListProjects(ctx, q)
}

// ListTasks retrieves a list of tasks (WBS3) for a given project and phase.
// Tasks are returned as Project records with SubLevel set to the WBS3 level.
func (c *Client) ListTasks(ctx context.Context, wbs1, wbs2 string, q *Query) ([]Project, error) {
	if q == nil {
		q = NewQuery()
	}
	q.WBSType("wbs3")
	q.Filter("WBS1", "eq", wbs1)
	q.Filter("WBS2", "eq", wbs2)
	return c.ListProjects(ctx, q)
}

// ProjectEmployee represents an employee assignment on a project.
type ProjectEmployee struct {
	RecordID        string  `json:"RecordID,omitempty"`
	WBS1            string  `json:"WBS1,omitempty"`
	WBS2            string  `json:"WBS2,omitempty"`
	WBS3            string  `json:"WBS3,omitempty"`
	Employee        string  `json:"Employee,omitempty"`
	Role            string  `json:"Role,omitempty"`
	RoleDescription string  `json:"RoleDescription,omitempty"`
	TeamStatus      string  `json:"TeamStatus,omitempty"`
	StartDate       string  `json:"StartDate,omitempty"`
	EndDate         string  `json:"EndDate,omitempty"`
	CRMHours        float64 `json:"CRMHours,omitempty"`
	CreateUser      string  `json:"CreateUser,omitempty"`
	CreateDate      string  `json:"CreateDate,omitempty"`
	ModUser         string  `json:"ModUser,omitempty"`
	ModDate         string  `json:"ModDate,omitempty"`
}

// ProjectTeamMember represents a team member (employee or contact) on a project.
type ProjectTeamMember struct {
	WBS1               string  `json:"WBS1,omitempty"`
	WBS2               string  `json:"WBS2,omitempty"`
	WBS3               string  `json:"WBS3,omitempty"`
	RecordID           string  `json:"RecordID,omitempty"`
	PrimaryInd         string  `json:"PrimaryInd,omitempty"`
	Role               string  `json:"Role,omitempty"`
	RoleDescription    string  `json:"RoleDescription,omitempty"`
	CFGRoleDescription string  `json:"CFGRoleDescription,omitempty"`
	LfName             string  `json:"lfName,omitempty"`
	FlName             string  `json:"flName,omitempty"`
	CtName             string  `json:"ctName,omitempty"`
	FirmName           string  `json:"FirmName,omitempty"`
	FirmID             string  `json:"FirmID,omitempty"`
	Title              string  `json:"Title,omitempty"`
	Phone              string  `json:"Phone,omitempty"`
	Address1           string  `json:"Address1,omitempty"`
	City               string  `json:"City,omitempty"`
	State              string  `json:"State,omitempty"`
	Zip                string  `json:"Zip,omitempty"`
	Email              string  `json:"email,omitempty"`
	RecordType         string  `json:"RecordType,omitempty"`
	StartDate          string  `json:"StartDate,omitempty"`
	EndDate            string  `json:"EndDate,omitempty"`
	Hours              float64 `json:"Hours,omitempty"`
	Employee           string  `json:"Employee,omitempty"`
	CreateUser         string  `json:"CreateUser,omitempty"`
	CreateDate         string  `json:"CreateDate,omitempty"`
	ModUser            string  `json:"ModUser,omitempty"`
	ModDate            string  `json:"ModDate,omitempty"`
}

// ProjectRevenue represents a revenue allocation record on a project.
type ProjectRevenue struct {
	RevAllocID     string  `json:"RevAllocID,omitempty"`
	WBS1           string  `json:"WBS1,omitempty"`
	WBS2           string  `json:"WBS2,omitempty"`
	WBS3           string  `json:"WBS3,omitempty"`
	Description    string  `json:"Description,omitempty"`
	RevenueDate    string  `json:"RevenueDate,omitempty"`
	RevenueAmt     float64 `json:"RevenueAmt,omitempty"`
	PercentRevenue float64 `json:"PercentRevenue,omitempty"`
	CreateUser     string  `json:"CreateUser,omitempty"`
	CreateDate     string  `json:"CreateDate,omitempty"`
	ModUser        string  `json:"ModUser,omitempty"`
	ModDate        string  `json:"ModDate,omitempty"`
}

// ProjectFirmClient represents a firm or client association on a project.
type ProjectFirmClient struct {
	PKey               int     `json:"PKey,omitempty"`
	ClientID           string  `json:"ClientID,omitempty"`
	WBS1               string  `json:"WBS1,omitempty"`
	WBS2               string  `json:"WBS2,omitempty"`
	WBS3               string  `json:"WBS3,omitempty"`
	Role               string  `json:"Role,omitempty"`
	RoleDescription    string  `json:"RoleDescription,omitempty"`
	RoleName           string  `json:"RoleName,omitempty"`
	TeamStatus         string  `json:"TeamStatus,omitempty"`
	Address            string  `json:"Address,omitempty"`
	Address1           string  `json:"Address1,omitempty"`
	Address2           string  `json:"Address2,omitempty"`
	City               string  `json:"City,omitempty"`
	State              string  `json:"State,omitempty"`
	Country            string  `json:"Country,omitempty"`
	ZIP                string  `json:"ZIP,omitempty"`
	PrimaryInd         string  `json:"PrimaryInd,omitempty"`
	ClientConfidential string  `json:"ClientConfidential,omitempty"`
	ClientInd          string  `json:"ClientInd,omitempty"`
	VendorInd          string  `json:"VendorInd,omitempty"`
	ClientName         string  `json:"ClientName,omitempty"`
	AddressPhone       string  `json:"AddressPhone,omitempty"`
	CostAmount         float64 `json:"CostAmount,omitempty"`
	IsClient           string  `json:"IsClient,omitempty"`
	IsSubconsultant    string  `json:"IsSubconsultant,omitempty"`
	WebSite            string  `json:"WebSite,omitempty"`
	CreateUser         string  `json:"CreateUser,omitempty"`
	CreateDate         string  `json:"CreateDate,omitempty"`
	ModUser            string  `json:"ModUser,omitempty"`
	ModDate            string  `json:"ModDate,omitempty"`
}

// ListProjectEmployees retrieves the employee assignments for a project.
func (c *Client) ListProjectEmployees(ctx context.Context, wbsKey string, q *Query) ([]ProjectEmployee, error) {
	var employees []ProjectEmployee
	if err := c.get(ctx, "project/"+wbsKey+"/employee", q, &employees); err != nil {
		return nil, fmt.Errorf("listing project employees for %s: %w", wbsKey, err)
	}
	return employees, nil
}

// ListProjectTeamMembers retrieves the team members for a project.
func (c *Client) ListProjectTeamMembers(ctx context.Context, wbsKey string, q *Query) ([]ProjectTeamMember, error) {
	var members []ProjectTeamMember
	if err := c.get(ctx, "project/"+wbsKey+"/teammember", q, &members); err != nil {
		return nil, fmt.Errorf("listing project team members for %s: %w", wbsKey, err)
	}
	return members, nil
}

// ListProjectRevenue retrieves the revenue allocations for a project.
func (c *Client) ListProjectRevenue(ctx context.Context, wbsKey string, q *Query) ([]ProjectRevenue, error) {
	var revenue []ProjectRevenue
	if err := c.get(ctx, "project/"+wbsKey+"/revenue", q, &revenue); err != nil {
		return nil, fmt.Errorf("listing project revenue for %s: %w", wbsKey, err)
	}
	return revenue, nil
}

// ListProjectFirmClients retrieves the firm/client associations for a project.
func (c *Client) ListProjectFirmClients(ctx context.Context, wbsKey string, q *Query) ([]ProjectFirmClient, error) {
	var clients []ProjectFirmClient
	if err := c.get(ctx, "project/"+wbsKey+"/firmClient", q, &clients); err != nil {
		return nil, fmt.Errorf("listing project firm clients for %s: %w", wbsKey, err)
	}
	return clients, nil
}
