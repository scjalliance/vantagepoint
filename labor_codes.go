package vantagepoint

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// LaborCode represents a Vantagepoint labor code record.
type LaborCode struct {
	LaborCode   string `json:"LaborCode,omitempty"`
	Description string `json:"Description,omitempty"`
	Category    string `json:"Category,omitempty"`
	Status      string `json:"Status,omitempty"`
}

// ListLaborCodes retrieves a list of labor codes from the Vantagepoint API.
// Uses GET /accountConfiguration/laborCode.
// Pass a *Query to filter, paginate, or sort results; nil retrieves defaults.
func (c *Client) ListLaborCodes(ctx context.Context, q *Query) ([]LaborCode, error) {
	var codes []LaborCode
	if err := c.get(ctx, "accountConfiguration/laborCode", q, &codes); err != nil {
		return nil, fmt.Errorf("listing labor codes: %w", err)
	}
	return codes, nil
}

// GetLaborCode retrieves a single labor code by its code and level. The
// Vantagepoint API returns an array even for single-resource lookups, so the
// response is decoded as a slice and the first element is returned.
// Uses GET /accountConfiguration/laborCode/{code}?LCLevel={level}.
func (c *Client) GetLaborCode(ctx context.Context, code string, lcLevel int) (*LaborCode, error) {
	var codes []LaborCode
	q := make(url.Values)
	if lcLevel > 0 {
		q.Set("LCLevel", strconv.Itoa(lcLevel))
	}
	if err := c.do(ctx, "GET", "accountConfiguration/laborCode/"+code, q, nil, &codes); err != nil {
		return nil, fmt.Errorf("getting labor code %s: %w", code, err)
	}
	if len(codes) == 0 {
		return nil, fmt.Errorf("getting labor code %s: %w", code, ErrNotFound)
	}
	return &codes[0], nil
}

// CreateLaborCode creates a new labor code record in Vantagepoint.
// Uses POST /accountConfiguration/laborCode.
func (c *Client) CreateLaborCode(ctx context.Context, laborCode *LaborCode) (*LaborCode, error) {
	var result LaborCode
	if err := c.post(ctx, "accountConfiguration/laborCode", laborCode, &result); err != nil {
		return nil, fmt.Errorf("creating labor code: %w", err)
	}
	return &result, nil
}

// UpdateLaborCode updates an existing labor code record.
// Uses PUT /accountConfiguration/laborCode/{code}.
func (c *Client) UpdateLaborCode(ctx context.Context, code string, laborCode *LaborCode) (*LaborCode, error) {
	var result LaborCode
	if err := c.put(ctx, "accountConfiguration/laborCode/"+code, laborCode, &result); err != nil {
		return nil, fmt.Errorf("updating labor code %s: %w", code, err)
	}
	return &result, nil
}

// DeleteLaborCode deletes a labor code record by its code.
// Uses DELETE /accountConfiguration/laborCode/{code}.
func (c *Client) DeleteLaborCode(ctx context.Context, code string) error {
	if err := c.delete(ctx, "accountConfiguration/laborCode/"+code); err != nil {
		return fmt.Errorf("deleting labor code %s: %w", code, err)
	}
	return nil
}
