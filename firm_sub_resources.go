package vantagepoint

import (
	"context"
	"fmt"
)

// FirmAddress represents an address record associated with a firm.
type FirmAddress struct {
	ClientID              string `json:"ClientID,omitempty"`
	CLAddressID           string `json:"CLAddressID,omitempty"`
	PrimaryInd            string `json:"PrimaryInd,omitempty"`
	Billing               string `json:"Billing,omitempty"`
	Accounting            string `json:"Accounting,omitempty"`
	Address               string `json:"Address,omitempty"`
	Addressee             string `json:"Addressee,omitempty"`
	Address1              string `json:"Address1,omitempty"`
	Address2              string `json:"Address2,omitempty"`
	Address3              string `json:"Address3,omitempty"`
	Address4              string `json:"Address4,omitempty"`
	City                  string `json:"City,omitempty"`
	State                 string `json:"State,omitempty"`
	Zip                   string `json:"Zip,omitempty"`
	Country               string `json:"Country,omitempty"`
	Phone                 string `json:"Phone,omitempty"`
	FAX                   string `json:"FAX,omitempty"`
	Email                 string `json:"Email,omitempty"`
	TaxRegistrationNumber string `json:"TaxRegistrationNumber,omitempty"`
	TaxCountryCode        string `json:"TaxCountryCode,omitempty"`
	PhoneFormat           string `json:"PhoneFormat,omitempty"`
	FaxFormat             string `json:"FaxFormat,omitempty"`
	Payment               string `json:"Payment,omitempty"`
	StateDescription      string `json:"StateDescription,omitempty"`
	CreateUser            string `json:"CreateUser,omitempty"`
	CreateDate            string `json:"CreateDate,omitempty"`
	ModUser               string `json:"ModUser,omitempty"`
	ModDate               string `json:"ModDate,omitempty"`
}

// FirmAlias represents an alias record for a firm.
type FirmAlias struct {
	PKey       string `json:"PKey,omitempty"`
	ClientID   string `json:"ClientID,omitempty"`
	Alias      string `json:"Alias,omitempty"`
	CreateUser string `json:"CreateUser,omitempty"`
	CreateDate string `json:"CreateDate,omitempty"`
	ModUser    string `json:"ModUser,omitempty"`
	ModDate    string `json:"ModDate,omitempty"`
}

// FirmLink represents a link associated with a firm.
type FirmLink struct {
	Description string `json:"Description,omitempty"`
	FilePath    string `json:"FilePath,omitempty"`
	Graphic     string `json:"Graphic,omitempty"`
}

// ListFirmAddresses retrieves address records for a firm.
// Uses GET /firm/{clientID}/address.
func (c *Client) ListFirmAddresses(ctx context.Context, clientID string, q *Query) ([]FirmAddress, error) {
	var results []FirmAddress
	if err := c.get(ctx, "firm/"+clientID+"/address", q, &results); err != nil {
		return nil, fmt.Errorf("listing firm addresses for %s: %w", clientID, err)
	}
	return results, nil
}

// CreateFirmAddress creates a new address for a firm.
// Uses POST /firm/{clientID}/address.
func (c *Client) CreateFirmAddress(ctx context.Context, clientID string, addr *FirmAddress) (*FirmAddress, error) {
	var results []FirmAddress
	if err := c.post(ctx, "firm/"+clientID+"/address", addr, &results); err != nil {
		return nil, fmt.Errorf("creating firm address for %s: %w", clientID, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("creating firm address for %s: %w", clientID, ErrNotFound)
	}
	return &results[0], nil
}

// ListFirmAliases retrieves alias records for a firm.
// Uses GET /firm/{clientID}/alias.
func (c *Client) ListFirmAliases(ctx context.Context, clientID string, q *Query) ([]FirmAlias, error) {
	var results []FirmAlias
	if err := c.get(ctx, "firm/"+clientID+"/alias", q, &results); err != nil {
		return nil, fmt.Errorf("listing firm aliases for %s: %w", clientID, err)
	}
	return results, nil
}

// ListFirmLinks retrieves links associated with a firm.
// Uses GET /firm/{clientID}/links.
func (c *Client) ListFirmLinks(ctx context.Context, clientID string, q *Query) ([]FirmLink, error) {
	var results []FirmLink
	if err := c.get(ctx, "firm/"+clientID+"/links", q, &results); err != nil {
		return nil, fmt.Errorf("listing firm links for %s: %w", clientID, err)
	}
	return results, nil
}

// ListFirmFiles retrieves supporting documents for a firm.
// Uses GET /firm/{clientID}/files.
func (c *Client) ListFirmFiles(ctx context.Context, clientID string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "firm/"+clientID+"/files", q, &results); err != nil {
		return nil, fmt.Errorf("listing firm files for %s: %w", clientID, err)
	}
	return results, nil
}

// ListFirmCustomTable retrieves custom table data for a firm.
// Uses GET /firm/{clientID}/customTable/{customTable}.
func (c *Client) ListFirmCustomTable(ctx context.Context, clientID, customTable string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "firm/"+clientID+"/customTable/"+customTable, q, &results); err != nil {
		return nil, fmt.Errorf("listing firm custom table %s for %s: %w", customTable, clientID, err)
	}
	return results, nil
}
