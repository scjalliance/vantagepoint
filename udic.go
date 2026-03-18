package vantagepoint

import (
	"context"
	"fmt"
)

// UDICRecord represents a user-defined information center (hub) record.
// UDIC records are highly configurable — the Cust* fields are user-defined
// and vary by InfocenterArea. The core fields (UDIC_UID, UDICNumber,
// UDICName, CreateUser, etc.) are always present.
type UDICRecord struct {
	UDIC_UID               string `json:"UDIC_UID,omitempty"`
	DescCustomCurrencyCode string `json:"desc_CustomCurrencyCode,omitempty"`
	CustNumber             string `json:"CustNumber,omitempty"`
	CustName               string `json:"CustName,omitempty"`
	CreateUser             string `json:"CreateUser,omitempty"`
	CreateDate             string `json:"CreateDate,omitempty"`
	ModUser                string `json:"ModUser,omitempty"`
	ModDate                string `json:"ModDate,omitempty"`
	HasPhoto               int    `json:"HasPhoto,omitempty"`
	PhotoModDate           string `json:"PhotoModDate,omitempty"`
	UDICNumber             string `json:"UDICNumber,omitempty"`
	UDICName               string `json:"UDICName,omitempty"`
}

// UDICFile represents a supporting document associated with a UDIC record.
type UDICFile struct {
	FileID              string `json:"FileID,omitempty"`
	PKey                string `json:"PKey,omitempty"`
	FileName            string `json:"FileName,omitempty"`
	FileDescription     string `json:"FileDescription,omitempty"`
	FileSize            int    `json:"FileSize,omitempty"`
	ContentType         string `json:"ContentType,omitempty"`
	Key1                string `json:"Key1,omitempty"`
	Key2                string `json:"Key2,omitempty"`
	Key3                string `json:"Key3,omitempty"`
	UDIC_UID            string `json:"UDIC_UID,omitempty"`
	CategoryCode        string `json:"CategoryCode,omitempty"`
	Application         string `json:"Application,omitempty"`
	CategoryDescription string `json:"CategoryDescription,omitempty"`
}

// ListUDICRecords retrieves a list of UDIC records for the given infocenter area.
// Uses GET /UDIC/{infocenterArea}.
func (c *Client) ListUDICRecords(ctx context.Context, infocenterArea string, q *Query) ([]UDICRecord, error) {
	var results []UDICRecord
	if err := c.get(ctx, "UDIC/"+infocenterArea, q, &results); err != nil {
		return nil, fmt.Errorf("listing UDIC records for %s: %w", infocenterArea, err)
	}
	return results, nil
}

// GetUDICRecord retrieves a single UDIC record by its area and UID.
// Uses GET /UDIC/{infocenterArea}/{udicID}.
func (c *Client) GetUDICRecord(ctx context.Context, infocenterArea, udicID string) (*UDICRecord, error) {
	var results []UDICRecord
	if err := c.get(ctx, "UDIC/"+infocenterArea+"/"+udicID, nil, &results); err != nil {
		return nil, fmt.Errorf("getting UDIC record %s/%s: %w", infocenterArea, udicID, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("getting UDIC record %s/%s: %w", infocenterArea, udicID, ErrNotFound)
	}
	return &results[0], nil
}

// CreateUDICRecord creates a new UDIC record in the given infocenter area.
// Uses POST /UDIC/{infocenterArea}.
func (c *Client) CreateUDICRecord(ctx context.Context, infocenterArea string, record *UDICRecord) (*UDICRecord, error) {
	var results []UDICRecord
	if err := c.post(ctx, "UDIC/"+infocenterArea, record, &results); err != nil {
		return nil, fmt.Errorf("creating UDIC record in %s: %w", infocenterArea, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("creating UDIC record in %s: %w", infocenterArea, ErrNotFound)
	}
	return &results[0], nil
}

// UpdateUDICRecord updates an existing UDIC record.
// Uses PUT /UDIC/{infocenterArea}/{udicUID}.
func (c *Client) UpdateUDICRecord(ctx context.Context, infocenterArea, udicUID string, record *UDICRecord) (*UDICRecord, error) {
	var results []UDICRecord
	if err := c.put(ctx, "UDIC/"+infocenterArea+"/"+udicUID, record, &results); err != nil {
		return nil, fmt.Errorf("updating UDIC record %s/%s: %w", infocenterArea, udicUID, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("updating UDIC record %s/%s: %w", infocenterArea, udicUID, ErrNotFound)
	}
	return &results[0], nil
}

// DeleteUDICRecord deletes a UDIC record.
// Uses DELETE /UDIC/{infocenterArea}/{udicUID}.
func (c *Client) DeleteUDICRecord(ctx context.Context, infocenterArea, udicUID string) error {
	if err := c.delete(ctx, "UDIC/"+infocenterArea+"/"+udicUID); err != nil {
		return fmt.Errorf("deleting UDIC record %s/%s: %w", infocenterArea, udicUID, err)
	}
	return nil
}

// ListUDICFiles retrieves supporting documents for a UDIC record.
// Uses GET /UDIC/{infocenterArea}/{udicUID}/files.
func (c *Client) ListUDICFiles(ctx context.Context, infocenterArea, udicUID string, q *Query) ([]UDICFile, error) {
	var results []UDICFile
	if err := c.get(ctx, "UDIC/"+infocenterArea+"/"+udicUID+"/files", q, &results); err != nil {
		return nil, fmt.Errorf("listing UDIC files for %s/%s: %w", infocenterArea, udicUID, err)
	}
	return results, nil
}

// ListUDICCustomTable retrieves custom table data for a UDIC record.
// Uses GET /UDIC/{infocenterArea}/{udicUID}/CustomTable/{customTable}.
func (c *Client) ListUDICCustomTable(ctx context.Context, infocenterArea, udicUID, customTable string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "UDIC/"+infocenterArea+"/"+udicUID+"/CustomTable/"+customTable, q, &results); err != nil {
		return nil, fmt.Errorf("listing UDIC custom table %s for %s/%s: %w", customTable, infocenterArea, udicUID, err)
	}
	return results, nil
}
