// Command vptest is a test CLI that exercises read-only Vantagepoint API calls
// to verify the client library works against a production instance.
//
// Configuration is via environment variables:
//
//	VP_BASE_URL      - API root URL (e.g., "https://host/Instance/api")
//	VP_DATABASE      - Database name
//	VP_CLIENT_ID     - OAuth 2.0 client ID
//	VP_CLIENT_SECRET - OAuth 2.0 client secret
//	VP_USERNAME      - Login username
//	VP_PASSWORD      - Login password
//
// Use -env to load variables from a .env file (parsed before commands run):
//
//	vptest -env .secrets/.env <command> [args...]
//
// Usage:
//
//	vptest [-env <file>] <command> [args...] [-s field:value] [-f field:op:value]
//
// List commands support filtering with -s (like/search) and -f (explicit operator).
// Each command has a default search field (e.g., Name for firms, LastName for employees).
// Operators: eq, ne, gt, lt, ge, le, like, in.
//
// Commands:
//
//	orgs                          - List organizations
//	projects [limit]              - List projects (default limit: 5)
//	project <wbs1>                - Get a single project by WBS1
//	employee <code>               - Get a single employee by code
//	phases <wbs1>                 - List phases for a project
//	project-employees <wbs1>      - List employee assignments on a project
//	project-team <wbs1>           - List team members on a project
//	project-revenue <wbs1>        - List revenue allocations on a project
//	project-firms <wbs1>          - List firm/client associations on a project
//	employees [limit]             - List employees (default limit: 5)
//	firms [limit]                 - List firms (default limit: 5)
//	labor-codes [limit]           - List labor codes (default limit: 10)
//	code-table <table>            - List entries in a code table
//	metadata <table>              - Get table metadata
//	code-table-metadata <t>       - Get code table metadata
//	labor-detail <wbs1> [limit]   - List posted labor detail for a project
//	psa-ledger <type> <wbs1> [n]  - List PSA ledger entries (type: L/E/AP/M)
//	ar-balances <wbs1>            - List AR balances for a project
//	billing-expenses <wbs1>       - List billing expenses for a project
//	billing-units <wbs1>          - List billing units for a project
//	billing-limits <wbs1>         - List billing limits for a project
//	invoice-headers <wbs1>        - List invoice headers for a project
//	project-awards <wbs1>         - List awards for a project
//	project-contracts <wbs1>      - List contracts for a project
//	project-proposals <wbs1>      - List proposals for a project
//	project-milestones <wbs1>     - List milestones for a project
//	project-competition <wbs1>    - List competition for a project
//	project-links <wbs1>          - List links for a project
//	project-files <wbs1>          - List supporting documents for a project
//	ts-batches [limit]            - List timesheet batches (default limit: 10)
//	ts-detail <batch> <wbs1>      - List timesheet detail for a batch/project
//	settings                      - General firm-wide settings
//	companies                     - List companies
//	active-company                - Active company
//	periods                       - List accounting periods
//	active-period                 - Active accounting period
//	system-labels [limit]         - List system labels (default limit: 20)
//	key-format                    - Key conversion format config
//	org-levels                    - Organization level definitions
//	login-config [database]       - Login configuration
//	boilerplates [limit]          - List boilerplates (default limit: 10)
//	campaigns [limit]             - List marketing campaigns (default limit: 10)
//	campaign <id>                 - Get a single campaign
//	udic <area> [limit]           - List UDIC records (default limit: 10)
//	firm-addresses <clientID>     - List addresses for a firm
//	firm-aliases <clientID>       - List aliases for a firm
//	firm-links <clientID>         - List links for a firm
//	contact-links <contactID>     - List links for a contact
//	contact-campaigns <contactID> - List campaigns for a contact
//	contact-categories <contactID> - List categories for a contact
//	raw <path>                    - Raw GET against any API path
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/scjalliance/vantagepoint"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))

	// Parse -env flag before subcommand.
	args := os.Args[1:]
	if len(args) >= 2 && args[0] == "-env" {
		if err := loadEnvFile(args[1]); err != nil {
			logger.Error("failed to load env file", "path", args[1], "err", err)
			os.Exit(1)
		}
		logger.Info("loaded env file", "path", args[1])
		args = args[2:]
	}

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: vptest [-env <file>] <command> [args...] [-s field:value] [-f field:op:value]\n")
		fmt.Fprintf(os.Stderr, "\nGeneral:\n")
		fmt.Fprintf(os.Stderr, "  orgs                          List organizations\n")
		fmt.Fprintf(os.Stderr, "  projects [limit]              List projects (default: 5)\n")
		fmt.Fprintf(os.Stderr, "  project <wbs1>                Get a single project\n")
		fmt.Fprintf(os.Stderr, "  employee <code>               Get a single employee\n")
		fmt.Fprintf(os.Stderr, "  employees [limit]             List employees (default: 5)\n")
		fmt.Fprintf(os.Stderr, "  firms [limit]                 List firms (default: 5)\n")
		fmt.Fprintf(os.Stderr, "  labor-codes [limit]           List labor codes (default: 10)\n")
		fmt.Fprintf(os.Stderr, "  code-table <table>            List code table entries\n")
		fmt.Fprintf(os.Stderr, "  metadata <table>              Get table metadata\n")
		fmt.Fprintf(os.Stderr, "  code-table-metadata <t>       Get code table metadata\n")
		fmt.Fprintf(os.Stderr, "\nProject structure:\n")
		fmt.Fprintf(os.Stderr, "  phases <wbs1>                 List phases for a project\n")
		fmt.Fprintf(os.Stderr, "  project-employees <wbs1>      List employee assignments\n")
		fmt.Fprintf(os.Stderr, "  project-team <wbs1>           List team members\n")
		fmt.Fprintf(os.Stderr, "  project-revenue <wbs1>        List revenue allocations\n")
		fmt.Fprintf(os.Stderr, "  project-firms <wbs1>          List firm/client associations\n")
		fmt.Fprintf(os.Stderr, "\nProject financials:\n")
		fmt.Fprintf(os.Stderr, "  labor-detail <wbs1> [limit]   Posted labor detail (default: 25)\n")
		fmt.Fprintf(os.Stderr, "  psa-ledger <type> <wbs1> [n]  PSA ledger (type: L/E/AP/M, default: 25)\n")
		fmt.Fprintf(os.Stderr, "  ar-balances <wbs1>            AR balances for a project\n")
		fmt.Fprintf(os.Stderr, "  billing-expenses <wbs1>       Billing expenses\n")
		fmt.Fprintf(os.Stderr, "  billing-units <wbs1>          Billing units\n")
		fmt.Fprintf(os.Stderr, "  billing-limits <wbs1>         Billing limits\n")
		fmt.Fprintf(os.Stderr, "  invoice-headers <wbs1>        Invoice headers\n")
		fmt.Fprintf(os.Stderr, "\nProject sub-resources:\n")
		fmt.Fprintf(os.Stderr, "  project-awards <wbs1>         Awards for a project\n")
		fmt.Fprintf(os.Stderr, "  project-contracts <wbs1>      Contracts for a project\n")
		fmt.Fprintf(os.Stderr, "  project-proposals <wbs1>      Proposals for a project\n")
		fmt.Fprintf(os.Stderr, "  project-milestones <wbs1>     Milestones for a project\n")
		fmt.Fprintf(os.Stderr, "  project-competition <wbs1>    Competition for a project\n")
		fmt.Fprintf(os.Stderr, "  project-links <wbs1>          Links for a project\n")
		fmt.Fprintf(os.Stderr, "  project-files <wbs1>          Supporting documents for a project\n")
		fmt.Fprintf(os.Stderr, "\nTimesheets:\n")
		fmt.Fprintf(os.Stderr, "  ts-batches [limit]            List timesheet batches (default: 10)\n")
		fmt.Fprintf(os.Stderr, "  ts-detail <batch> <wbs1>      Timesheet detail filtered by project\n")
		fmt.Fprintf(os.Stderr, "\nSettings & config:\n")
		fmt.Fprintf(os.Stderr, "  settings                      General settings\n")
		fmt.Fprintf(os.Stderr, "  companies                     List companies\n")
		fmt.Fprintf(os.Stderr, "  active-company                Active company\n")
		fmt.Fprintf(os.Stderr, "  periods                       List accounting periods\n")
		fmt.Fprintf(os.Stderr, "  active-period                 Active accounting period\n")
		fmt.Fprintf(os.Stderr, "  system-labels [limit]         System labels (default: 20)\n")
		fmt.Fprintf(os.Stderr, "  key-format                    Key conversion format config\n")
		fmt.Fprintf(os.Stderr, "  org-levels                    Organization level definitions\n")
		fmt.Fprintf(os.Stderr, "  login-config [database]       Login configuration\n")
		fmt.Fprintf(os.Stderr, "\nCRM:\n")
		fmt.Fprintf(os.Stderr, "  boilerplates [limit]          List boilerplates (default: 10)\n")
		fmt.Fprintf(os.Stderr, "  campaigns [limit]             List marketing campaigns (default: 10)\n")
		fmt.Fprintf(os.Stderr, "  campaign <id>                 Get a single campaign\n")
		fmt.Fprintf(os.Stderr, "  udic <area> [limit]           List UDIC records (default: 10)\n")
		fmt.Fprintf(os.Stderr, "\nFirm sub-resources:\n")
		fmt.Fprintf(os.Stderr, "  firm-addresses <clientID>     Addresses for a firm\n")
		fmt.Fprintf(os.Stderr, "  firm-aliases <clientID>       Aliases for a firm\n")
		fmt.Fprintf(os.Stderr, "  firm-links <clientID>         Links for a firm\n")
		fmt.Fprintf(os.Stderr, "\nContact sub-resources:\n")
		fmt.Fprintf(os.Stderr, "  contact-links <contactID>     Links for a contact\n")
		fmt.Fprintf(os.Stderr, "  contact-campaigns <contactID> Campaigns for a contact\n")
		fmt.Fprintf(os.Stderr, "  contact-categories <contactID> Categories for a contact\n")
		fmt.Fprintf(os.Stderr, "\nDebugging:\n")
		fmt.Fprintf(os.Stderr, "  raw <path>                    Raw GET against any API path\n")
		fmt.Fprintf(os.Stderr, "\nFiltering (works with list commands):\n")
		fmt.Fprintf(os.Stderr, "  -s field:value                Search (like) on a field\n")
		fmt.Fprintf(os.Stderr, "  -s value                      Search on command's default field\n")
		fmt.Fprintf(os.Stderr, "  -f field:op:value             Filter with operator (eq,ne,gt,lt,ge,le,like,in)\n")
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  firms 50 -s Name:Alliance     Firms with 'Alliance' in name\n")
		fmt.Fprintf(os.Stderr, "  employees 20 -s LastName:Smith Employees named Smith\n")
		fmt.Fprintf(os.Stderr, "  projects 10 -s Name:Sewer -f Status:eq:A  Active projects with 'Sewer'\n")
		os.Exit(1)
	}

	client, err := clientFromEnv()
	if err != nil {
		logger.Error("configuration error", "err", err)
		os.Exit(1)
	}

	ctx := context.Background()
	logger.Info("authenticating")
	if err := client.Authenticate(ctx, mustEnv("VP_USERNAME"), mustEnv("VP_PASSWORD")); err != nil {
		logger.Error("authentication failed", "err", err)
		os.Exit(1)
	}
	logger.Info("authenticated successfully")

	cmd := args[0]
	cmdArgs := args[1:]

	if err := run(ctx, logger, client, cmd, cmdArgs); err != nil {
		logger.Error("command failed", "command", cmd, "err", err)
		os.Exit(1)
	}
}

