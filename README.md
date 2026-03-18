# vantagepoint

Go client library for the [Deltek Vantagepoint](https://www.deltek.com/en/products/project-erp/vantagepoint) REST API.

Provides typed access to projects, employees, timesheets, firms, contacts, general ledger, labor detail, AR review, and more.

## Installation

```bash
go get github.com/scjalliance/vantagepoint
```

## Authentication

The Vantagepoint API uses OAuth 2.0 password grant authentication.

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/scjalliance/vantagepoint"
)

func main() {
	client := vantagepoint.NewClient(
		"https://example.deltekfirst.com/Instance/api",
		"DatabaseName",
		"client-id",
		"client-secret",
		vantagepoint.WithAutoRefresh(true),
	)

	ctx := context.Background()
	if err := client.Authenticate(ctx, "username", "password"); err != nil {
		log.Fatal(err)
	}

	// List active projects
	q := vantagepoint.NewQuery().
		Limit(10).
		Fields("WBS1", "Name", "Status").
		Filter("Status", "eq", "Active")
	projects, err := client.ListProjects(ctx, q)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range projects {
		fmt.Printf("%s  %s\n", p.WBS1, p.Name)
	}
}
```

## Query Builder

Use `NewQuery()` to construct filtered, paginated, and sorted requests:

```go
q := vantagepoint.NewQuery().
	Limit(50).
	Offset(100).
	Fields("WBS1", "Name", "Status", "ClientID").
	Filter("Status", "eq", "Active").
	OrderBy("Name")

projects, err := client.ListProjects(ctx, q)
```

Supported filter operators: `eq`, `ne`, `gt`, `lt`, `ge`, `le`, `like`, `in`.

## Resources

The client covers the following Vantagepoint API resources:

- **Projects** — WBS1/WBS2/WBS3 (phases and tasks)
- **Project Sub-Resources** — employees, team members, revenue, firms, awards, contracts, proposals, milestones, links, files
- **Employees** — with project assignment sub-resource
- **Timesheets** — three-tier DataEntry (tsControl/tsMaster/tsDetail)
- **Firms** — clients/vendors with addresses, aliases, links
- **Contacts** — with links, campaigns, categories
- **Organizations**
- **Activities** — with contacts, employees, clients, links
- **Labor Codes**
- **Code Tables**
- **Invoices** — three-tier DataEntry (inControl/inMaster/inDetail)
- **Expenses** — three-tier DataEntry (exControl/exMaster/exDetail)
- **Units** — three-tier DataEntry (unControl/unMaster/unDetail)
- **Interactive Billing** — read-only billing detail endpoints
- **GL Summary** — aggregated general ledger data
- **Labor Detail** — per-transaction labor costs
- **PSA Ledger** — general ledger entries by transaction type
- **AR Review** — accounts receivable balances and cash receipts
- **Metadata** — field definitions per table
- **Settings** — firm-wide configuration, companies, periods, key format
- **Phase Budget Report** — computed budget/spent/remaining metrics per WBS2

## CLI Tool

The `cmd/vptest` directory contains a read-only CLI for exercising API calls against a Vantagepoint instance. See [cmd/vptest/VPTEST.md](cmd/vptest/VPTEST.md) for usage.

## API Version Compatibility

This client is developed against the Deltek Vantagepoint REST API. It has been tested with recent Vantagepoint releases. Some endpoints or fields may vary between versions.

## License

MIT — see [LICENSE](LICENSE).
