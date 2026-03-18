package vantagepoint

import (
	"context"
	"fmt"
)

// Employee expenses in Vantagepoint follow the three-tier DataEntry pattern:
// exControl → exMaster → exDetail.

// ExpenseControl represents an expense batch control record.
type ExpenseControl struct {
	Batch               string  `json:"Batch,omitempty"`
	Description         string  `json:"Description,omitempty"`
	Recurring           string  `json:"Recurring,omitempty"`
	EndDate             string  `json:"EndDate,omitempty"`
	DefaultCurrencyCode string  `json:"DefaultCurrencyCode,omitempty"`
	Total               float64 `json:"Total,omitempty"`
	SumTotal            float64 `json:"SumTotal,omitempty"`
	DiffTotal           float64 `json:"DiffTotal,omitempty"`
	AdvanceAmount       float64 `json:"AdvanceAmount,omitempty"`
	SumAdvanceAmount    float64 `json:"SumAdvanceAmount,omitempty"`
	DiffAdvanceAmount   float64 `json:"DiffAdvanceAmount,omitempty"`
	Selected            string  `json:"Selected,omitempty"`
	Posted              string  `json:"Posted,omitempty"`
	Creator             string  `json:"Creator,omitempty"`
	Period              int     `json:"Period,omitempty"`
	Company             string  `json:"Company,omitempty"`
	Diary               string  `json:"Diary,omitempty"`
	DiaryNo             int     `json:"DiaryNo,omitempty"`
}

// ExpenseMaster represents a per-employee expense report within a batch.
type ExpenseMaster struct {
	Employee                       string  `json:"Employee,omitempty"`
	EmplName                       string  `json:"EmplName,omitempty"`
	ReportDate                     string  `json:"ReportDate,omitempty"`
	ReportName                     string  `json:"ReportName,omitempty"`
	AdvanceAmount                  float64 `json:"AdvanceAmount,omitempty"`
	TotalDue                       float64 `json:"TotalDue,omitempty"`
	Batch                          string  `json:"Batch,omitempty"`
	MasterPKey                     string  `json:"MasterPKey,omitempty"`
	Posted                         string  `json:"Posted,omitempty"`
	Seq                            int     `json:"Seq,omitempty"`
	DefaultCurrencyCode            string  `json:"DefaultCurrencyCode,omitempty"`
	CurrencyExchangeOverrideMethod string  `json:"CurrencyExchangeOverrideMethod,omitempty"`
	CurrencyExchangeOverrideDate   string  `json:"CurrencyExchangeOverrideDate,omitempty"`
	CurrencyExchangeOverrideRate   float64 `json:"CurrencyExchangeOverrideRate,omitempty"`
	PaymentCurrencyCode            string  `json:"PaymentCurrencyCode,omitempty"`
	PaymentExchangeOverrideMethod  string  `json:"PaymentExchangeOverrideMethod,omitempty"`
	PaymentExchangeOverrideDate    string  `json:"PaymentExchangeOverrideDate,omitempty"`
	PaymentExchangeOverrideRate    float64 `json:"PaymentExchangeOverrideRate,omitempty"`
	Status                         string  `json:"Status,omitempty"`
	StatusDescription              string  `json:"StatusDescription,omitempty"`
	AuthorizedByName               string  `json:"AuthorizedByName,omitempty"`
	AuthorizedBy                   string  `json:"AuthorizedBy,omitempty"`
	RejectReason                   string  `json:"RejectReason,omitempty"`
	ModUser                        string  `json:"ModUser,omitempty"`
	ModUserEmployee                string  `json:"ModUserEmployee,omitempty"`
	ModDate                        string  `json:"ModDate,omitempty"`
}

// ExpenseDetail represents an individual expense line item.
type ExpenseDetail struct {
	Batch                        string  `json:"Batch,omitempty"`
	MasterPKey                   string  `json:"MasterPKey,omitempty"`
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
	Org                          string  `json:"Org,omitempty"`
	Account                      string  `json:"Account,omitempty"`
	AccountName                  string  `json:"AccountName,omitempty"`
	AccountCurrencyCode          string  `json:"AccountCurrencyCode,omitempty"`
	GlobalAccount                string  `json:"GlobalAccount,omitempty"`
	AccountType                  string  `json:"AccountType,omitempty"`
	CurrencyCode                 string  `json:"CurrencyCode,omitempty"`
	Amount                       float64 `json:"Amount,omitempty"`
	TotalTax                     float64 `json:"TotalTax,omitempty"`
	OriginatingVendor            string  `json:"OriginatingVendor,omitempty"`
	NetAmount                    float64 `json:"NetAmount,omitempty"`
	CurrencyExchangeOverrideRate float64 `json:"CurrencyExchangeOverrideRate,omitempty"`
	PaymentAmount                float64 `json:"PaymentAmount,omitempty"`
	PaymentExchangeRate          float64 `json:"PaymentExchangeRate,omitempty"`
	PaymentExchangeInfo          string  `json:"PaymentExchangeInfo,omitempty"`
	SuppressBill                 string  `json:"SuppressBill,omitempty"`
	OriginatingVendorName        string  `json:"OriginatingVendorName,omitempty"`
	DefaultsSet                  string  `json:"DefaultsSet,omitempty"`
	TaxToolTipColumn             string  `json:"TaxToolTipColumn,omitempty"`
}