// run dispatches the given command to the appropriate handler.
func run(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, cmd string, args []string) error {
	switch cmd {
	// General
	case "orgs":
		return cmdOrgs(ctx, logger, c)
	case "projects":
		return cmdProjects(ctx, logger, c, args)
	case "project":
		return cmdProject(ctx, logger, c, args)
	case "employee":
		return cmdEmployee(ctx, logger, c, args)
	case "employees":
		return cmdEmployees(ctx, logger, c, args)
	case "firms":
		return cmdFirms(ctx, logger, c, args)
	case "labor-codes":
		return cmdLaborCodes(ctx, logger, c, args)
	case "code-table":
		return cmdCodeTable(ctx, logger, c, args)
	case "metadata":
		return cmdMetadata(ctx, logger, c, args)
	case "code-table-metadata":
		return cmdCodeTableMetadata(ctx, logger, c, args)
	// Project structure
	case "phases":
		return cmdPhases(ctx, logger, c, args)
	case "project-employees":
		return cmdProjectEmployees(ctx, logger, c, args)
	case "project-team":
		return cmdProjectTeam(ctx, logger, c, args)
	case "project-revenue":
		return cmdProjectRevenue(ctx, logger, c, args)
	case "project-firms":
		return cmdProjectFirms(ctx, logger, c, args)
	// Project financials
	case "labor-detail":
		return cmdLaborDetail(ctx, logger, c, args)
	case "psa-ledger":
		return cmdPSALedger(ctx, logger, c, args)
	case "ar-balances":
		return cmdARBalances(ctx, logger, c, args)
	case "billing-expenses":
		return cmdBillingExpenses(ctx, logger, c, args)
	case "billing-units":
		return cmdBillingUnits(ctx, logger, c, args)
	case "billing-limits":
		return cmdBillingLimits(ctx, logger, c, args)
	case "invoice-headers":
		return cmdInvoiceHeaders(ctx, logger, c, args)
	// Project sub-resources
	case "project-awards":
		return cmdProjectAwards(ctx, logger, c, args)
	case "project-contracts":
		return cmdProjectContracts(ctx, logger, c, args)
	case "project-proposals":
		return cmdProjectProposals(ctx, logger, c, args)
	case "project-milestones":
		return cmdProjectMilestones(ctx, logger, c, args)
	case "project-competition":
		return cmdProjectCompetition(ctx, logger, c, args)
	case "project-links":
		return cmdProjectLinks(ctx, logger, c, args)
	case "project-files":
		return cmdProjectFiles(ctx, logger, c, args)
	// Timesheets
	case "ts-batches":
		return cmdTSBatches(ctx, logger, c, args)
	case "ts-detail":
		return cmdTSDetail(ctx, logger, c, args)
	// Settings & config
	case "settings":
		return cmdSettings(ctx, logger, c)
	case "companies":
		return cmdCompanies(ctx, logger, c)
	case "active-company":
		return cmdActiveCompany(ctx, logger, c)
	case "periods":
		return cmdPeriods(ctx, logger, c)
	case "active-period":
		return cmdActivePeriod(ctx, logger, c)
	case "system-labels":
		return cmdSystemLabels(ctx, logger, c, args)
	case "key-format":
		return cmdKeyFormat(ctx, logger, c)
	case "org-levels":
		return cmdOrgLevels(ctx, logger, c)
	case "login-config":
		return cmdLoginConfig(ctx, logger, c, args)
	// CRM
	case "boilerplates":
		return cmdBoilerplates(ctx, logger, c, args)
	case "campaigns":
		return cmdCampaigns(ctx, logger, c, args)
	case "campaign":
		return cmdCampaign(ctx, logger, c, args)
	case "udic":
		return cmdUDIC(ctx, logger, c, args)
	// Firm sub-resources
	case "firm-addresses":
		return cmdFirmAddresses(ctx, logger, c, args)
	case "firm-aliases":
		return cmdFirmAliases(ctx, logger, c, args)
	case "firm-links":
		return cmdFirmLinks(ctx, logger, c, args)
	// Contact sub-resources
	case "contact-links":
		return cmdContactLinks(ctx, logger, c, args)
	case "contact-campaigns":
		return cmdContactCampaigns(ctx, logger, c, args)
	case "contact-categories":
		return cmdContactCategories(ctx, logger, c, args)
	// Debugging
	case "raw":
		return cmdRaw(ctx, logger, c, args)
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
}

func cmdOrgs(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("listing organizations")
	orgs, err := c.ListOrganizations(ctx, nil)
	if err != nil {
		return fmt.Errorf("listing organizations: %w", err)
	}
	logger.Info("fetched organizations", "count", len(orgs))
	return printJSON(orgs)
}

func cmdProjects(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 5)
	q := vantagepoint.NewQuery().Limit(limit).Fields("WBS1", "Name", "Status", "ClientID", "ProjMgr")
	applyFilters(q, p.filters, "Name")
	logger.Info("listing projects", "limit", limit)
	projects, err := c.ListProjects(ctx, q)
	if err != nil {
		return fmt.Errorf("listing projects: %w", err)
	}
	logger.Info("fetched projects", "count", len(projects))
	return printJSON(projects)
}

