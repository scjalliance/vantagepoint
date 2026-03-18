package vantagepoint

import (
	"context"
	"fmt"
)

// Boilerplate represents a boilerplate (text library) record used for
// CRM proposal and document templates.
type Boilerplate struct {
	Name               string `json:"Name,omitempty"`
	DisplayName        string `json:"DisplayName,omitempty"`
	Description        string `json:"Description,omitempty"`
	Document           string `json:"Document,omitempty"`
	CustomCurrencyCode string `json:"CustomCurrencyCode,omitempty"`
	CreateUser         string `json:"CreateUser,omitempty"`
	CreateDate         string `json:"CreateDate,omitempty"`
	ModUser            string `json:"ModUser,omitempty"`
	ModDate            string `json:"ModDate,omitempty"`
}

// ListBoilerplates retrieves a list of boilerplate records.
// Uses GET /boilerplate.
func (c *Client) ListBoilerplates(ctx context.Context, q *Query) ([]Boilerplate, error) {
	var results []Boilerplate
	if err := c.get(ctx, "boilerplate", q, &results); err != nil {
		return nil, fmt.Errorf("listing boilerplates: %w", err)
	}
	return results, nil
}

// GetBoilerplate retrieves a single boilerplate by name.
// Uses GET /boilerplate/{name}.
func (c *Client) GetBoilerplate(ctx context.Context, name string) (*Boilerplate, error) {
	var results []Boilerplate
	if err := c.get(ctx, "boilerplate/"+name, nil, &results); err != nil {
		return nil, fmt.Errorf("getting boilerplate %s: %w", name, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("getting boilerplate %s: %w", name, ErrNotFound)
	}
	return &results[0], nil
}

// CreateBoilerplate creates a new boilerplate record.
// Uses POST /boilerplate.
func (c *Client) CreateBoilerplate(ctx context.Context, bp *Boilerplate) (*Boilerplate, error) {
	var results []Boilerplate
	if err := c.post(ctx, "boilerplate", bp, &results); err != nil {
		return nil, fmt.Errorf("creating boilerplate: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("creating boilerplate: %w", ErrNotFound)
	}
	return &results[0], nil
}

// UpdateBoilerplate updates an existing boilerplate record.
// Uses PUT /boilerplate/{name}.
func (c *Client) UpdateBoilerplate(ctx context.Context, name string, bp *Boilerplate) (*Boilerplate, error) {
	var results []Boilerplate
	if err := c.put(ctx, "boilerplate/"+name, bp, &results); err != nil {
		return nil, fmt.Errorf("updating boilerplate %s: %w", name, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("updating boilerplate %s: %w", name, ErrNotFound)
	}
	return &results[0], nil
}

// DeleteBoilerplate deletes a boilerplate record by name.
// Uses DELETE /boilerplate/{name}.
func (c *Client) DeleteBoilerplate(ctx context.Context, name string) error {
	if err := c.delete(ctx, "boilerplate/"+name); err != nil {
		return fmt.Errorf("deleting boilerplate %s: %w", name, err)
	}
	return nil
}

// BoilerplateLink represents a link associated with a boilerplate record.
type BoilerplateLink struct {
	Description string `json:"Description,omitempty"`
	FilePath    string `json:"FilePath,omitempty"`
	Graphic     string `json:"Graphic,omitempty"`
}

// ListBoilerplateLinks retrieves links associated with a boilerplate.
// Uses GET /boilerplate/{name}/links.
func (c *Client) ListBoilerplateLinks(ctx context.Context, name string, q *Query) ([]BoilerplateLink, error) {
	var results []BoilerplateLink
	if err := c.get(ctx, "boilerplate/"+name+"/links", q, &results); err != nil {
		return nil, fmt.Errorf("listing boilerplate links for %s: %w", name, err)
	}
	return results, nil
}

// ListBoilerplateFiles retrieves supporting documents for a boilerplate.
// Uses GET /boilerplate/{name}/files.
func (c *Client) ListBoilerplateFiles(ctx context.Context, name string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "boilerplate/"+name+"/files", q, &results); err != nil {
		return nil, fmt.Errorf("listing boilerplate files for %s: %w", name, err)
	}
	return results, nil
}

// ListBoilerplateCustomTable retrieves custom table data for a boilerplate.
// Uses GET /boilerplate/{name}/customTable/{customTable}.
func (c *Client) ListBoilerplateCustomTable(ctx context.Context, name, customTable string, q *Query) ([]map[string]any, error) {
	var results []map[string]any
	if err := c.get(ctx, "boilerplate/"+name+"/customTable/"+customTable, q, &results); err != nil {
		return nil, fmt.Errorf("listing boilerplate custom table %s for %s: %w", customTable, name, err)
	}
	return results, nil
}
