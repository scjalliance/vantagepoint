# vptest CLI

A test CLI for exercising read-only Deltek Vantagepoint API calls against a production instance. It uses the `github.com/scjalliance/vantagepoint` Go client library.

All commands are **read-only**. Nothing is created, updated, or deleted.

## Running

```bash
go run ./cmd/vptest/ [-env <file>] <command> [args...]
```

The `-env` flag is optional and must appear **before** the command. It loads a `.env` file into the process environment. Real environment variables take precedence over file values.

```bash
go run ./cmd/vptest/ -env .secrets/.env project 20-000001
```

## Configuration

Credentials are read from environment variables. Set them directly or load them from a `.env` file with `-env`.

| Variable | Description |
|---|---|
| `VP_BASE_URL` | API root URL (e.g., `https://host/Instance/api`) |
| `VP_DATABASE` | Database name |
| `VP_CLIENT_ID` | OAuth 2.0 client ID |
| `VP_CLIENT_SECRET` | OAuth 2.0 client secret |
| `VP_USERNAME` | Login username |
| `VP_PASSWORD` | Login password |

All six are required. The `.env` file format supports `KEY=VALUE` lines, `#` comments, and optional single/double quoting of values.

## Output

- **JSON data** is written to **stdout** (pipe to `jq` for filtering).
- **Status/log messages** are written to **stderr** (structured slog format).

This means you can do:

```bash
go run ./cmd/vptest/ -env .secrets/.env projects 10 2>/dev/null | jq '.[].Name'
```

## Filtering & Search

List commands support filtering via `-s` (search/like) and `-f` (explicit filter) flags. These can be mixed with positional arguments like `[limit]` in any order.

### `-s` — Search (like)

```bash
# Search firms whose Name contains "Alliance"
vptest firms 50 -s Name:Alliance

# If no field is specified, the command's default field is used
vptest firms 50 -s Alliance
```

### `-f` — Filter with operator

```bash
# Exact match
vptest firms 50 -f Status:eq:A

# Combine search with filter
vptest projects 20 -s Name:Sewer -f Status:eq:A
```

Supported operators: `eq`, `ne`, `gt`, `lt`, `ge`, `le`, `like`, `in`.

### Default search fields by command

| Command | Default `-s` field |
|---|---|
| `projects` | `Name` |
| `employees` | `LastName` |
| `firms` | `Name` |
| `labor-codes` | `Description` |
| `system-labels` | `LabelName` |
| `boilerplates` | `Name` |
| `campaigns` | `Name` |
| `ts-batches` | `Employee` |

## Commands

### General

| Command | Args | Description |
|---|---|---|
| `orgs` | | List all organizations. |
| `projects` | `[limit]` | List projects. Default limit: 5. Returns WBS1, Name, Status, ClientID, ProjMgr fields. |
| `project` | `<wbs1>` | Get a single project by its WBS1 code (e.g., `20-000001`). Returns all fields. |
| `employee` | `<code>` | Get a single employee by their employee code (e.g., `JDOE`). Returns all fields. |
| `employees` | `[limit]` | List employees. Default limit: 5. Returns Employee, FirstName, LastName, Title, EMail, Status fields. |
| `firms` | `[limit]` | List firms (clients/vendors). Default limit: 5. Returns ClientID, Name, Status fields. |
| `labor-codes` | `[limit]` | List labor codes. Default limit: 10. |
| `code-table` | `<table>` | List all entries in a named code table. |
| `metadata` | `<table>` | Get field metadata (names, types, constraints) for a table. Known tables: `activity`, `contact`, `employee`, `firm`, `organization`, `project`, `APEntry`, `TSEntry`, `GLSummary`, etc. |
| `code-table-metadata` | `<table>` | Get field metadata for a code table. |

### Project Structure

All of these take a WBS1 project code (e.g., `20-000001`) and return data scoped to that project.

| Command | Args | Description |
|---|---|---|
| `phases` | `<wbs1>` | List WBS2 phases under the project. |
| `project-employees` | `<wbs1>` | List employee assignments (role, dates, hours). |
| `project-team` | `<wbs1>` | List team members including both employees and contacts. |
| `project-revenue` | `<wbs1>` | List revenue allocation records. |
| `project-firms` | `<wbs1>` | List firm/client associations (subconsultants, clients). |

### Project Financials

All of these take a WBS1 project code and return financial data filtered to that project.

| Command | Args | Description |
|---|---|---|
| `labor-detail` | `<wbs1> [limit]` | Posted labor transactions (hours, amounts, rates, billing status). Default limit: 25. |
| `psa-ledger` | `<type> <wbs1> [limit]` | Posted PSA ledger entries filtered by transaction type and project. Default limit: 25. Transaction types: `L` (labor), `E` (expense), `AP` (accounts payable), `M` (miscellaneous). |
| `ar-balances` | `<wbs1>` | Accounts receivable balances (invoice-level with amounts and billing client). |
| `billing-expenses` | `<wbs1>` | Interactive billing expense records. |
| `billing-units` | `<wbs1>` | Interactive billing unit records. |
| `billing-limits` | `<wbs1>` | Billing limits/caps for the project. |
| `invoice-headers` | `<wbs1>` | Invoice master/header records for the project. |

### Project Sub-Resources

All of these take a WBS1 project code (e.g., `20-000001`).