func cmdProject(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("getting project", "wbs1", wbs1)
	project, err := c.GetProject(ctx, wbs1)
	if err != nil {
		return fmt.Errorf("getting project %s: %w", wbs1, err)
	}
	return printJSON(project)
}

func cmdEmployee(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest employee <code>")
	}
	code := args[0]
	logger.Info("getting employee", "code", code)
	employee, err := c.GetEmployee(ctx, code)
	if err != nil {
		return fmt.Errorf("getting employee %s: %w", code, err)
	}
	return printJSON(employee)
}

func cmdEmployees(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 5)
	q := vantagepoint.NewQuery().Limit(limit).Fields("Employee", "FirstName", "LastName", "Title", "EMail", "Status")
	applyFilters(q, p.filters, "LastName")
	logger.Info("listing employees", "limit", limit)
	employees, err := c.ListEmployees(ctx, q)
	if err != nil {
		return fmt.Errorf("listing employees: %w", err)
	}
	logger.Info("fetched employees", "count", len(employees))
	return printJSON(employees)
}

func cmdFirms(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 5)
	q := vantagepoint.NewQuery().Limit(limit).Fields("ClientID", "Name", "Status")
	applyFilters(q, p.filters, "Name")
	logger.Info("listing firms", "limit", limit)
	firms, err := c.ListFirms(ctx, q)
	if err != nil {
		return fmt.Errorf("listing firms: %w", err)
	}
	logger.Info("fetched firms", "count", len(firms))
	return printJSON(firms)
}

