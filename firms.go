package vantagepoint

import (
	"context"
	"fmt"
)

// Firm represents a Vantagepoint firm (client/vendor) record. This struct covers
// core business fields from the ~123 fields returned by the API. Additional fields
// (billing/payment address blocks, custom cust* fields, owner denormalized fields,
// third-party integration fields) exist in the API response but are omitted here.
// Expand from the Postman collection as needed.
//
// Note: The primary key is ClientID (a GUID-style string), not a short code.
type Firm struct {
	ClientID                      string  `json:"ClientID,omitempty"`
	Client                        string  `json:"Client,omitempty"`
	Name                          string  `json:"Name,omitempty"`
	Status                        string  `json:"Status,omitempty"`
	ExportInd                     string  `json:"ExportInd,omitempty"`
	WebSite                       string  `json:"WebSite,omitempty"`
	Memo                          string  `json:"Memo,omitempty"`
	CurrentStatus                 string  `json:"CurrentStatus,omitempty"`
	CustomCurrencyCode            string  `json:"CustomCurrencyCode,omitempty"`
	ClientInd                     string  `json:"ClientInd,omitempty"`
	VendorInd                     string  `json:"VendorInd,omitempty"`
	PriorWork                     string  `json:"PriorWork,omitempty"`
	Recommend                     string  `json:"Recommend,omitempty"`
	DisadvBusiness                string  `json:"DisadvBusiness,omitempty"`
	MinorityBusiness              string  `json:"MinorityBusiness,omitempty"`
	SmallBusiness                 string  `json:"SmallBusiness,omitempty"`
	WomanOwned                    string  `json:"WomanOwned,omitempty"`
	VetOwnedSmallBusiness         string  `json:"VetOwnedSmallBusiness,omitempty"`
	DisabledVetOwnedSmallBusiness string  `json:"DisabledVetOwnedSmallBusiness,omitempty"`
	AlaskaNative                  string  `json:"AlaskaNative,omitempty"`
	HBCU                          string  `json:"HBCU,omitempty"`
	HUBZone                       string  `json:"HUBZone,omitempty"`
	EightA                        string  `json:"EightA,omitempty"`
	GovernmentAgency              string  `json:"GovernmentAgency,omitempty"`
	Competitor                    string  `json:"Competitor,omitempty"`
	Incumbent                     string  `json:"Incumbent,omitempty"`
	SpecialtyType                 string  `json:"SpecialtyType,omitempty"`
	Specialty                     string  `json:"Specialty,omitempty"`
	ParentID                      string  `json:"ParentID,omitempty"`
	ParentLevel1                  string  `json:"ParentLevel1,omitempty"`
	ParentLevel2                  string  `json:"ParentLevel2,omitempty"`
	ParentLevel3                  string  `json:"ParentLevel3,omitempty"`
	ParentLevel4                  string  `json:"ParentLevel4,omitempty"`
	ParentName                    string  `json:"ParentName,omitempty"`
	Employees                     int     `json:"Employees,omitempty"`
	AnnualRevenue                 float64 `json:"AnnualRevenue,omitempty"`
	Category                      string  `json:"Category,omitempty"`
	SortName                      string  `json:"SortName,omitempty"`
	Vendor                        string  `json:"Vendor,omitempty"`
	Type                          string  `json:"Type,omitempty"`
	TypeOfTIN                     string  `json:"TypeOfTIN,omitempty"`
	FedID                         string  `json:"FedID,omitempty"`
	Org                           string  `json:"Org,omitempty"`
	Owner                         string  `json:"Owner,omitempty"`
	AvailableForCRM               string  `json:"AvailableForCRM,omitempty"`
	ReadyForApproval              string  `json:"ReadyForApproval,omitempty"`
	ReadyForProcessing            string  `json:"ReadyForProcessing,omitempty"`
	HasPhoto                      int     `json:"HasPhoto,omitempty"`
	PhotoModDate                  string  `json:"PhotoModDate,omitempty"`
	CreateUser                    string  `json:"CreateUser,omitempty"`
	CreateDate                    string  `json:"CreateDate,omitempty"`
	ModUser                       string  `json:"ModUser,omitempty"`
	ModDate                       string  `json:"ModDate,omitempty"`
	PrimaryAddress1               string  `json:"PrimaryAddress1,omitempty"`
	PrimaryAddress2               string  `json:"PrimaryAddress2,omitempty"`
	PrimaryCity                   string  `json:"PrimaryCity,omitempty"`
	PrimaryState                  string  `json:"PrimaryState,omitempty"`
	PrimaryZip                    string  `json:"PrimaryZip,omitempty"`
	PrimaryCountry                string  `json:"PrimaryCountry,omitempty"`
	PrimaryPhone                  string  `json:"PrimaryPhone,omitempty"`
	PrimaryFax                    string  `json:"PrimaryFax,omitempty"`
	PrimaryEmail                  string  `json:"PrimaryEmail,omitempty"`
	HasHierarchy                  int     `json:"HasHierarchy,omitempty"`
}

