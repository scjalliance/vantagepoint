package vantagepoint

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// ARBalance represents an accounts receivable balance record from the
// ARReview/ARBalance endpoint, containing invoice-level balance and billing
// client information for a project.
type ARBalance struct {
	WBS1            string  `json:"WBS1,omitempty"`
	WBS1Name        string  `json:"WBS1Name,omitempty"`
	Company         string  `json:"Company,omitempty"`
	Invoice         string  `json:"Invoice,omitempty"`
	InvoiceDate     string  `json:"InvoiceDate,omitempty"`
	CreditMemoRefNo string  `json:"CreditMemoRefNo,omitempty"`
	Balance         float64 `json:"Balance,omitempty"`
	OriginalAmt     float64 `json:"OriginalAmt,omitempty"`
	BillingClientID string  `json:"BillingClientID,omitempty"`
	BillingClient   string  `json:"BillingClient,omitempty"`
}

// ARBalanceDetail represents a detailed breakdown of an AR balance from the
// ARReview/ARBalanceDetail endpoint, including WBS hierarchy, account, and
// billing group information.
type ARBalanceDetail struct {
	Invoice          string  `json:"Invoice,omitempty"`
	CreditMemoRefNo  string  `json:"CreditMemoRefNo,omitempty"`
	WBS1             string  `json:"WBS1,omitempty"`
	WBS2             string  `json:"WBS2,omitempty"`
	WBS3             string  `json:"WBS3,omitempty"`
	Account          string  `json:"Account,omitempty"`
	WBS1Name         string  `json:"WBS1Name,omitempty"`
	WBS2Name         string  `json:"WBS2Name,omitempty"`
	WBS3Name         string  `json:"WBS3Name,omitempty"`
	BillClientID     string  `json:"BillClientID,omitempty"`
	BillClient       string  `json:"BillClient,omitempty"`
	BillingGroupWBS1 string  `json:"BillingGroupWBS1,omitempty"`
	BillingGroupName string  `json:"BillingGroupName,omitempty"`
	InvoiceAmount    float64 `json:"InvoiceAmount,omitempty"`
}

// ARBalanceDetailCR represents a cash receipt detail record from the
// ARReview/ARBalanceDetailCR endpoint, containing payment and retainer
// information applied against an AR balance.
type ARBalanceDetailCR struct {
	WBS1            string  `json:"WBS1,omitempty"`
	WBS2            string  `json:"WBS2,omitempty"`
	WBS3            string  `json:"WBS3,omitempty"`
	Invoice         string  `json:"Invoice,omitempty"`
	CashReceiptDate string  `json:"CashReceiptDate,omitempty"`
	Account         string  `json:"Account,omitempty"`
	TransType       string  `json:"TransType,omitempty"`
	SubType         string  `json:"SubType,omitempty"`
	InvoiceSection  string  `json:"InvoiceSection,omitempty"`
	Amount          float64 `json:"Amount,omitempty"`
	RetainerAmount  float64 `json:"RetainerAmount,omitempty"`
	WBS1Name        string  `json:"WBS1Name,omitempty"`
	WBS2Name        string  `json:"WBS2Name,omitempty"`
	WBS3Name        string  `json:"WBS3Name,omitempty"`
}

// ListARBalances retrieves accounts receivable balance records from the
// ARReview/ARBalance endpoint. Use the Query parameter to filter, limit,
// and paginate results.
func (c *Client) ListARBalances(ctx context.Context, q *Query) ([]ARBalance, error) {
	var results []ARBalance
	if err := c.get(ctx, "ARReview/ARBalance", q, &results); err != nil {
		return nil, fmt.Errorf("listing AR balances: %w", err)
	}
	return results, nil
}

// ListARBalanceDetails retrieves detailed AR balance breakdowns from the
// ARReview/ARBalanceDetail endpoint for the specified WBS1/invoice keys.
// Keys are passed as a comma-separated ARWBS1InvoiceKey query parameter.
func (c *Client) ListARBalanceDetails(ctx context.Context, keys []string) ([]ARBalanceDetail, error) {
	vals := url.Values{}
	vals.Set("ARWBS1InvoiceKey", strings.Join(keys, ","))
	var results []ARBalanceDetail
	if err := c.do(ctx, "GET", "ARReview/ARBalanceDetail", vals, nil, &results); err != nil {
		return nil, fmt.Errorf("listing AR balance details: %w", err)
	}
	return results, nil
}

// ListARBalanceDetailCRs retrieves cash receipt detail records from the
// ARReview/ARBalanceDetailCR endpoint for the specified WBS1/invoice keys.
// Keys are passed as a comma-separated ARWBS1InvoiceKey query parameter.
func (c *Client) ListARBalanceDetailCRs(ctx context.Context, keys []string) ([]ARBalanceDetailCR, error) {
	vals := url.Values{}
	vals.Set("ARWBS1InvoiceKey", strings.Join(keys, ","))
	var results []ARBalanceDetailCR
	if err := c.do(ctx, "GET", "ARReview/ARBalanceDetailCR", vals, nil, &results); err != nil {
		return nil, fmt.Errorf("listing AR balance detail CRs: %w", err)
	}
	return results, nil
}
