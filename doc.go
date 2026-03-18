// Package vantagepoint provides a Go client for the Deltek Vantagepoint REST API.
//
// Vantagepoint is a project-based ERP system used for professional services
// firms. This package wraps the REST API to provide typed access to resources
// such as projects, employees, timesheets, firms, contacts, and more.
//
// # Authentication
//
// The API uses OAuth 2.0 password grant authentication. Create a client and
// authenticate before making requests:
//
//	client := vantagepoint.NewClient(
//		"https://example.deltekfirst.com/Instance/api",
//		"DatabaseName",
//		"client-id",
//		"client-secret",
//		vantagepoint.WithAutoRefresh(true),
//	)
//	err := client.Authenticate(ctx, "username", "password")
//
// # Querying
//
// Use the Query builder to construct filtered, paginated, and sorted requests:
//
//	q := vantagepoint.NewQuery().
//		Limit(50).
//		Fields("WBS1", "Name", "Status").
//		Filter("Status", "eq", "Active").
//		OrderBy("Name")
//	projects, err := client.ListProjects(ctx, q)
//
// # Resources
//
// The client provides methods for the following Vantagepoint resources:
//   - Projects (WBS1/WBS2/WBS3 — phases and tasks are projects with WBSType filter)
//   - Project Sub-Resources (employees, team members, revenue, firm clients)
//   - Employees (with project assignment sub-resource)
//   - Timesheets (three-tier DataEntry: tsControl/tsMaster/tsDetail)
//   - Firms (clients/vendors, keyed by ClientID, with employee/project sub-resources)
//   - Contacts (with qualification workflow)
//   - Organizations (read-only)
//   - Activities (with sub-resources: contacts, employees, clients, links)
//   - Labor Codes (under accountConfiguration/laborCode)
//   - Code Tables (codeTable, camelCase path)
//   - Invoices (three-tier DataEntry: inControl/inMaster/inDetail)
//   - Expenses (three-tier DataEntry: exControl/exMaster/exDetail)
//   - Units (three-tier DataEntry: unControl/unMaster/unDetail)
//   - Units By Project (three-tier DataEntry: upControl/upMaster/upDetail)
//   - Interactive Billing (read-only InteractiveDetail endpoints)
//   - GL Summary (aggregated general ledger data)
//   - Labor Detail (per-transaction labor costs)
//   - PSA Ledger (general ledger entries by transaction type)
//   - AR Review (accounts receivable balances and cash receipts)
//   - All Landing (cross-transaction-type batch listing)
//   - Metadata (per-table, no generic list endpoint)
package vantagepoint