func cmdLaborCodes(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 10)
	q := vantagepoint.NewQuery().Limit(limit)
	applyFilters(q, p.filters, "Description")
	logger.Info("listing labor codes", "limit", limit)
	codes, err := c.ListLaborCodes(ctx, q)
	if err != nil {
		return fmt.Errorf("listing labor codes: %w", err)
	}
	logger.Info("fetched labor codes", "count", len(codes))
	return printJSON(codes)
}

func cmdCodeTable(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest code-table <table-name>")
	}
	tableName := args[0]
	logger.Info("listing code table entries", "table", tableName)
	entries, err := c.ListCodeTables(ctx, tableName, nil)
	if err != nil {
		return fmt.Errorf("listing code table %s: %w", tableName, err)
	}
	logger.Info("fetched code table entries", "table", tableName, "count", len(entries))
	return printJSON(entries)
}

func cmdMetadata(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest metadata <table-name>")
	}
	tableName := args[0]
	logger.Info("getting metadata", "table", tableName)
	meta, err := c.GetMetadata(ctx, tableName)
	if err != nil {
		return fmt.Errorf("getting metadata for %s: %w", tableName, err)
	}
	return printJSON(meta)
}

func cmdCodeTableMetadata(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest code-table-metadata <code-table>")
	}
	tableName := args[0]
	logger.Info("getting code table metadata", "codeTable", tableName)
	meta, err := c.GetCodeTableMetadata(ctx, tableName)
	if err != nil {
		return fmt.Errorf("getting code table metadata for %s: %w", tableName, err)
	}
	return printJSON(meta)
}

