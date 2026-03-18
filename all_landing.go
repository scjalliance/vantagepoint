package vantagepoint

import (
	"context"
	"fmt"
)

// AllLanding represents a cross-transaction-type batch record from the
// Vantagepoint allLanding endpoint. It provides a unified view of batches
// across different transaction types (e.g., AP, AR, timesheet, expense)
// with their posting status, approval state, and scheduling metadata.
type AllLanding struct {
	TransType           string  `json:"TransType,omitempty"`
	TransTypeDesc       string  `json:"TransTypeDesc,omitempty"`
	Batch               string  `json:"Batch,omitempty"`
	Company             string  `json:"Company,omitempty"`
	Period              string  `json:"Period,omitempty"`
	RawPeriod           int     `json:"RawPeriod,omitempty"`
	TransDate           string  `json:"TransDate,omitempty"`
	RefNo               string  `json:"RefNo,omitempty"`
	Creator             string  `json:"Creator,omitempty"`
	Description         string  `json:"Description,omitempty"`
	EndDate             string  `json:"EndDate,omitempty"`
	Total               float64 `json:"Total,omitempty"`
	CurrencyCode        string  `json:"CurrencyCode,omitempty"`
	ApprovalStatus      string  `json:"ApprovalStatus,omitempty"`
	ApprovalStatusDesc  string  `json:"ApprovalStatusDesc,omitempty"`
	PostStatus          string  `json:"PostStatus,omitempty"`
	PostStatusDesc      string  `json:"PostStatusDesc,omitempty"`
	PostComment         string  `json:"PostComment,omitempty"`
	PostDate            string  `json:"PostDate,omitempty"`
	PostUser            string  `json:"PostUser,omitempty"`
	PostSeq             int     `json:"PostSeq,omitempty"`
	TransListReportPath string  `json:"TransListReportPath,omitempty"`
	PostLogReportPath   string  `json:"PostLogReportPath,omitempty"`
	PostLogReportName   string  `json:"PostLogReportName,omitempty"`
	Diary               string  `json:"Diary,omitempty"`
	DiaryNo             int     `json:"DiaryNo,omitempty"`
	RevGen              string  `json:"RevGen,omitempty"`
	Posted              string  `json:"Posted,omitempty"`
	AutoPost            string  `json:"AutoPost,omitempty"`
	RecurSchedule       string  `json:"RecurSchedule,omitempty"`
	Reverse             string  `json:"Reverse,omitempty"`
	Reversed            string  `json:"Reversed,omitempty"`
	Recurring           string  `json:"Recurring,omitempty"`
	EntryAccess         string  `json:"EntryAccess,omitempty"`
	ListsAccess         string  `json:"ListsAccess,omitempty"`
	PostAccess          string  `json:"PostAccess,omitempty"`
}

// ListAllBatches retrieves a list of batches across all transaction types
// from the Vantagepoint allLanding endpoint. Pass a *Query to filter,
// paginate, or sort results; nil retrieves defaults.
func (c *Client) ListAllBatches(ctx context.Context, q *Query) ([]AllLanding, error) {
	var batches []AllLanding
	if err := c.get(ctx, "DataEntry/allLanding", q, &batches); err != nil {
		return nil, fmt.Errorf("listing all batches: %w", err)
	}
	return batches, nil
}
