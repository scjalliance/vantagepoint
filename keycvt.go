package vantagepoint

import (
	"context"
	"fmt"
)

// KeyFormat represents the key conversion format configuration from the
// KeyCvt/CFGFormat endpoint. It defines length, delimiter, and formatting
// rules for all entity key types (WBS, Employee, Account, Client, etc.).
type KeyFormat struct {
	PKey                      string `json:"PKey,omitempty"`
	Entity                    string `json:"Entity,omitempty"`
	WBS1Length                int    `json:"WBS1Length,omitempty"`
	WBS1LeadZeros             string `json:"WBS1LeadZeros,omitempty"`
	WBS1Delimiter             string `json:"WBS1Delimiter,omitempty"`
	WBS1DelimiterPosition     int    `json:"WBS1DelimiterPosition,omitempty"`
	WBS1ChangeSide            string `json:"WBS1ChangeSide,omitempty"`
	WBS1Label                 string `json:"WBS1Label,omitempty"`
	WBS2Length                int    `json:"WBS2Length,omitempty"`
	WBS2LeadZeros             string `json:"WBS2LeadZeros,omitempty"`
	WBS2Delimiter             string `json:"WBS2Delimiter,omitempty"`
	WBS2DelimiterPosition     int    `json:"WBS2DelimiterPosition,omitempty"`
	WBS2ChangeSide            string `json:"WBS2ChangeSide,omitempty"`
	WBS2Label                 string `json:"WBS2Label,omitempty"`
	WBS3Length                int    `json:"WBS3Length,omitempty"`
	WBS3LeadZeros             string `json:"WBS3LeadZeros,omitempty"`
	WBS3Delimiter             string `json:"WBS3Delimiter,omitempty"`
	WBS3DelimiterPosition     int    `json:"WBS3DelimiterPosition,omitempty"`
	WBS3ChangeSide            string `json:"WBS3ChangeSide,omitempty"`
	WBS3Label                 string `json:"WBS3Label,omitempty"`
	WBS4Length                int    `json:"WBS4Length,omitempty"`
	EmployeeLength            int    `json:"EmployeeLength,omitempty"`
	EmployeeLeadZeros         string `json:"EmployeeLeadZeros,omitempty"`
	EmployeeDelimiter         string `json:"EmployeeDelimiter,omitempty"`
	EmployeeDelimiterPosition int    `json:"EmployeeDelimiterPosition,omitempty"`
	AccountLength             int    `json:"AccountLength,omitempty"`
	AccountLeadZeros          string `json:"AccountLeadZeros,omitempty"`
	AccountDelimiter          string `json:"AccountDelimiter,omitempty"`
	AccountDelimiterPosition  int    `json:"AccountDelimiterPosition,omitempty"`
	ClientLength              int    `json:"ClientLength,omitempty"`
	ClientLeadZeros           string `json:"ClientLeadZeros,omitempty"`
	ClientDelimiter           string `json:"ClientDelimiter,omitempty"`
	ClientDelimiterPosition   int    `json:"ClientDelimiterPosition,omitempty"`
	VendorLength              int    `json:"VendorLength,omitempty"`
	VendorLeadZeros           string `json:"VendorLeadZeros,omitempty"`
	VendorDelimiter           string `json:"VendorDelimiter,omitempty"`
	VendorDelimiterPosition   int    `json:"VendorDelimiterPosition,omitempty"`
	UnitLength                int    `json:"UnitLength,omitempty"`
	UnitLeadZeros             string `json:"UnitLeadZeros,omitempty"`
	UnitDelimiter             string `json:"UnitDelimiter,omitempty"`
	UnitDelimiterPosition     int    `json:"UnitDelimiterPosition,omitempty"`
	RefnoLength               int    `json:"RefnoLength,omitempty"`
	OrgLevels                 int    `json:"OrgLevels,omitempty"`
	OrgDelimiter              string `json:"OrgDelimiter,omitempty"`
	LCLevels                  int    `json:"LCLevels,omitempty"`
	LCDelimiter               string `json:"LCDelimiter,omitempty"`
	ChangeSide                string `json:"ChangeSide,omitempty"`
}

// OrgLevel represents a single organization level definition from the
// KeyCvt/CFGFormatOrg endpoint. Each level has a label, start position,
// and length within the org key string.
type OrgLevel struct {
	CLevel int    `json:"cLevel,omitempty"`
	Label  string `json:"label,omitempty"`
	Start  int    `json:"Start,omitempty"`
	Length int    `json:"Length,omitempty"`
}

// GetKeyFormat retrieves the key conversion format configuration that defines
// how entity keys (WBS, Employee, Account, etc.) are structured.
// Uses GET /KeyCvt/CFGFormat.
func (c *Client) GetKeyFormat(ctx context.Context) ([]KeyFormat, error) {
	var results []KeyFormat
	if err := c.get(ctx, "KeyCvt/CFGFormat", nil, &results); err != nil {
		return nil, fmt.Errorf("getting key format: %w", err)
	}
	return results, nil
}

// GetOrgLevels retrieves the organization level definitions that describe
// how org keys are segmented (e.g., Company, Office, Discipline).
// Uses GET /KeyCvt/CFGFormatOrg.
func (c *Client) GetOrgLevels(ctx context.Context) ([]OrgLevel, error) {
	var results []OrgLevel
	if err := c.get(ctx, "KeyCvt/CFGFormatOrg", nil, &results); err != nil {
		return nil, fmt.Errorf("getting org levels: %w", err)
	}
	return results, nil
}