// --- Project structure commands ---

func cmdPhases(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest phases <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing phases", "wbs1", wbs1)
	phases, err := c.ListPhases(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing phases for %s: %w", wbs1, err)
	}
	logger.Info("fetched phases", "count", len(phases))
	return printJSON(phases)
}

func cmdProjectEmployees(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-employees <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project employees", "wbs1", wbs1)
	employees, err := c.ListProjectEmployees(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project employees for %s: %w", wbs1, err)
	}
	logger.Info("fetched project employees", "count", len(employees))
	return printJSON(employees)
}

func cmdProjectTeam(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-team <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project team members", "wbs1", wbs1)
	members, err := c.ListProjectTeamMembers(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project team members for %s: %w", wbs1, err)
	}
	logger.Info("fetched project team members", "count", len(members))
	return printJSON(members)
}

func cmdProjectRevenue(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-revenue <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project revenue", "wbs1", wbs1)
	revenue, err := c.ListProjectRevenue(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project revenue for %s: %w", wbs1, err)
	}
	logger.Info("fetched project revenue", "count", len(revenue))
	return printJSON(revenue)
}

func cmdProjectFirms(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-firms <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project firm/clients", "wbs1", wbs1)
	clients, err := c.ListProjectFirmClients(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project firm/clients for %s: %w", wbs1, err)
	}
	logger.Info("fetched project firm/clients", "count", len(clients))
	return printJSON(clients)
}

// --- Project financial commands ---

func cmdLaborDetail(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest labor-detail <wbs1> [limit]")
	}
	wbs1 := args[0]
	limit := parseIntArg(args, 1, 25)
	q := vantagepoint.NewQuery().Limit(limit).Filter("WBS1", "eq", wbs1)
	logger.Info("listing labor detail", "wbs1", wbs1, "limit", limit)
	entries, err := c.ListLaborDetail(ctx, q)
	if err != nil {
		return fmt.Errorf("listing labor detail for %s: %w", wbs1, err)
	}
	logger.Info("fetched labor detail", "count", len(entries))
	return printJSON(entries)
}

func cmdPSALedger(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: vptest psa-ledger <type> <wbs1> [limit]\n  types: L (labor), E (expense), AP (accounts payable), M (miscellaneous)")
	}
	transType := args[0]
	wbs1 := args[1]
	limit := parseIntArg(args, 2, 25)
	q := vantagepoint.NewQuery().Limit(limit).MaxRowsOnly().Filter("WBS1", "eq", wbs1)
	logger.Info("listing PSA ledger", "type", transType, "wbs1", wbs1, "limit", limit)
	entries, err := c.ListPSALedger(ctx, transType, q)
	if err != nil {
		return fmt.Errorf("listing PSA ledger type %s for %s: %w", transType, wbs1, err)
	}
	logger.Info("fetched PSA ledger entries", "count", len(entries))
	return printJSON(entries)
}

func cmdARBalances(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest ar-balances <wbs1>")
	}
	wbs1 := args[0]
	q := vantagepoint.NewQuery().Filter("WBS1", "eq", wbs1)
	logger.Info("listing AR balances", "wbs1", wbs1)
	balances, err := c.ListARBalances(ctx, q)
	if err != nil {
		return fmt.Errorf("listing AR balances for %s: %w", wbs1, err)
	}
	logger.Info("fetched AR balances", "count", len(balances))
	return printJSON(balances)
}

func cmdBillingExpenses(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest billing-expenses <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing billing expenses", "wbs1", wbs1)
	expenses, err := c.GetBillingExpenses(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing billing expenses for %s: %w", wbs1, err)
	}
	logger.Info("fetched billing expenses", "count", len(expenses))
	return printJSON(expenses)
}

func cmdBillingUnits(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest billing-units <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing billing units", "wbs1", wbs1)
	units, err := c.GetBillingUnits(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing billing units for %s: %w", wbs1, err)
	}
	logger.Info("fetched billing units", "count", len(units))
	return printJSON(units)
}

func cmdBillingLimits(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest billing-limits <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing billing limits", "wbs1", wbs1)
	limits, err := c.GetBillingLimits(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing billing limits for %s: %w", wbs1, err)
	}
	logger.Info("fetched billing limits", "count", len(limits))
	return printJSON(limits)
}