// ListFirms retrieves a list of firms from the Vantagepoint API.
// Pass a *Query to filter, paginate, or sort results; nil retrieves defaults.
func (c *Client) ListFirms(ctx context.Context, q *Query) ([]Firm, error) {
	var firms []Firm
	if err := c.get(ctx, "firm", q, &firms); err != nil {
		return nil, fmt.Errorf("listing firms: %w", err)
	}
	return firms, nil
}

// GetFirm retrieves a single firm by its ClientID. The Vantagepoint API
// returns an array even for single-resource lookups, so the response is decoded
// as a slice and the first element is returned.
func (c *Client) GetFirm(ctx context.Context, clientID string) (*Firm, error) {
	var firms []Firm
	if err := c.get(ctx, "firm/"+clientID, nil, &firms); err != nil {
		return nil, fmt.Errorf("getting firm %s: %w", clientID, err)
	}
	if len(firms) == 0 {
		return nil, fmt.Errorf("getting firm %s: %w", clientID, ErrNotFound)
	}
	return &firms[0], nil
}

// CreateFirm creates a new firm record in Vantagepoint.
func (c *Client) CreateFirm(ctx context.Context, firm *Firm) (*Firm, error) {
	var result Firm
	if err := c.post(ctx, "firm", firm, &result); err != nil {
		return nil, fmt.Errorf("creating firm: %w", err)
	}
	return &result, nil
}

// UpdateFirm updates an existing firm record identified by ClientID.
func (c *Client) UpdateFirm(ctx context.Context, clientID string, firm *Firm) (*Firm, error) {
	var result Firm
	if err := c.put(ctx, "firm/"+clientID, firm, &result); err != nil {
		return nil, fmt.Errorf("updating firm %s: %w", clientID, err)
	}
	return &result, nil
}

// DeleteFirm deletes a firm record by its ClientID.
func (c *Client) DeleteFirm(ctx context.Context, clientID string) error {
	if err := c.delete(ctx, "firm/"+clientID); err != nil {
		return fmt.Errorf("deleting firm %s: %w", clientID, err)
	}
	return nil
}

