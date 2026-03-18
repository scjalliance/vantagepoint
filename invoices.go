package vantagepoint

import (
	"context"
	"fmt"
)

// Invoices (AP voucher entry) in Vantagepoint follow the three-tier DataEntry
// pattern: inControl -> inMaster -> inDetail.

// InvoiceControl represents an invoice batch control record.
type InvoiceControl struct {
	Batch          string  `json:"Batch,omitempty"`
	Description    string  `json:"Description,omitempty"`
	Recurring      string  `json:"Recurring,omitempty"`
	EndDate        string  `json:"EndDate,omitempty"`
	DefaultTaxCode string  `json:"DefaultTaxCode,omitempty"`
	Total          float64 `json:"Total,omitempty"`
	SumTotal       float64 `json:"SumTotal,omitempty"`
	DiffTotal      float64 `json:"DiffTotal,omitempty"`
	RetTotal       float64 `json:"RetTotal,omitempty"`
	SumRetTotal    float64 `json:"SumRetTotal,omitempty"`
	DiffRetTotal   float64 `json:"DiffRetTotal,omitempty"`
	Selected       string  `json:"Selected,omitempty"`
	Posted         string  `json:"Posted,omitempty"`
	Creator        string  `json:"Creator,omitempty"`
	Period         int     `json:"Period,omitempty"`
	Company        string  `json:"Company,omitempty"`
}

// InvoiceMaster represents a per-invoice record within a batch.
type InvoiceMaster struct {
	Batch                          string  `json:"Batch,omitempty"`
	Invoice                        string  `json:"Invoice,omitempty"`
	WBS1                           string  `json:"WBS1,omitempty"`
	WBS1Name                       string  `json:"WBS1Name,omitempty"`
	ClientName                     string  `json:"ClientName,omitempty"`
	WBS2Level                      string  `json:"WBS2Level,omitempty"`
	ChargeType                     string  `json:"ChargeType,omitempty"`
	WBS2                           string  `json:"WBS2,omitempty"`
	WBS2Name                       string  `json:"WBS2Name,omitempty"`
	WBS3Level                      string  `json:"WBS3Level,omitempty"`
	WBS3                           string  `json:"WBS3,omitempty"`
	WBS3Name                       string  `json:"WBS3Name,omitempty"`
	Org                            string  `json:"Org,omitempty"`
	TransDate                      string  `json:"TransDate,omitempty"`
	DueDate                        string  `json:"DueDate,omitempty"`
	TransComment                   string  `json:"TransComment,omitempty"`
	Posted                         string  `json:"Posted,omitempty"`
	Seq                            int     `json:"Seq,omitempty"`
	CurrencyCode                   string  `json:"CurrencyCode,omitempty"`
	CurrencyExchangeOverrideMethod string  `json:"CurrencyExchangeOverrideMethod,omitempty"`
	CurrencyExchangeOverrideDate   string  `json:"CurrencyExchangeOverrideDate,omitempty"`
	CurrencyExchangeOverrideRate   float64 `json:"CurrencyExchangeOverrideRate,omitempty"`
	Status                         string  `json:"Status,omitempty"`
	StatusDescription              string  `json:"StatusDescription,omitempty"`
	AuthorizedByName               string  `json:"AuthorizedByName,omitempty"`
	AuthorizedBy                   string  `json:"AuthorizedBy,omitempty"`
	RejectReason                   string  `json:"RejectReason,omitempty"`
	ModUser                        string  `json:"ModUser,omitempty"`
	ModUserEmployee                string  `json:"ModUserEmployee,omitempty"`
	ModDate                        string  `json:"ModDate,omitempty"`
	Diary                          string  `json:"Diary,omitempty"`
	DiaryNo                        int     `json:"DiaryNo,omitempty"`
}