func cmdInvoiceHeaders(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest invoice-headers <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing invoice headers", "wbs1", wbs1)
	headers, err := c.GetInvoiceHeaders(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing invoice headers for %s: %w", wbs1, err)
	}
	logger.Info("fetched invoice headers", "count", len(headers))
	return printJSON(headers)
}

// --- Timesheet commands ---

func cmdTSBatches(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 10)
	q := vantagepoint.NewQuery().Limit(limit)
	applyFilters(q, p.filters, "Employee")
	logger.Info("listing timesheet batches", "limit", limit)
	batches, err := c.ListTimesheetBatches(ctx, q)
	if err != nil {
		return fmt.Errorf("listing timesheet batches: %w", err)
	}
	logger.Info("fetched timesheet batches", "count", len(batches))
	return printJSON(batches)
}

func cmdTSDetail(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: vptest ts-detail <batch> <wbs1>")
	}
	batch := args[0]
	wbs1 := args[1]
	q := vantagepoint.NewQuery().Filter("WBS1", "eq", wbs1)
	logger.Info("listing timesheet detail", "batch", batch, "wbs1", wbs1)
	detail, err := c.GetTimesheetDetail(ctx, batch, q)
	if err != nil {
		return fmt.Errorf("listing timesheet detail for batch %s project %s: %w", batch, wbs1, err)
	}
	logger.Info("fetched timesheet detail", "count", len(detail))
	return printJSON(detail)
}

// --- Project sub-resource commands ---

func cmdProjectAwards(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-awards <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project awards", "wbs1", wbs1)
	results, err := c.ListProjectAwards(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project awards for %s: %w", wbs1, err)
	}
	logger.Info("fetched project awards", "count", len(results))
	return printJSON(results)
}

func cmdProjectContracts(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-contracts <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project contracts", "wbs1", wbs1)
	results, err := c.ListProjectContracts(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project contracts for %s: %w", wbs1, err)
	}
	logger.Info("fetched project contracts", "count", len(results))
	return printJSON(results)
}

func cmdProjectProposals(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-proposals <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project proposals", "wbs1", wbs1)
	results, err := c.ListProjectProposals(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project proposals for %s: %w", wbs1, err)
	}
	logger.Info("fetched project proposals", "count", len(results))
	return printJSON(results)
}

func cmdProjectMilestones(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-milestones <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project milestones", "wbs1", wbs1)
	results, err := c.ListProjectMilestones(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project milestones for %s: %w", wbs1, err)
	}
	logger.Info("fetched project milestones", "count", len(results))
	return printJSON(results)
}

func cmdProjectCompetition(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-competition <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project competition", "wbs1", wbs1)
	results, err := c.ListProjectCompetition(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project competition for %s: %w", wbs1, err)
	}
	logger.Info("fetched project competition", "count", len(results))
	return printJSON(results)
}

func cmdProjectLinks(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-links <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project links", "wbs1", wbs1)
	results, err := c.ListProjectLinks(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project links for %s: %w", wbs1, err)
	}
	logger.Info("fetched project links", "count", len(results))
	return printJSON(results)
}

func cmdProjectFiles(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest project-files <wbs1>")
	}
	wbs1 := args[0]
	logger.Info("listing project files", "wbs1", wbs1)
	results, err := c.ListProjectFiles(ctx, wbs1, nil)
	if err != nil {
		return fmt.Errorf("listing project files for %s: %w", wbs1, err)
	}
	logger.Info("fetched project files", "count", len(results))
	return printJSON(results)
}

// --- Settings & config commands ---

func cmdSettings(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("getting general settings")
	result, err := c.GetGeneralSettings(ctx)
	if err != nil {
		return fmt.Errorf("getting general settings: %w", err)
	}
	return printJSON(result)
}

func cmdCompanies(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("listing companies")
	results, err := c.ListCompanies(ctx, nil)
	if err != nil {
		return fmt.Errorf("listing companies: %w", err)
	}
	logger.Info("fetched companies", "count", len(results))
	return printJSON(results)
}

func cmdActiveCompany(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("getting active company")
	result, err := c.GetActiveCompany(ctx)
	if err != nil {
		return fmt.Errorf("getting active company: %w", err)
	}
	return printJSON(result)
}

func cmdPeriods(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("listing accounting periods")
	results, err := c.ListAccountingPeriods(ctx, nil)
	if err != nil {
		return fmt.Errorf("listing accounting periods: %w", err)
	}
	logger.Info("fetched accounting periods", "count", len(results))
	return printJSON(results)
}

func cmdActivePeriod(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("getting active period")
	result, err := c.GetActivePeriod(ctx)
	if err != nil {
		return fmt.Errorf("getting active period: %w", err)
	}
	return printJSON(result)
}

func cmdSystemLabels(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 20)
	q := vantagepoint.NewQuery().Limit(limit)
	applyFilters(q, p.filters, "LabelName")
	logger.Info("listing system labels", "limit", limit)
	results, err := c.ListSystemLabels(ctx, q)
	if err != nil {
		return fmt.Errorf("listing system labels: %w", err)
	}
	logger.Info("fetched system labels", "count", len(results))
	return printJSON(results)
}

