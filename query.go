package vantagepoint

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Query builds query parameters for Vantagepoint API requests, including
// pagination, field selection, filtering, sorting, and WBS type selection.
type Query struct {
	limit     int
	offset    int
	page      int
	pageSize  int
	fields    []string
	filters   []filter
	order     string
	wbsType   string
	postAsGet bool
	noLimit   bool
}

// filter represents a single filter condition in the filterHash query system.
type filter struct {
	Name      string
	Value     string
	TableName string
	Opp       string
}

// NewQuery creates a new empty Query builder.
func NewQuery() *Query {
	return &Query{}
}

// Limit sets the maximum number of records to return.
func (q *Query) Limit(n int) *Query {
	q.limit = n
	return q
}

// Offset sets the number of records to skip before returning results.
func (q *Query) Offset(n int) *Query {
	q.offset = n
	return q
}

// Page sets the page number for paginated results.
func (q *Query) Page(n int) *Query {
	q.page = n
	return q
}

// PageSize sets the number of records per page.
func (q *Query) PageSize(n int) *Query {
	q.pageSize = n
	return q
}

// Fields specifies which fields to include in the response. This maps to the
// fieldFilter query parameter as a comma-separated list.
func (q *Query) Fields(fields ...string) *Query {
	q.fields = append(q.fields, fields...)
	return q
}

// Filter adds a filter condition. The opp parameter specifies the comparison
// operator (e.g., "eq", "ne", "gt", "lt", "ge", "le", "like", "in").
func (q *Query) Filter(name, opp, value string) *Query {
	q.filters = append(q.filters, filter{
		Name:  name,
		Opp:   opp,
		Value: value,
	})
	return q
}

// FilterWithTable adds a filter condition that targets a specific table name.
// This is used when filtering on fields that belong to related tables.
func (q *Query) FilterWithTable(name, opp, value, tableName string) *Query {
	q.filters = append(q.filters, filter{
		Name:      name,
		Opp:       opp,
		Value:     value,
		TableName: tableName,
	})
	return q
}

// OrderBy sets the sort order for results. Use the field name optionally
// followed by " asc" or " desc" (e.g., "Name asc", "WBS1 desc").
func (q *Query) OrderBy(field string) *Query {
	q.order = field
	return q
}

// WBSType sets the WBS type filter. Valid values are "wbs1" (project),
// "wbs2" (phase), or "wbs3" (task).
func (q *Query) WBSType(t string) *Query {
	q.wbsType = t
	return q
}

// MaxRowsOnly configures the query to use only the q.maxrows parameter for
// row limiting instead of the standard limit parameter. Some Vantagepoint
// endpoints (e.g., PSALedger, DataEntry) reject the limit parameter with a 500
// error and only accept q.maxrows.
func (q *Query) MaxRowsOnly() *Query {
	q.noLimit = true
	return q
}

// PostAsGet marks the query to be sent as a POST request with the
// "dvp-postAsGet: true" header. This is useful for queries with parameters
// that exceed URL length limits.
func (q *Query) PostAsGet() *Query {
	q.postAsGet = true
	return q
}

// Values encodes the query parameters into url.Values suitable for use as
// query string parameters on an API request.
func (q *Query) Values() url.Values {
	v := url.Values{}

	if q.limit > 0 {
		s := strconv.Itoa(q.limit)
		if !q.noLimit {
			v.Set("limit", s)
		}
		v.Set("q.maxrows", s)
	}
	if q.offset > 0 {
		v.Set("offset", strconv.Itoa(q.offset))
	}
	if q.page > 0 {
		v.Set("page", strconv.Itoa(q.page))
	}
	if q.pageSize > 0 {
		v.Set("pageSize", strconv.Itoa(q.pageSize))
	}
	if len(q.fields) > 0 {
		v.Set("fieldFilter", strings.Join(q.fields, ","))
	}
	if q.order != "" {
		v.Set("order", q.order)
	}
	if q.wbsType != "" {
		v.Set("wbstype", q.wbsType)
	}

	for i, f := range q.filters {
		prefix := fmt.Sprintf("filterHash[%d]", i)
		v.Set(prefix+"[name]", f.Name)
		v.Set(prefix+"[value]", f.Value)
		v.Set(prefix+"[opp]", f.Opp)
		if f.TableName != "" {
			v.Set(prefix+"[tablename]", f.TableName)
		}
	}

	return v
}

// NeedsPostAsGet returns true if the query was marked for POST-as-GET mode.
func (q *Query) NeedsPostAsGet() bool {
	return q.postAsGet
}
