package vantagepoint

import (
	"context"
	"fmt"
)

// Units By Project DataEntry endpoints follow the Landing/Control/Master/Detail
// tier pattern. The API prefix is "upLanding", "upControl", "upMaster", "upDetail".
// Unlike the standard Units DataEntry (utControl), the master record uses RefNo
// instead of Unit/Name, and the detail record includes both RefNo and Unit.
// The control record omits TSCount.

// UnitByProjectLanding represents a batch summary returned by the upLanding endpoint.
// This shares the same schema as the standard UnitLanding type, where Period is a
// formatted string and RawPeriod is the integer form.
type UnitByProjectLanding struct {
	Batch               string `json:"Batch,omitempty"`
	Description         string `json:"Description,omitempty"`
	Recurring           string `json:"Recurring,omitempty"`
	EndDate             string `json:"EndDate,omitempty"`
	Creator             string `json:"Creator,omitempty"`
	Period              string `json:"Period,omitempty"`
	RawPeriod           int    `json:"RawPeriod,omitempty"`
	Company             string `json:"Company,omitempty"`
	TransType           string `json:"TransType,omitempty"`
	ApprovalStatus      string `json:"ApprovalStatus,omitempty"`
	ApprovalStatusDesc  string `json:"ApprovalStatusDesc,omitempty"`
	AutoPost            string `json:"AutoPost,omitempty"`
	PostComment         string `json:"PostComment,omitempty"`
	PostDate            string `json:"PostDate,omitempty"`
	PostLogReportName   string `json:"PostLogReportName,omitempty"`
	PostLogReportPath   string `json:"PostLogReportPath,omitempty"`
	PostSeq             int    `json:"PostSeq,omitempty"`
	PostStatus          string `json:"PostStatus,omitempty"`
	PostStatusDesc      string `json:"PostStatusDesc,omitempty"`
	PostUser            string `json:"PostUser,omitempty"`
	RecurSchedule       string `json:"RecurSchedule,omitempty"`
	Reverse             string `json:"Reverse,omitempty"`
	Reversed            string `json:"Reversed,omitempty"`
	TransListReportPath string `json:"TransListReportPath,omitempty"`
}

// UnitByProjectControl represents a units-by-project batch control record.
// Unlike UnitControl, this type does not include a TSCount field.
type UnitByProjectControl struct {
	Batch       string  `json:"Batch,omitempty"`
	Description string  `json:"Description,omitempty"`
	Recurring   string  `json:"Recurring,omitempty"`
	EndDate     string  `json:"EndDate,omitempty"`
	Total       float64 `json:"Total,omitempty"`
	SumTotal    float64 `json:"SumTotal,omitempty"`
	DiffTotal   float64 `json:"DiffTotal,omitempty"`
	Selected    string  `json:"Selected,omitempty"`
	Posted      string  `json:"Posted,omitempty"`
	Creator     string  `json:"Creator,omitempty"`
	Period      int     `json:"Period,omitempty"`
	Company     string  `json:"Company,omitempty"`
	Diary       string  `json:"Diary,omitempty"`
	DiaryNo     int     `json:"DiaryNo,omitempty"`
}

// UnitByProjectMaster represents a per-project master record within a
// units-by-project batch. Unlike UnitMaster, this type uses RefNo instead
// of Unit/Name as the primary identifier.
type UnitByProjectMaster struct {
	Batch                          string `json:"Batch,omitempty"`
	RefNo                          string `json:"RefNo,omitempty"`
	Posted                         string `json:"Posted,omitempty"`
	Seq                            int    `json:"Seq,omitempty"`
	CurrencyExchangeOverrideMethod string `json:"CurrencyExchangeOverrideMethod,omitempty"`
	CurrencyExchangeOverrideDate   string `json:"CurrencyExchangeOverrideDate,omitempty"`
	Status                         string `json:"Status,omitempty"`
	StatusDescription              string `json:"StatusDescription,omitempty"`
	AuthorizedByName               string `json:"AuthorizedByName,omitempty"`
	AuthorizedBy                   string `json:"AuthorizedBy,omitempty"`
	RejectReason                   string `json:"RejectReason,omitempty"`
	ModUser                        string `json:"ModUser,omitempty"`
	ModUserEmployee                string `json:"ModUserEmployee,omitempty"`
	ModDate                        string `json:"ModDate,omitempty"`
}

