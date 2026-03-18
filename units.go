package vantagepoint

import (
	"context"
	"fmt"
)

// Unit entries in Vantagepoint follow the three-tier DataEntry pattern:
// unControl → unMaster → unDetail.

// UnitLanding represents a batch summary returned by the unLanding endpoint.
// This shares the same schema as TimesheetLanding — Period is a string here
// (e.g. "06/2008") versus an int in UnitControl.
type UnitLanding struct {
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

// UnitControl represents a unit batch control record.
type UnitControl struct {
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
	TSCount     int     `json:"TSCount,omitempty"`
}

// UnitMaster represents a per-unit master record within a batch.
type UnitMaster struct {
	Batch                          string `json:"Batch,omitempty"`
	Unit                           string `json:"Unit,omitempty"`
	Name                           string `json:"Name,omitempty"`
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

// UnitDetail represents an individual unit entry line.
type UnitDetail struct {
	Batch                        string  `json:"Batch,omitempty"`
	Unit                         string  `json:"Unit,omitempty"`
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
	TkPKey                       string  `json:"tkPKey,omitempty"`
}

// ListUnitBatches retrieves a list of unit batch summaries from the
// DataEntry unLanding endpoint. Pass a *Query to filter, paginate, or sort
// results; nil retrieves defaults.
func (c *Client) ListUnitBatches(ctx context.Context, q *Query) ([]UnitLanding, error) {
	var results []UnitLanding
	if err := c.get(ctx, "DataEntry/unLanding", q, &results); err != nil {
		return nil, fmt.Errorf("listing unit batches: %w", err)
	}
	return results, nil
}

// GetUnitControl retrieves the control records for a unit batch identified
// by the given batch key.
func (c *Client) GetUnitControl(ctx context.Context, batch string, q *Query) ([]UnitControl, error) {
	var results []UnitControl
	if err := c.get(ctx, "DataEntry/unControl/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting unit control for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetUnitMaster retrieves the per-unit master records for a unit batch
// identified by the given batch key.
func (c *Client) GetUnitMaster(ctx context.Context, batch string, q *Query) ([]UnitMaster, error) {
	var results []UnitMaster
	if err := c.get(ctx, "DataEntry/unMaster/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting unit master for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetUnitDetail retrieves the individual unit entry lines for a unit batch
// identified by the given batch key.
func (c *Client) GetUnitDetail(ctx context.Context, batch string, q *Query) ([]UnitDetail, error) {
	var results []UnitDetail
	if err := c.get(ctx, "DataEntry/unDetail/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting unit detail for batch %s: %w", batch, err)
	}
	return results, nil
}

// CreateUnitBatch creates a new unit batch via the unControl endpoint.
// Uses POST /DataEntry/unControl.
func (c *Client) CreateUnitBatch(ctx context.Context, control *UnitControl) ([]UnitControl, error) {
	var results []UnitControl
	if err := c.post(ctx, "DataEntry/unControl", control, &results); err != nil {
		return nil, fmt.Errorf("creating unit batch: %w", err)
	}
	return results, nil
}

// UpdateUnitBatch updates an existing unit batch control record identified
// by the given batch key.
// Uses PUT /DataEntry/unControl/{batch}.
func (c *Client) UpdateUnitBatch(ctx context.Context, batch string, control *UnitControl) ([]UnitControl, error) {
	var results []UnitControl
	if err := c.put(ctx, "DataEntry/unControl/"+batch, control, &results); err != nil {
		return nil, fmt.Errorf("updating unit batch %s: %w", batch, err)
	}
	return results, nil
}

// DeleteUnitBatch deletes a unit batch.
// Uses DELETE /DataEntry/unControl/{batch}.
func (c *Client) DeleteUnitBatch(ctx context.Context, batch string) error {
	if err := c.delete(ctx, "DataEntry/unControl/"+batch); err != nil {
		return fmt.Errorf("deleting unit batch %s: %w", batch, err)
	}
	return nil
}