func cmdKeyFormat(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("getting key format")
	results, err := c.GetKeyFormat(ctx)
	if err != nil {
		return fmt.Errorf("getting key format: %w", err)
	}
	logger.Info("fetched key format", "count", len(results))
	return printJSON(results)
}

func cmdOrgLevels(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client) error {
	logger.Info("getting org levels")
	results, err := c.GetOrgLevels(ctx)
	if err != nil {
		return fmt.Errorf("getting org levels: %w", err)
	}
	logger.Info("fetched org levels", "count", len(results))
	return printJSON(results)
}

func cmdLoginConfig(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) > 0 {
		database := args[0]
		logger.Info("getting login config for database", "database", database)
		result, err := c.GetLoginConfigForDatabase(ctx, database)
		if err != nil {
			return fmt.Errorf("getting login config for %s: %w", database, err)
		}
		return printJSON(result)
	}
	logger.Info("getting login config")
	result, err := c.GetLoginConfig(ctx)
	if err != nil {
		return fmt.Errorf("getting login config: %w", err)
	}
	return printJSON(result)
}

// --- CRM commands ---

func cmdBoilerplates(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 10)
	q := vantagepoint.NewQuery().Limit(limit)
	applyFilters(q, p.filters, "Name")
	logger.Info("listing boilerplates", "limit", limit)
	results, err := c.ListBoilerplates(ctx, q)
	if err != nil {
		return fmt.Errorf("listing boilerplates: %w", err)
	}
	logger.Info("fetched boilerplates", "count", len(results))
	return printJSON(results)
}

func cmdCampaigns(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	p := parseArgs(args)
	limit := parseIntArg(p.positional, 0, 10)
	q := vantagepoint.NewQuery().Limit(limit)
	applyFilters(q, p.filters, "Name")
	logger.Info("listing campaigns", "limit", limit)
	results, err := c.ListCampaigns(ctx, q)
	if err != nil {
		return fmt.Errorf("listing campaigns: %w", err)
	}
	logger.Info("fetched campaigns", "count", len(results))
	return printJSON(results)
}

func cmdCampaign(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest campaign <campaignID>")
	}
	campaignID := args[0]
	logger.Info("getting campaign", "campaignID", campaignID)
	result, err := c.GetCampaign(ctx, campaignID)
	if err != nil {
		return fmt.Errorf("getting campaign %s: %w", campaignID, err)
	}
	return printJSON(result)
}

func cmdUDIC(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest udic <infocenterArea> [limit]")
	}
	area := args[0]
	limit := parseIntArg(args, 1, 10)
	q := vantagepoint.NewQuery().Limit(limit)
	logger.Info("listing UDIC records", "area", area, "limit", limit)
	results, err := c.ListUDICRecords(ctx, area, q)
	if err != nil {
		return fmt.Errorf("listing UDIC records for %s: %w", area, err)
	}
	logger.Info("fetched UDIC records", "count", len(results))
	return printJSON(results)
}

// --- Firm sub-resource commands ---

func cmdFirmAddresses(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest firm-addresses <clientID>")
	}
	clientID := args[0]
	logger.Info("listing firm addresses", "clientID", clientID)
	results, err := c.ListFirmAddresses(ctx, clientID, nil)
	if err != nil {
		return fmt.Errorf("listing firm addresses for %s: %w", clientID, err)
	}
	logger.Info("fetched firm addresses", "count", len(results))
	return printJSON(results)
}

func cmdFirmAliases(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest firm-aliases <clientID>")
	}
	clientID := args[0]
	logger.Info("listing firm aliases", "clientID", clientID)
	results, err := c.ListFirmAliases(ctx, clientID, nil)
	if err != nil {
		return fmt.Errorf("listing firm aliases for %s: %w", clientID, err)
	}
	logger.Info("fetched firm aliases", "count", len(results))
	return printJSON(results)
}

func cmdFirmLinks(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest firm-links <clientID>")
	}
	clientID := args[0]
	logger.Info("listing firm links", "clientID", clientID)
	results, err := c.ListFirmLinks(ctx, clientID, nil)
	if err != nil {
		return fmt.Errorf("listing firm links for %s: %w", clientID, err)
	}
	logger.Info("fetched firm links", "count", len(results))
	return printJSON(results)
}

// --- Contact sub-resource commands ---

func cmdContactLinks(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest contact-links <contactID>")
	}
	contactID := args[0]
	logger.Info("listing contact links", "contactID", contactID)
	results, err := c.ListContactLinks(ctx, contactID, nil)
	if err != nil {
		return fmt.Errorf("listing contact links for %s: %w", contactID, err)
	}
	logger.Info("fetched contact links", "count", len(results))
	return printJSON(results)
}