// InvoiceDetail represents an invoice detail line.
type InvoiceDetail struct {
	Batch               string  `json:"Batch,omitempty"`
	Invoice             string  `json:"Invoice,omitempty"`
	WBS1                string  `json:"WBS1,omitempty"`
	WBS2                string  `json:"WBS2,omitempty"`
	WBS3                string  `json:"WBS3,omitempty"`
	PKey                string  `json:"PKey,omitempty"`
	Seq                 int     `json:"Seq,omitempty"`
	InvoiceSection      string  `json:"InvoiceSection,omitempty"`
	TaxCode             string  `json:"TaxCode,omitempty"`
	Account             string  `json:"Account,omitempty"`
	AccountName         string  `json:"AccountName,omitempty"`
	AccountCurrencyCode string  `json:"AccountCurrencyCode,omitempty"`
	GlobalAccount       string  `json:"GlobalAccount,omitempty"`
	AccountType         string  `json:"AccountType,omitempty"`
	Amount              float64 `json:"Amount,omitempty"`
	RetAmount           float64 `json:"RetAmount,omitempty"`
	TaxBasis            float64 `json:"TaxBasis,omitempty"`
	Retainer            string  `json:"Retainer,omitempty"`
	LinkWBS1            string  `json:"LinkWBS1,omitempty"`
	LinkWBS2            string  `json:"LinkWBS2,omitempty"`
	LinkWBS3            string  `json:"LinkWBS3,omitempty"`
	CreditMemoRefNo     string  `json:"CreditMemoRefNo,omitempty"`
}

// ListInvoiceBatches retrieves a list of invoice batch summaries from the
// inLanding endpoint. Uses GET /DataEntry/inLanding.
func (c *Client) ListInvoiceBatches(ctx context.Context, q *Query) ([]InvoiceControl, error) {
	var results []InvoiceControl
	if err := c.get(ctx, "DataEntry/inLanding", q, &results); err != nil {
		return nil, fmt.Errorf("listing invoice batches: %w", err)
	}
	return results, nil
}

// GetInvoiceControl retrieves the control record for an invoice batch.
// Uses GET /DataEntry/inControl/{batch}.
func (c *Client) GetInvoiceControl(ctx context.Context, batch string, q *Query) ([]InvoiceControl, error) {
	var results []InvoiceControl
	if err := c.get(ctx, "DataEntry/inControl/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice control for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetInvoiceMasterRecords retrieves the master records for an invoice batch.
// Uses GET /DataEntry/inMaster/{batch}.
func (c *Client) GetInvoiceMasterRecords(ctx context.Context, batch string, q *Query) ([]InvoiceMaster, error) {
	var results []InvoiceMaster
	if err := c.get(ctx, "DataEntry/inMaster/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice master records for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetInvoiceDetailRecords retrieves the detail lines for an invoice batch.
// Uses GET /DataEntry/inDetail/{batch}.
func (c *Client) GetInvoiceDetailRecords(ctx context.Context, batch string, q *Query) ([]InvoiceDetail, error) {
	var results []InvoiceDetail
	if err := c.get(ctx, "DataEntry/inDetail/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice detail records for batch %s: %w", batch, err)
	}
	return results, nil
}

// CreateInvoiceBatch creates a new invoice batch via the inControl endpoint.
// Uses POST /DataEntry/inControl.
func (c *Client) CreateInvoiceBatch(ctx context.Context, control *InvoiceControl) ([]InvoiceControl, error) {
	var results []InvoiceControl
	if err := c.post(ctx, "DataEntry/inControl", control, &results); err != nil {
		return nil, fmt.Errorf("creating invoice batch: %w", err)
	}
	return results, nil
}

// UpdateInvoiceBatch updates an existing invoice batch control record.
// Uses PUT /DataEntry/inControl/{batch}.
func (c *Client) UpdateInvoiceBatch(ctx context.Context, batch string, control *InvoiceControl) ([]InvoiceControl, error) {
	var results []InvoiceControl
	if err := c.put(ctx, "DataEntry/inControl/"+batch, control, &results); err != nil {
		return nil, fmt.Errorf("updating invoice batch %s: %w", batch, err)
	}
	return results, nil
}

// DeleteInvoiceBatch deletes an invoice batch.
// Uses DELETE /DataEntry/inControl/{batch}.
func (c *Client) DeleteInvoiceBatch(ctx context.Context, batch string) error {
	if err := c.delete(ctx, "DataEntry/inControl/"+batch); err != nil {
		return fmt.Errorf("deleting invoice batch %s: %w", batch, err)
	}
	return nil
}
