package vantagepoint

import (
	"context"
	"fmt"
)

// TableMetadata represents metadata about a Vantagepoint API table.
type TableMetadata struct {
	TableName string          `json:"TableName,omitempty"`
	Fields    []FieldMetadata `json:"Fields,omitempty"`
}

// FieldMetadata represents metadata about a single field in a Vantagepoint table.
type FieldMetadata struct {
	FieldName  string `json:"FieldName,omitempty"`
	FieldType  string `json:"FieldType,omitempty"`
	MaxLength  int    `json:"MaxLength,omitempty"`
	IsRequired bool   `json:"IsRequired,omitempty"`
	IsReadOnly bool   `json:"IsReadOnly,omitempty"`
}

// GetMetadata retrieves metadata for a given API table, including field names,
// types, and constraints. The tableName parameter is case-sensitive for some
// resources (e.g., "APEntry", "CDEntry" use capital M in the path).
//
// Known metadata endpoints include: activity, boilerplate, campaign, contact,
// employee, firm, organization, project, APEntry, CDEntry, CREntry, CVEntry,
// EREntry, EXEntry, GLSummary, INEntry, JEEntry, LAEntry, MIEntry, PREntry,
// TSEntry, UNEntry, UPEntry.
func (c *Client) GetMetadata(ctx context.Context, tableName string) (*TableMetadata, error) {
	var meta TableMetadata
	if err := c.get(ctx, "metadata/"+tableName, nil, &meta); err != nil {
		return nil, fmt.Errorf("getting metadata for %s: %w", tableName, err)
	}
	return &meta, nil
}

// GetCodeTableMetadata retrieves metadata for a specific code table.
// Uses GET /metadata/codeTable/{codeTable}.
func (c *Client) GetCodeTableMetadata(ctx context.Context, codeTable string) (*TableMetadata, error) {
	var meta TableMetadata
	if err := c.get(ctx, "metadata/codeTable/"+codeTable, nil, &meta); err != nil {
		return nil, fmt.Errorf("getting code table metadata for %s: %w", codeTable, err)
	}
	return &meta, nil
}
