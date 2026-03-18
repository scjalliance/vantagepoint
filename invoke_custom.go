package vantagepoint

import "context"

// InvokeCustom calls a custom stored procedure via the Utilities/InvokeCustom
// endpoint. The procName is the procedure name, params are passed as the POST
// body, and result receives the JSON-decoded response.
func (c *Client) InvokeCustom(ctx context.Context, procName string, params map[string]any, result any) error {
	return c.post(ctx, "Utilities/InvokeCustom/"+procName, params, result)
}
