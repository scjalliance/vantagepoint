package vantagepoint

import (
	"context"
	"fmt"
)

// Employee represents a Vantagepoint employee record. This struct covers core
// business fields from the ~210 fields returned by the API. Additional fields
// (ADP/Paychex payroll, passport/visa, emergency contacts, mailing address,
// demographics, QBO integration, custom Cust* fields) exist in the API response
// but are omitted here. Expand from the Postman collection as needed.
//
// Note: VP uses quirky PascalCase for some fields (e.g., EMail, FAX, ZIP).
// Boolean values are "Y"/"N" strings, not actual booleans. Status values are
// "A" (active), "I" (inactive), "T" (terminated).
type Employee struct {
	Employee             string  `json:"Employee,omitempty"`
	PIMID                string  `json:"PIMID,omitempty"`
	EmployeeCompany      string  `json:"EmployeeCompany,omitempty"`
	Org                  string  `json:"Org,omitempty"`
	Supervisor           string  `json:"Supervisor,omitempty"`
	SupervisorName       string  `json:"SupervisorName,omitempty"`
	HireDate             string  `json:"HireDate,omitempty"`
	Status               string  `json:"Status,omitempty"`
	YearsOtherFirms      int     `json:"YearsOtherFirms,omitempty"`
	PriorYearsFirm       int     `json:"PriorYearsFirm,omitempty"`
	ReadyForProcessing   string  `json:"ReadyForProcessing,omitempty"`
	PayRate              float64 `json:"PayRate,omitempty"`
	PayOvtPct            float64 `json:"PayOvtPct,omitempty"`
	PaySpecialOvtPct     float64 `json:"PaySpecialOvtPct,omitempty"`
	PayType              string  `json:"PayType,omitempty"`
	UseTotalHrsAsStd     string  `json:"UseTotalHrsAsStd,omitempty"`
	ChangeDefaultLC      string  `json:"ChangeDefaultLC,omitempty"`
	Locale               string  `json:"Locale,omitempty"`
	LocaleMethod         string  `json:"LocaleMethod,omitempty"`
	HomeCompany          string  `json:"HomeCompany,omitempty"`
	Salutation           string  `json:"Salutation,omitempty"`
	FirstName            string  `json:"FirstName,omitempty"`
	MiddleName           string  `json:"MiddleName,omitempty"`
	LastName             string  `json:"LastName,omitempty"`
	Suffix               string  `json:"Suffix,omitempty"`
	PreferredName        string  `json:"PreferredName,omitempty"`
	Title                string  `json:"Title,omitempty"`
	EMail                string  `json:"EMail,omitempty"`
	WorkPhone            string  `json:"WorkPhone,omitempty"`
	WorkPhoneExt         string  `json:"WorkPhoneExt,omitempty"`
	MobilePhone          string  `json:"MobilePhone,omitempty"`
	HomePhone            string  `json:"HomePhone,omitempty"`
	FAX                  string  `json:"FAX,omitempty"`
	AvailableForCRM      string  `json:"AvailableForCRM,omitempty"`
	ReadyForApproval     string  `json:"ReadyForApproval,omitempty"`
	Address1             string  `json:"Address1,omitempty"`
	Address2             string  `json:"Address2,omitempty"`
	Address3             string  `json:"Address3,omitempty"`
	City                 string  `json:"City,omitempty"`
	State                string  `json:"State,omitempty"`
	ZIP                  string  `json:"ZIP,omitempty"`
	Country              string  `json:"Country,omitempty"`
	Memo                 string  `json:"Memo,omitempty"`
	UtilizationRatio     float64 `json:"UtilizationRatio,omitempty"`
	TargetRatio          float64 `json:"TargetRatio,omitempty"`
	Language             string  `json:"Language,omitempty"`
	ConsultantInd        string  `json:"ConsultantInd,omitempty"`
	ClientVendorInd      string  `json:"ClientVendorInd,omitempty"`
	Vendor               string  `json:"Vendor,omitempty"`
	ClientID             string  `json:"ClientID,omitempty"`
	HasPhoto             int     `json:"HasPhoto,omitempty"`
	PhotoModDate         string  `json:"PhotoModDate,omitempty"`
	CreateUser           string  `json:"CreateUser,omitempty"`
	CreateDate           string  `json:"CreateDate,omitempty"`
	ModUser              string  `json:"ModUser,omitempty"`
	ModDate              string  `json:"ModDate,omitempty"`
	HoursPerDay          float64 `json:"HoursPerDay,omitempty"`
	RaiseDate            string  `json:"RaiseDate,omitempty"`
	TerminationDate      string  `json:"TerminationDate,omitempty"`
	SSN                  string  `json:"SSN,omitempty"`
	JobCostRate          float64 `json:"JobCostRate,omitempty"`
	JCOvtPct             float64 `json:"JCOvtPct,omitempty"`
	JCSpecialOvtPct      float64 `json:"JCSpecialOvtPct,omitempty"`
	JobCostType          string  `json:"JobCostType,omitempty"`
	ProvCostRate         float64 `json:"ProvCostRate,omitempty"`
	ProvCostOTPct        float64 `json:"ProvCostOTPct,omitempty"`
	ProvCostSpecialOTPct float64 `json:"ProvCostSpecialOTPct,omitempty"`
	ProvBillRate         float64 `json:"ProvBillRate,omitempty"`
	ProvBillOTPct        float64 `json:"ProvBillOTPct,omitempty"`
	ProvBillSpecialOTPct float64 `json:"ProvBillSpecialOTPct,omitempty"`
	Type                 string  `json:"Type,omitempty"`
	BillingCategory      int     `json:"BillingCategory,omitempty"`
	BillingCategoryCode  string  `json:"BillingCategoryCode,omitempty"`
}