// UnitByProjectDetail represents an individual unit-by-project line item.
// This type includes both RefNo (project reference) and Unit (unit code),
// and does not include a TkPKey field (unlike UnitDetail).
type UnitByProjectDetail struct {
	Batch                        string  `json:"Batch,omitempty"`
	RefNo                        string  `json:"RefNo,omitempty"`
	PKey                         string  `json:"PKey,omitempty"`
	Seq                          int     `json:"Seq,omitempty"`
	TransDate                    string  `json:"TransDate,omitempty"`
	Description                  string  `json:"Description,omitempty"`
	WBS1                         string  `json:"WBS1,omitempty"`
	WBS1Name                     string  `json:"WBS1Name,omitempty"`
	ClientName                   string  `json:"ClientName,omitempty"`
	WBS2Level                    string  `json:"WBS2Level,omitempty"`
	ChargeType                   string  `json:"ChargeType,omitempty"`
	WBS2                         string  `json:"WBS2,omitempty"`
	WBS2Name                     string  `json:"WBS2Name,omitempty"`
	WBS3Level                    string  `json:"WBS3Level,omitempty"`
	WBS3                         string  `json:"WBS3,omitempty"`
	WBS3Name                     string  `json:"WBS3Name,omitempty"`
	FunctionalCurrencyCode       string  `json:"FunctionalCurrencyCode,omitempty"`
	CurrencyCode                 string  `json:"CurrencyCode,omitempty"`
	BillingCurrencyCode          string  `json:"BillingCurrencyCode,omitempty"`
	Org                          string  `json:"Org,omitempty"`
	DefaultUnitTable             string  `json:"DefaultUnitTable,omitempty"`
	Unit                         string  `json:"Unit,omitempty"`
	UnitName                     string  `json:"UnitName,omitempty"`
	CostRate                     float64 `json:"CostRate,omitempty"`
	BillingRate                  float64 `json:"BillingRate,omitempty"`
	SingleLabel                  string  `json:"SingleLabel,omitempty"`
	TableNo                      string  `json:"TableNo,omitempty"`
	Account                      string  `json:"Account,omitempty"`
	AccountName                  string  `json:"AccountName,omitempty"`
	AccountCurrencyCode          string  `json:"AccountCurrencyCode,omitempty"`
	GlobalAccount                string  `json:"GlobalAccount,omitempty"`
	AccountType                  int     `json:"AccountType,omitempty"`
	Quantity                     float64 `json:"Quantity,omitempty"`
	Employee                     string  `json:"Employee,omitempty"`
	EmplName                     string  `json:"EmplName,omitempty"`
	EmployeeSpecificRevenue      string  `json:"EmployeeSpecificRevenue,omitempty"`
	EmployeeCompany              string  `json:"EmployeeCompany,omitempty"`
	Document                     string  `json:"Document,omitempty"`
	DocumentToolTipColumn        string  `json:"DocumentToolTipColumn,omitempty"`
	CurrencyExchangeOverrideRate float64 `json:"CurrencyExchangeOverrideRate,omitempty"`
}

// ListUnitByProjectBatches retrieves a list of units-by-project batch summaries
// from the DataEntry upLanding endpoint.
func (c *Client) ListUnitByProjectBatches(ctx context.Context, q *Query) ([]UnitByProjectLanding, error) {
	var results []UnitByProjectLanding
	if err := c.get(ctx, "DataEntry/upLanding", q, &results); err != nil {
		return nil, fmt.Errorf("listing unit-by-project batches: %w", err)
	}
	return results, nil
}

// GetUnitByProjectControl retrieves the control records for a units-by-project
// batch identified by the given batch key.
func (c *Client) GetUnitByProjectControl(ctx context.Context, batch string, q *Query) ([]UnitByProjectControl, error) {
	var results []UnitByProjectControl
	if err := c.get(ctx, "DataEntry/upControl/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting unit-by-project control for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetUnitByProjectMaster retrieves the per-project master records for a
// units-by-project batch identified by the given batch key.
func (c *Client) GetUnitByProjectMaster(ctx context.Context, batch string, q *Query) ([]UnitByProjectMaster, error) {
	var results []UnitByProjectMaster
	if err := c.get(ctx, "DataEntry/upMaster/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting unit-by-project master for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetUnitByProjectDetail retrieves the individual unit-by-project line items
// for a batch identified by the given batch key.
func (c *Client) GetUnitByProjectDetail(ctx context.Context, batch string, q *Query) ([]UnitByProjectDetail, error) {
	var results []UnitByProjectDetail
	if err := c.get(ctx, "DataEntry/upDetail/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting unit-by-project detail for batch %s: %w", batch, err)
	}
	return results, nil
}

// CreateUnitByProjectBatch creates a new units-by-project batch via the
// upControl endpoint.
func (c *Client) CreateUnitByProjectBatch(ctx context.Context, control *UnitByProjectControl) ([]UnitByProjectControl, error) {
	var results []UnitByProjectControl
	if err := c.post(ctx, "DataEntry/upControl", control, &results); err != nil {
		return nil, fmt.Errorf("creating unit-by-project batch: %w", err)
	}
	return results, nil
}

// UpdateUnitByProjectBatch updates an existing units-by-project batch control
// record identified by the given batch key.
func (c *Client) UpdateUnitByProjectBatch(ctx context.Context, batch string, control *UnitByProjectControl) ([]UnitByProjectControl, error) {
	var results []UnitByProjectControl
	if err := c.put(ctx, "DataEntry/upControl/"+batch, control, &results); err != nil {
		return nil, fmt.Errorf("updating unit-by-project batch %s: %w", batch, err)
	}
	return results, nil
}

// DeleteUnitByProjectBatch deletes a units-by-project batch via the upControl
// endpoint.
func (c *Client) DeleteUnitByProjectBatch(ctx context.Context, batch string) error {
	if err := c.delete(ctx, "DataEntry/upControl/"+batch); err != nil {
		return fmt.Errorf("deleting unit-by-project batch %s: %w", batch, err)
	}
	return nil
}
