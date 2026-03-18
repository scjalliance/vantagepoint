package vantagepoint

import (
	"context"
	"fmt"
)

// CodeTable represents a Vantagepoint code table entry.
type CodeTable struct {
	TableName   string `json:"TableName,omitempty"`
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
	Status      string `json:"Status,omitempty"`
}

// ListCodeTables retrieves a list of code table entries from the Vantagepoint API.
// The tableName parameter specifies which code table to query.
// Uses GET /codeTable/{codeTable}.
func (c *Client) ListCodeTables(ctx context.Context, tableName string, q *Query) ([]CodeTable, error) {
	var entries []CodeTable
	if err := c.get(ctx, "codeTable/"+tableName, q, &entries); err != nil {
		return nil, fmt.Errorf("listing code table %s: %w", tableName, err)
	}
	return entries, nil
}

// GetCodeTableEntry retrieves a single code table entry by table name and code.
// The Vantagepoint API returns an array even for single-resource lookups, so
// the response is decoded as a slice and the first element is returned.
// Uses GET /codeTable/{codeTable}/{code}.
func (c *Client) GetCodeTableEntry(ctx context.Context, tableName, code string) (*CodeTable, error) {
	var entries []CodeTable
	if err := c.get(ctx, "codeTable/"+tableName+"/"+code, nil, &entries); err != nil {
		return nil, fmt.Errorf("getting code table %s entry %s: %w", tableName, code, err)
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("getting code table %s entry %s: %w", tableName, code, ErrNotFound)
	}
	return &entries[0], nil
}

// CreateCodeTableEntry creates a new code table entry in Vantagepoint.
// Uses POST /codeTable.
func (c *Client) CreateCodeTableEntry(ctx context.Context, entry *CodeTable) (*CodeTable, error) {
	var result CodeTable
	if err := c.post(ctx, "codeTable", entry, &result); err != nil {
		return nil, fmt.Errorf("creating code table entry: %w", err)
	}
	return &result, nil
}

// UpdateCodeTableEntry updates an existing code table entry.
// Uses PUT /codeTable/{codeTable}/{code}.
func (c *Client) UpdateCodeTableEntry(ctx context.Context, tableName, code string, entry *CodeTable) (*CodeTable, error) {
	var result CodeTable
	if err := c.put(ctx, "codeTable/"+tableName+"/"+code, entry, &result); err != nil {
		return nil, fmt.Errorf("updating code table %s entry %s: %w", tableName, code, err)
	}
	return &result, nil
}

// DeleteCodeTableEntry deletes a code table entry by table name and code.
// Uses DELETE /codeTable/{codeTable}/{code}.
func (c *Client) DeleteCodeTableEntry(ctx context.Context, tableName, code string) error {
	if err := c.delete(ctx, "codeTable/"+tableName+"/"+code); err != nil {
		return fmt.Errorf("deleting code table %s entry %s: %w", tableName, code, err)
	}
	return nil
}
