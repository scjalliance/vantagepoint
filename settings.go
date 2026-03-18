package vantagepoint

import (
	"context"
	"fmt"
)

// GeneralSettings represents firm-wide general configuration from the
// Settings/General endpoint.
type GeneralSettings struct {
	Company   string `json:"Company,omitempty"`
	FirmName  string `json:"FirmName,omitempty"`
	Byline    string `json:"Byline,omitempty"`
	Address1  string `json:"Address1,omitempty"`
	Address2  string `json:"Address2,omitempty"`
	Address3  string `json:"Address3,omitempty"`
	Address4  string `json:"Address4,omitempty"`
	Country   string `json:"Country,omitempty"`
	FirmEmail string `json:"FirmEmail,omitempty"`
	FirmPhone string `json:"FirmPhone,omitempty"`
}

// SystemLabel represents a single system label record from the
// Settings/SystemLabels endpoint. Labels control UI display names for
// entities, fields, and system terms.
type SystemLabel struct {
	UICultureName         string `json:"UICultureName,omitempty"`
	LabelName             string `json:"LabelName,omitempty"`
	LabelValue            string `json:"LabelValue,omitempty"`
	PlaceHolder           string `json:"PlaceHolder,omitempty"`
	Gender                string `json:"Gender,omitempty"`
	LeadingVowelTreatment string `json:"LeadingVowelTreatment,omitempty"`
	UDIC_ID               string `json:"UDIC_ID,omitempty"`
}

// Company represents a company record from the Settings/Company endpoint.
type Company struct {
	Company                string `json:"Company,omitempty"`
	FirmName               string `json:"FirmName,omitempty"`
	FunctionalCurrencyCode string `json:"FunctionalCurrencyCode,omitempty"`
	Country                string `json:"Country,omitempty"`
	Status                 string `json:"Status,omitempty"`
}

// AccountingPeriod represents an accounting period from the Settings/Period
// endpoint.
type AccountingPeriod struct {
	Period              int    `json:"Period,omitempty"`
	FmtPeriod           string `json:"FmtPeriod,omitempty"`
	AccountPdStart      string `json:"AccountPdStart,omitempty"`
	AccountPdEnd        string `json:"AccountPdEnd,omitempty"`
	FYStart             string `json:"FYStart,omitempty"`
	FYEnd               string `json:"FYEnd,omitempty"`
	ActiveCompanyClosed string `json:"ActiveCompanyClosed,omitempty"`
	AllCompaniesClosed  string `json:"AllCompaniesClosed,omitempty"`
	ApplytoAllCompanies string `json:"ApplytoAllCompanies,omitempty"`
}

// GetGeneralSettings retrieves firm-wide general configuration.
// Uses GET /Settings/General.
func (c *Client) GetGeneralSettings(ctx context.Context) (*GeneralSettings, error) {
	var results []GeneralSettings
	if err := c.get(ctx, "Settings/General", nil, &results); err != nil {
		return nil, fmt.Errorf("getting general settings: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("getting general settings: %w", ErrNotFound)
	}
	return &results[0], nil
}

// ListSystemLabels retrieves all system labels used for UI display names.
// Uses GET /Settings/SystemLabels.
func (c *Client) ListSystemLabels(ctx context.Context, q *Query) ([]SystemLabel, error) {
	var results []SystemLabel
	if err := c.get(ctx, "Settings/SystemLabels", q, &results); err != nil {
		return nil, fmt.Errorf("listing system labels: %w", err)
	}
	return results, nil
}

// ListCompanies retrieves all company records.
// Uses GET /Settings/Company.
func (c *Client) ListCompanies(ctx context.Context, q *Query) ([]Company, error) {
	var results []Company
	if err := c.get(ctx, "Settings/Company", q, &results); err != nil {
		return nil, fmt.Errorf("listing companies: %w", err)
	}
	return results, nil
}

// GetActiveCompany retrieves the currently active company.
// Uses GET /Settings/ActiveCompany.
func (c *Client) GetActiveCompany(ctx context.Context) (*Company, error) {
	var results []Company
	if err := c.get(ctx, "Settings/ActiveCompany", nil, &results); err != nil {
		return nil, fmt.Errorf("getting active company: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("getting active company: %w", ErrNotFound)
	}
	return &results[0], nil
}

// SetActiveCompany sets the active company by company code.
// Uses PUT /Settings/ActiveCompany/{company}.
func (c *Client) SetActiveCompany(ctx context.Context, company string) (*Company, error) {
	var results []Company
	if err := c.put(ctx, "Settings/ActiveCompany/"+company, nil, &results); err != nil {
		return nil, fmt.Errorf("setting active company to %s: %w", company, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("setting active company to %s: %w", company, ErrNotFound)
	}
	return &results[0], nil
}

// ListAccountingPeriods retrieves all accounting periods.
// Uses GET /Settings/Period.
func (c *Client) ListAccountingPeriods(ctx context.Context, q *Query) ([]AccountingPeriod, error) {
	var results []AccountingPeriod
	if err := c.get(ctx, "Settings/Period", q, &results); err != nil {
		return nil, fmt.Errorf("listing accounting periods: %w", err)
	}
	return results, nil
}

// GetActivePeriod retrieves the currently active accounting period.
// Uses GET /Settings/ActivePeriod.
func (c *Client) GetActivePeriod(ctx context.Context) (*AccountingPeriod, error) {
	var results []AccountingPeriod
	if err := c.get(ctx, "Settings/ActivePeriod", nil, &results); err != nil {
		return nil, fmt.Errorf("getting active period: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("getting active period: %w", ErrNotFound)
	}
	return &results[0], nil
}

// SetActivePeriod sets the active accounting period.
// Uses PUT /Settings/ActivePeriod/{period}.
func (c *Client) SetActivePeriod(ctx context.Context, period int) (*AccountingPeriod, error) {
	var results []AccountingPeriod
	if err := c.put(ctx, "Settings/ActivePeriod/"+fmt.Sprintf("%d", period), nil, &results); err != nil {
		return nil, fmt.Errorf("setting active period to %d: %w", period, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("setting active period to %d: %w", period, ErrNotFound)
	}
	return &results[0], nil
}