// ListExpenseBatches retrieves a list of expense batch summaries.
// Uses GET /DataEntry/exLanding.
func (c *Client) ListExpenseBatches(ctx context.Context, q *Query) ([]ExpenseControl, error) {
	var results []ExpenseControl
	if err := c.get(ctx, "DataEntry/exLanding", q, &results); err != nil {
		return nil, fmt.Errorf("listing expense batches: %w", err)
	}
	return results, nil
}

// GetExpenseControl retrieves the control record for an expense batch.
// Uses GET /DataEntry/exControl/{batch}.
func (c *Client) GetExpenseControl(ctx context.Context, batch string, q *Query) ([]ExpenseControl, error) {
	var results []ExpenseControl
	if err := c.get(ctx, "DataEntry/exControl/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting expense control for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetExpenseMaster retrieves the per-employee master records for an expense batch.
// Uses GET /DataEntry/exMaster/{batch}.
func (c *Client) GetExpenseMaster(ctx context.Context, batch string, q *Query) ([]ExpenseMaster, error) {
	var results []ExpenseMaster
	if err := c.get(ctx, "DataEntry/exMaster/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting expense master for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetExpenseDetail retrieves the individual expense lines for an expense batch.
// Uses GET /DataEntry/exDetail/{batch}.
func (c *Client) GetExpenseDetail(ctx context.Context, batch string, q *Query) ([]ExpenseDetail, error) {
	var results []ExpenseDetail
	if err := c.get(ctx, "DataEntry/exDetail/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting expense detail for batch %s: %w", batch, err)
	}
	return results, nil
}

// CreateExpenseBatch creates a new expense batch via the exControl endpoint.
// Uses POST /DataEntry/exControl.
func (c *Client) CreateExpenseBatch(ctx context.Context, control *ExpenseControl) ([]ExpenseControl, error) {
	var results []ExpenseControl
	if err := c.post(ctx, "DataEntry/exControl", control, &results); err != nil {
		return nil, fmt.Errorf("creating expense batch: %w", err)
	}
	return results, nil
}

// UpdateExpenseBatch updates an existing expense batch control record.
// Uses PUT /DataEntry/exControl/{batch}.
func (c *Client) UpdateExpenseBatch(ctx context.Context, batch string, control *ExpenseControl) ([]ExpenseControl, error) {
	var results []ExpenseControl
	if err := c.put(ctx, "DataEntry/exControl/"+batch, control, &results); err != nil {
		return nil, fmt.Errorf("updating expense batch %s: %w", batch, err)
	}
	return results, nil
}

// DeleteExpenseBatch deletes an expense batch.
// Uses DELETE /DataEntry/exControl/{batch}.
func (c *Client) DeleteExpenseBatch(ctx context.Context, batch string) error {
	if err := c.delete(ctx, "DataEntry/exControl/"+batch); err != nil {
		return fmt.Errorf("deleting expense batch %s: %w", batch, err)
	}
	return nil
}

// CreateExpenseMaster creates a per-employee expense report within a batch.
// Uses POST /DataEntry/exMaster/{batch|masterPKey}.
func (c *Client) CreateExpenseMaster(ctx context.Context, batch, masterPKey string, master *ExpenseMaster) ([]ExpenseMaster, error) {
	var results []ExpenseMaster
	if err := c.post(ctx, "DataEntry/exMaster/"+batch+"|"+masterPKey, master, &results); err != nil {
		return nil, fmt.Errorf("creating expense master for batch %s: %w", batch, err)
	}
	return results, nil
}

// UpdateExpenseMaster updates a per-employee expense report within a batch.
// Uses PUT /DataEntry/exMaster/{batch|masterPKey}.
func (c *Client) UpdateExpenseMaster(ctx context.Context, batch, masterPKey string, master *ExpenseMaster) ([]ExpenseMaster, error) {
	var results []ExpenseMaster
	if err := c.put(ctx, "DataEntry/exMaster/"+batch+"|"+masterPKey, master, &results); err != nil {
		return nil, fmt.Errorf("updating expense master for batch %s: %w", batch, err)
	}
	return results, nil
}

// DeleteExpenseMaster deletes a per-employee expense report from a batch.
// Uses DELETE /DataEntry/exMaster/{batch|masterPKey}.
func (c *Client) DeleteExpenseMaster(ctx context.Context, batch, masterPKey string) error {
	if err := c.delete(ctx, "DataEntry/exMaster/"+batch+"|"+masterPKey); err != nil {
		return fmt.Errorf("deleting expense master for batch %s: %w", batch, err)
	}
	return nil
}