func cmdContactCampaigns(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest contact-campaigns <contactID>")
	}
	contactID := args[0]
	logger.Info("listing contact campaigns", "contactID", contactID)
	results, err := c.ListContactCampaigns(ctx, contactID, nil)
	if err != nil {
		return fmt.Errorf("listing contact campaigns for %s: %w", contactID, err)
	}
	logger.Info("fetched contact campaigns", "count", len(results))
	return printJSON(results)
}

func cmdContactCategories(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest contact-categories <contactID>")
	}
	contactID := args[0]
	logger.Info("listing contact categories", "contactID", contactID)
	results, err := c.ListContactCategories(ctx, contactID, nil)
	if err != nil {
		return fmt.Errorf("listing contact categories for %s: %w", contactID, err)
	}
	logger.Info("fetched contact categories", "count", len(results))
	return printJSON(results)
}

// --- Debugging commands ---

func cmdRaw(ctx context.Context, logger *slog.Logger, c *vantagepoint.Client, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: vptest raw <api-path>")
	}
	path := args[0]
	logger.Info("raw GET", "path", path)
	var result any
	if err := c.RawGet(ctx, path, nil, &result); err != nil {
		return fmt.Errorf("raw GET %s: %w", path, err)
	}
	return printJSON(result)
}

// loadEnvFile reads a .env file and sets each KEY=VALUE pair as an environment
// variable. Lines that are empty, whitespace-only, or start with # are skipped.
// Values may optionally be wrapped in single or double quotes, which are stripped.
// Existing environment variables are NOT overwritten, so real env takes precedence.
func loadEnvFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("opening env file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		// Strip matching surrounding quotes.
		if len(val) >= 2 && ((val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'')) {
			val = val[1 : len(val)-1]
		}
		// Don't overwrite existing env vars — real environment takes precedence.
		if os.Getenv(key) == "" {
			os.Setenv(key, val)
		}
	}
	return scanner.Err()
}

// clientFromEnv creates a vantagepoint.Client from environment variables.
func clientFromEnv() (*vantagepoint.Client, error) {
	baseURL := os.Getenv("VP_BASE_URL")
	database := os.Getenv("VP_DATABASE")
	clientID := os.Getenv("VP_CLIENT_ID")
	clientSecret := os.Getenv("VP_CLIENT_SECRET")

	if baseURL == "" || database == "" || clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("required environment variables: VP_BASE_URL, VP_DATABASE, VP_CLIENT_ID, VP_CLIENT_SECRET, VP_USERNAME, VP_PASSWORD")
	}

	return vantagepoint.NewClient(baseURL, database, clientID, clientSecret), nil
}

// mustEnv returns the value of the environment variable or exits with an error.
func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		fmt.Fprintf(os.Stderr, "error: %s is required\n", key)
		os.Exit(1)
	}
	return v
}

// parseIntArg parses args[index] as an integer, returning defaultVal if missing or invalid.
func parseIntArg(args []string, index, defaultVal int) int {
	if index >= len(args) {
		return defaultVal
	}
	n, err := strconv.Atoi(args[index])
	if err != nil || n <= 0 {
		return defaultVal
	}
	return n
}

// parsedArgs holds positional arguments and any filter flags extracted from
// command arguments. Filters are specified as:
//
//	-f field:op:value   (e.g., -f Name:like:Alliance -f Status:eq:A)
//	-s field:value      (shorthand for -f field:like:value)
//
// Supported operators: eq, ne, gt, lt, ge, le, like, in.
type parsedArgs struct {
	positional []string
	filters    []filter
}

// parseArgs separates filter flags (-f, -s) from positional arguments.
func parseArgs(args []string) parsedArgs {
	var p parsedArgs
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-f":
			i++
			if i >= len(args) {
				break
			}
			parts := strings.SplitN(args[i], ":", 3)
			if len(parts) == 3 {
				p.filters = append(p.filters, filter{field: parts[0], op: parts[1], value: parts[2]})
			}
		case "-s":
			i++
			if i >= len(args) {
				break
			}
			parts := strings.SplitN(args[i], ":", 2)
			if len(parts) == 2 {
				p.filters = append(p.filters, filter{field: parts[0], op: "like", value: parts[1]})
			} else {
				// If no field specified, this is a bare search term — each command
				// decides which field to apply it to.
				p.filters = append(p.filters, filter{field: "", op: "like", value: parts[0]})
			}
		default:
			p.positional = append(p.positional, args[i])
		}
	}
	return p
}

// filter is a parsed filter flag from the CLI.
type filter struct {
	field string
	op    string
	value string
}

// applyFilters adds parsed filter flags to a Query. If a filter has an empty
// field name, defaultField is used instead.
func applyFilters(q *vantagepoint.Query, filters []filter, defaultField string) {
	for _, f := range filters {
		field := f.field
		if field == "" {
			field = defaultField
		}
		q.Filter(field, f.op, f.value)
	}
}

// printJSON writes v as indented JSON to stdout.
func printJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