| Command | Args | Description |
|---|---|---|
| `project-awards` | `<wbs1>` | List award records for a project. |
| `project-contracts` | `<wbs1>` | List contract records for a project (fee breakdowns, dates, status). |
| `project-proposals` | `<wbs1>` | List proposal records for a project. |
| `project-milestones` | `<wbs1>` | List milestone records for a project. |
| `project-competition` | `<wbs1>` | List competition records for a project. |
| `project-links` | `<wbs1>` | List links (file paths, URLs) for a project. |
| `project-files` | `<wbs1>` | List supporting documents attached to a project. |

### Timesheets

| Command | Args | Description |
|---|---|---|
| `ts-batches` | `[limit]` | List timesheet batch summaries from the landing endpoint. Default limit: 10. |
| `ts-detail` | `<batch> <wbs1>` | List individual timesheet line items for a specific batch, filtered to entries charged to the given project. |

### Settings & Configuration

| Command | Args | Description |
|---|---|---|
| `settings` | | General firm-wide settings (name, address, phone). |
| `companies` | | List all company records. |
| `active-company` | | Get the currently active company. |
| `periods` | | List all accounting periods. |
| `active-period` | | Get the currently active accounting period. |
| `system-labels` | `[limit]` | List system UI labels. Default limit: 20. |
| `key-format` | | Key conversion format config (WBS/Employee/Account key structure). |
| `org-levels` | | Organization level definitions. |
| `login-config` | `[database]` | Login configuration. Optionally pass a database name for database-specific config. |

### CRM

| Command | Args | Description |
|---|---|---|
| `boilerplates` | `[limit]` | List boilerplate (text library) records. Default limit: 10. |
| `campaigns` | `[limit]` | List marketing campaigns. Default limit: 10. |
| `campaign` | `<id>` | Get a single campaign by its CampaignID. |
| `udic` | `<area> [limit]` | List UDIC (user-defined info center) records for an infocenter area. Default limit: 10. |

### Firm Sub-Resources

| Command | Args | Description |
|---|---|---|
| `firm-addresses` | `<clientID>` | List addresses for a firm. |
| `firm-aliases` | `<clientID>` | List aliases for a firm. |
| `firm-links` | `<clientID>` | List links for a firm. |

### Contact Sub-Resources

| Command | Args | Description |
|---|---|---|
| `contact-links` | `<contactID>` | List links for a contact. |
| `contact-campaigns` | `<contactID>` | List campaign associations for a contact. |
| `contact-categories` | `<contactID>` | List category assignments for a contact. |

### Debugging

| Command | Args | Description |
|---|---|---|
| `raw` | `<path>` | Perform a raw GET against any API path and dump the JSON response. The path is relative to the API root (e.g., `DataEntry/tsLanding`, `project`, `employee/JDOE`). Useful for exploring endpoints or debugging unexpected results. |

## Examples

```bash
# Get a project and pipe through jq
go run ./cmd/vptest/ -env .secrets/.env project 20-000001 2>/dev/null | jq .

# List 20 projects
go run ./cmd/vptest/ -env .secrets/.env projects 20

# See all phases under a project
go run ./cmd/vptest/ -env .secrets/.env phases 20-000001

# Who is assigned to a project?
go run ./cmd/vptest/ -env .secrets/.env project-employees 20-000001

# Get posted labor for a project (last 50 entries)
go run ./cmd/vptest/ -env .secrets/.env labor-detail 20-000001 50

# Get all expense ledger entries for a project
go run ./cmd/vptest/ -env .secrets/.env psa-ledger E 20-000001 100

# Check AR balances
go run ./cmd/vptest/ -env .secrets/.env ar-balances 20-000001

# View invoices for a project
go run ./cmd/vptest/ -env .secrets/.env invoice-headers 20-000001

# Explore a raw API endpoint
go run ./cmd/vptest/ -env .secrets/.env raw DataEntry/tsLanding

# Get field schema for the project table
go run ./cmd/vptest/ -env .secrets/.env metadata project

# View contracts on a project
go run ./cmd/vptest/ -env .secrets/.env project-contracts 20-000001

# View supporting documents attached to a project
go run ./cmd/vptest/ -env .secrets/.env project-files 20-000001

# Get firm-wide settings
go run ./cmd/vptest/ -env .secrets/.env settings

# List accounting periods
go run ./cmd/vptest/ -env .secrets/.env periods

# View addresses for a firm
go run ./cmd/vptest/ -env .secrets/.env firm-addresses 1XX-ALL

# Search firms by name (substring match)
go run ./cmd/vptest/ -env .secrets/.env firms 50 -s Name:Alliance

# Search active firms only
go run ./cmd/vptest/ -env .secrets/.env firms 50 -s Name:Alliance -f Status:eq:A

# Find employees by last name
go run ./cmd/vptest/ -env .secrets/.env employees 20 -s LastName:Smith

# Search projects by name and filter to active
go run ./cmd/vptest/ -env .secrets/.env projects 20 -s Name:Sewer -f Status:eq:A
```

## Key Concepts

- **WBS1** is the project number (e.g., `20-000001`). Most project-centric commands take this as their primary argument.
- **WBS2** is a phase within a project. **WBS3** is a task within a phase.
- **Limit** arguments are optional integers that cap the number of results returned. If omitted, each command uses its own sensible default.
- **PSA Ledger transaction types**: `L` = labor, `E` = expense, `AP` = accounts payable, `M` = miscellaneous.
- All API responses are JSON arrays, even for single-resource lookups. The CLI handles this transparently — single-resource commands (`project`, etc.) return the unwrapped object.