// ListEmployees retrieves a list of employees from the Vantagepoint API.
// Pass a *Query to filter, paginate, or sort results; nil retrieves defaults.
func (c *Client) ListEmployees(ctx context.Context, q *Query) ([]Employee, error) {
	var employees []Employee
	if err := c.get(ctx, "employee", q, &employees); err != nil {
		return nil, fmt.Errorf("listing employees: %w", err)
	}
	return employees, nil
}

// GetEmployee retrieves a single employee by their employee code. The
// Vantagepoint API returns an array even for single-resource lookups, so the
// response is decoded as a slice and the first element is returned.
func (c *Client) GetEmployee(ctx context.Context, employeeCode string) (*Employee, error) {
	var employees []Employee
	if err := c.get(ctx, "employee/"+employeeCode, nil, &employees); err != nil {
		return nil, fmt.Errorf("getting employee %s: %w", employeeCode, err)
	}
	if len(employees) == 0 {
		return nil, fmt.Errorf("getting employee %s: %w", employeeCode, ErrNotFound)
	}
	return &employees[0], nil
}

// CreateEmployee creates a new employee record in Vantagepoint.
func (c *Client) CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error) {
	var result Employee
	if err := c.post(ctx, "employee", employee, &result); err != nil {
		return nil, fmt.Errorf("creating employee: %w", err)
	}
	return &result, nil
}

// UpdateEmployee updates an existing employee record identified by employee code.
func (c *Client) UpdateEmployee(ctx context.Context, employeeCode string, employee *Employee) (*Employee, error) {
	var result Employee
	if err := c.put(ctx, "employee/"+employeeCode, employee, &result); err != nil {
		return nil, fmt.Errorf("updating employee %s: %w", employeeCode, err)
	}
	return &result, nil
}

// DeleteEmployee deletes an employee record by their employee code.
func (c *Client) DeleteEmployee(ctx context.Context, employeeCode string) error {
	if err := c.delete(ctx, "employee/"+employeeCode); err != nil {
		return fmt.Errorf("deleting employee %s: %w", employeeCode, err)
	}
	return nil
}

// EmployeeProject represents a project assignment for an employee.
// Retrieved via GET /employee/{Employee}/projects.
type EmployeeProject struct {
	RecordID        string `json:"RecordID,omitempty"`
	WBS1            string `json:"WBS1,omitempty"`
	WBS2            string `json:"WBS2,omitempty"`
	WBS3            string `json:"WBS3,omitempty"`
	Employee        string `json:"Employee,omitempty"`
	Role            string `json:"Role,omitempty"`
	RoleDescription string `json:"RoleDescription,omitempty"`
	StartDate       string `json:"StartDate,omitempty"`
	EndDate         string `json:"EndDate,omitempty"`
}

// ListEmployeeProjects retrieves projects assigned to an employee.
// Uses GET /employee/{employeeCode}/projects.
func (c *Client) ListEmployeeProjects(ctx context.Context, employeeCode string, q *Query) ([]EmployeeProject, error) {
	var results []EmployeeProject
	if err := c.get(ctx, "employee/"+employeeCode+"/projects", q, &results); err != nil {
		return nil, fmt.Errorf("listing projects for employee %s: %w", employeeCode, err)
	}
	return results, nil
}
