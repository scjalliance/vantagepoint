package vantagepoint

import (
	"context"
	"fmt"
)

// Interactive Billing in Vantagepoint lives under /InteractiveDetail/ with
// specialized read-only endpoints for billing data. There is no generic
// /billing endpoint.

// BillingUnit represents a billing unit record from interactive billing.
type BillingUnit map[string]any

// BillingExpense represents a billing expense record from interactive billing.
type BillingExpense map[string]any

// BillingLimit represents billing limits for a project.
type BillingLimit map[string]any

// InvoiceHeader represents an invoice header from interactive billing.
type InvoiceHeader map[string]any

// InvoiceLaborItem represents a labor line item on an interactive billing invoice.
type InvoiceLaborItem map[string]any

// InvoiceExpenseItem represents an expense line item on an interactive billing invoice.
type InvoiceExpenseItem map[string]any

// InvoiceUnitItem represents a unit line item on an interactive billing invoice.
type InvoiceUnitItem map[string]any

// InvoiceConsultantItem represents a consultant line item on an interactive billing invoice.
type InvoiceConsultantItem map[string]any

// InvoiceAddOn represents an add-on item on an interactive billing invoice.
type InvoiceAddOn map[string]any

// InvoiceFixedFeeItem represents a fixed fee item on an interactive billing invoice.
type InvoiceFixedFeeItem map[string]any

// InvoiceInterestItem represents an interest item on an interactive billing invoice.
type InvoiceInterestItem map[string]any

// GetBillingUnits retrieves billing units for a project.
// Uses GET /InteractiveDetail/{wbs1}/Units.
func (c *Client) GetBillingUnits(ctx context.Context, wbs1 string, q *Query) ([]BillingUnit, error) {
	var results []BillingUnit
	if err := c.get(ctx, "InteractiveDetail/"+wbs1+"/Units", q, &results); err != nil {
		return nil, fmt.Errorf("getting billing units for %s: %w", wbs1, err)
	}
	return results, nil
}

// GetBillingExpenses retrieves billing expenses for a project.
// Uses GET /InteractiveDetail/{wbs1}/Expense.
func (c *Client) GetBillingExpenses(ctx context.Context, wbs1 string, q *Query) ([]BillingExpense, error) {
	var results []BillingExpense
	if err := c.get(ctx, "InteractiveDetail/"+wbs1+"/Expense", q, &results); err != nil {
		return nil, fmt.Errorf("getting billing expenses for %s: %w", wbs1, err)
	}
	return results, nil
}

// GetBillingLimits retrieves billing limits for a WBS key.
// Uses GET /InteractiveDetail/{wbsKey}/BillingLimits.
func (c *Client) GetBillingLimits(ctx context.Context, wbsKey string, q *Query) ([]BillingLimit, error) {
	var results []BillingLimit
	if err := c.get(ctx, "InteractiveDetail/"+wbsKey+"/BillingLimits", q, &results); err != nil {
		return nil, fmt.Errorf("getting billing limits for %s: %w", wbsKey, err)
	}
	return results, nil
}

// GetInvoiceHeaders retrieves the invoice header (master) records for a project.
// Uses GET /InteractiveDetail/{wbs1}/InvoiceMaster.
func (c *Client) GetInvoiceHeaders(ctx context.Context, wbs1 string, q *Query) ([]InvoiceHeader, error) {
	var results []InvoiceHeader
	if err := c.get(ctx, "InteractiveDetail/"+wbs1+"/InvoiceMaster", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice headers for %s: %w", wbs1, err)
	}
	return results, nil
}

// GetInvoiceLaborItems retrieves invoice labor line items for a WBS key.
// Uses GET /InteractiveDetail/{wbsKey}/InvoiceLabor.
func (c *Client) GetInvoiceLaborItems(ctx context.Context, wbsKey string, q *Query) ([]InvoiceLaborItem, error) {
	var results []InvoiceLaborItem
	if err := c.get(ctx, "InteractiveDetail/"+wbsKey+"/InvoiceLabor", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice labor items for %s: %w", wbsKey, err)
	}
	return results, nil
}

// GetInvoiceExpenseItems retrieves invoice expense line items for a WBS key.
// Uses GET /InteractiveDetail/{wbsKey}/InvoiceExpenses.
func (c *Client) GetInvoiceExpenseItems(ctx context.Context, wbsKey string, q *Query) ([]InvoiceExpenseItem, error) {
	var results []InvoiceExpenseItem
	if err := c.get(ctx, "InteractiveDetail/"+wbsKey+"/InvoiceExpenses", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice expense items for %s: %w", wbsKey, err)
	}
	return results, nil
}

// GetInvoiceUnitItems retrieves invoice unit line items for a WBS key.
// Uses GET /InteractiveDetail/{wbsKey}/InvoiceUnit.
func (c *Client) GetInvoiceUnitItems(ctx context.Context, wbsKey string, q *Query) ([]InvoiceUnitItem, error) {
	var results []InvoiceUnitItem
	if err := c.get(ctx, "InteractiveDetail/"+wbsKey+"/InvoiceUnit", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice unit items for %s: %w", wbsKey, err)
	}
	return results, nil
}

// GetInvoiceConsultantItems retrieves invoice consultant line items for a WBS key.
// Uses GET /InteractiveDetail/{wbsKey}/InvoiceConsultant.
func (c *Client) GetInvoiceConsultantItems(ctx context.Context, wbsKey string, q *Query) ([]InvoiceConsultantItem, error) {
	var results []InvoiceConsultantItem
	if err := c.get(ctx, "InteractiveDetail/"+wbsKey+"/InvoiceConsultant", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice consultant items for %s: %w", wbsKey, err)
	}
	return results, nil
}

// GetInvoiceAddOns retrieves invoice add-on items for a project.
// Uses GET /InteractiveDetail/{wbs1}/InvoiceAddOns.
func (c *Client) GetInvoiceAddOns(ctx context.Context, wbs1 string, q *Query) ([]InvoiceAddOn, error) {
	var results []InvoiceAddOn
	if err := c.get(ctx, "InteractiveDetail/"+wbs1+"/InvoiceAddOns", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice add-ons for %s: %w", wbs1, err)
	}
	return results, nil
}

// GetInvoiceFixedFeeItems retrieves invoice fixed fee items for a project.
// Uses GET /InteractiveDetail/{wbs1}/InvoiceFixedFee.
func (c *Client) GetInvoiceFixedFeeItems(ctx context.Context, wbs1 string, q *Query) ([]InvoiceFixedFeeItem, error) {
	var results []InvoiceFixedFeeItem
	if err := c.get(ctx, "InteractiveDetail/"+wbs1+"/InvoiceFixedFee", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice fixed fee items for %s: %w", wbs1, err)
	}
	return results, nil
}

// GetInvoiceInterestItems retrieves invoice interest items for a project.
// Uses GET /InteractiveDetail/{wbs1}/InvoiceInterest.
func (c *Client) GetInvoiceInterestItems(ctx context.Context, wbs1 string, q *Query) ([]InvoiceInterestItem, error) {
	var results []InvoiceInterestItem
	if err := c.get(ctx, "InteractiveDetail/"+wbs1+"/InvoiceInterest", q, &results); err != nil {
		return nil, fmt.Errorf("getting invoice interest items for %s: %w", wbs1, err)
	}
	return results, nil
}