// FirmEmployee represents an employee associated with a firm.
// Retrieved via GET /firm/{ClientID}/employee.
type FirmEmployee struct {
	Employee         string  `json:"Employee,omitempty"`
	ClientID         string  `json:"ClientID,omitempty"`
	Name             string  `json:"Name,omitempty"`
	Relationship     string  `json:"Relationship,omitempty"`
	RelationshipDesc string  `json:"RelationshipDesc,omitempty"`
	LastName         string  `json:"LastName,omitempty"`
	FirstName        string  `json:"FirstName,omitempty"`
	MiddleName       string  `json:"MiddleName,omitempty"`
	Title            string  `json:"Title,omitempty"`
	Status           string  `json:"Status,omitempty"`
	Type             string  `json:"Type,omitempty"`
	Org              string  `json:"Org,omitempty"`
	EMail            string  `json:"EMail,omitempty"`
	HomePhone        string  `json:"HomePhone,omitempty"`
	Address1         string  `json:"Address1,omitempty"`
	Address2         string  `json:"Address2,omitempty"`
	City             string  `json:"City,omitempty"`
	State            string  `json:"State,omitempty"`
	ZIP              string  `json:"ZIP,omitempty"`
	Country          string  `json:"Country,omitempty"`
	HireDate         string  `json:"HireDate,omitempty"`
	HomeCompany      string  `json:"HomeCompany,omitempty"`
	EmployeeCompany  string  `json:"EmployeeCompany,omitempty"`
	EmployeeName     string  `json:"EmployeeName,omitempty"`
	BillingCategory  int     `json:"BillingCategory,omitempty"`
	BillingPool      int     `json:"BillingPool,omitempty"`
	PayRate          float64 `json:"PayRate,omitempty"`
	PayType          string  `json:"PayType,omitempty"`
	JobCostRate      float64 `json:"JobCostRate,omitempty"`
	JobCostType      string  `json:"JobCostType,omitempty"`
	ProvCostRate     float64 `json:"ProvCostRate,omitempty"`
	ProvBillRate     float64 `json:"ProvBillRate,omitempty"`
	SSN              string  `json:"SSN,omitempty"`
	CreateUser       string  `json:"CreateUser,omitempty"`
	CreateDate       string  `json:"CreateDate,omitempty"`
	ModUser          string  `json:"ModUser,omitempty"`
	ModDate          string  `json:"ModDate,omitempty"`
}

// FirmProject represents a project associated with a firm.
// Retrieved via GET /firm/{ClientID}/project.
type FirmProject struct {
	ClientID        string `json:"ClientID,omitempty"`
	WBS1            string `json:"WBS1,omitempty"`
	Role            string `json:"Role,omitempty"`
	RoleName        string `json:"RoleName,omitempty"`
	RoleDescription string `json:"RoleDescription,omitempty"`
	ClientInd       string `json:"ClientInd,omitempty"`
	VendorInd       string `json:"VendorInd,omitempty"`
	PRName          string `json:"PRName,omitempty"`
	PRClientID      string `json:"PRClientID,omitempty"`
	PRClientName    string `json:"PRClientName,omitempty"`
	PRContactID     string `json:"PRContactID,omitempty"`
	PRContactName   string `json:"PRContactName,omitempty"`
	PRStatus        string `json:"PRStatus,omitempty"`
	DefaultType     string `json:"DefaultType,omitempty"`
	CreateUser      string `json:"CreateUser,omitempty"`
	CreateDate      string `json:"CreateDate,omitempty"`
	ModUser         string `json:"ModUser,omitempty"`
	ModDate         string `json:"ModDate,omitempty"`
}

// ListFirmEmployees retrieves employees associated with a firm.
// Uses GET /firm/{clientID}/employee.
func (c *Client) ListFirmEmployees(ctx context.Context, clientID string, q *Query) ([]FirmEmployee, error) {
	var results []FirmEmployee
	if err := c.get(ctx, "firm/"+clientID+"/employee", q, &results); err != nil {
		return nil, fmt.Errorf("listing firm employees for %s: %w", clientID, err)
	}
	return results, nil
}

// ListFirmProjects retrieves projects associated with a firm.
// Uses GET /firm/{clientID}/project.
func (c *Client) ListFirmProjects(ctx context.Context, clientID string, q *Query) ([]FirmProject, error) {
	var results []FirmProject
	if err := c.get(ctx, "firm/"+clientID+"/project", q, &results); err != nil {
		return nil, fmt.Errorf("listing firm projects for %s: %w", clientID, err)
	}
	return results, nil
}
