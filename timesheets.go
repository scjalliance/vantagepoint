package vantagepoint

import (
	"context"
	"fmt"
)

// TimesheetLanding represents a batch summary returned by the tsLanding endpoint.
// This is a different schema from TimesheetControl — they share only a few fields,
// and Period is a string here (e.g. "06/2008") versus an int in TimesheetControl.
type TimesheetLanding struct {
	Batch               string `json:"Batch,omitempty"`
	Description         string `json:"Description,omitempty"`
	Recurring           string `json:"Recurring,omitempty"`
	EndDate             string `json:"EndDate,omitempty"`
	Creator             string `json:"Creator,omitempty"`
	Period              string `json:"Period,omitempty"`
	RawPeriod           int    `json:"RawPeriod,omitempty"`
	Company             string `json:"Company,omitempty"`
	TransType           string `json:"TransType,omitempty"`
	ApprovalStatus      string `json:"ApprovalStatus,omitempty"`
	ApprovalStatusDesc  string `json:"ApprovalStatusDesc,omitempty"`
	AutoPost            string `json:"AutoPost,omitempty"`
	PostComment         string `json:"PostComment,omitempty"`
	PostDate            string `json:"PostDate,omitempty"`
	PostLogReportName   string `json:"PostLogReportName,omitempty"`
	PostLogReportPath   string `json:"PostLogReportPath,omitempty"`
	PostSeq             int    `json:"PostSeq,omitempty"`
	PostStatus          string `json:"PostStatus,omitempty"`
	PostStatusDesc      string `json:"PostStatusDesc,omitempty"`
	PostUser            string `json:"PostUser,omitempty"`
	RecurSchedule       string `json:"RecurSchedule,omitempty"`
	Reverse             string `json:"Reverse,omitempty"`
	Reversed            string `json:"Reversed,omitempty"`
	TransListReportPath string `json:"TransListReportPath,omitempty"`
}

// TimesheetControl represents a timesheet batch control record.
type TimesheetControl struct {
	Batch                  string  `json:"Batch,omitempty"`
	Description            string  `json:"Description,omitempty"`
	Recurring              string  `json:"Recurring,omitempty"`
	StartDate              string  `json:"StartDate,omitempty"`
	EndDate                string  `json:"EndDate,omitempty"`
	RegHrsTotal            float64 `json:"RegHrsTotal,omitempty"`
	SumRegHrsTotal         float64 `json:"SumRegHrsTotal,omitempty"`
	DiffRegHrsTotal        float64 `json:"DiffRegHrsTotal,omitempty"`
	OvtHrsTotal            float64 `json:"OvtHrsTotal,omitempty"`
	SumOvtHrsTotal         float64 `json:"SumOvtHrsTotal,omitempty"`
	DiffOvtHrsTotal        float64 `json:"DiffOvtHrsTotal,omitempty"`
	SpecialOvtHrsTotal     float64 `json:"SpecialOvtHrsTotal,omitempty"`
	SumSpecialOvtHrsTotal  float64 `json:"SumSpecialOvtHrsTotal,omitempty"`
	DiffSpecialOvtHrsTotal float64 `json:"DiffSpecialOvtHrsTotal,omitempty"`
	Selected               string  `json:"Selected,omitempty"`
	Posted                 string  `json:"Posted,omitempty"`
	Creator                string  `json:"Creator,omitempty"`
	Period                 int     `json:"Period,omitempty"`
	Company                string  `json:"Company,omitempty"`
	Diary                  string  `json:"Diary,omitempty"`
	DiaryNo                int     `json:"DiaryNo,omitempty"`
}

// TimesheetMaster represents per-employee timesheet data within a batch.
type TimesheetMaster struct {
	Batch             string `json:"Batch,omitempty"`
	Employee          string `json:"Employee,omitempty"`
	EmplName          string `json:"EmplName,omitempty"`
	BillingCategory   string `json:"BillingCategory,omitempty"`
	LocaleMethod      string `json:"LocaleMethod,omitempty"`
	Locale            string `json:"Locale,omitempty"`
	Posted            string `json:"Posted,omitempty"`
	Seq               int    `json:"Seq,omitempty"`
	Status            string `json:"Status,omitempty"`
	StatusDescription string `json:"StatusDescription,omitempty"`
	AuthorizedByName  string `json:"AuthorizedByName,omitempty"`
	AuthorizedBy      string `json:"AuthorizedBy,omitempty"`
	RejectReason      string `json:"RejectReason,omitempty"`
	ModUser           string `json:"ModUser,omitempty"`
	ModUserEmployee   string `json:"ModUserEmployee,omitempty"`
	ModDate           string `json:"ModDate,omitempty"`
}

// TimesheetDetail represents an individual time entry line.
type TimesheetDetail struct {
	Batch         string  `json:"Batch,omitempty"`
	Employee      string  `json:"Employee,omitempty"`
	PKey          string  `json:"PKey,omitempty"`
	Seq           int     `json:"Seq,omitempty"`
	TransDate     string  `json:"TransDate,omitempty"`
	WBS1          string  `json:"WBS1,omitempty"`
	WBS1Name      string  `json:"WBS1Name,omitempty"`
	ClientName    string  `json:"ClientName,omitempty"`
	WBS2Level     string  `json:"WBS2Level,omitempty"`
	WBS1Locale    string  `json:"WBS1Locale,omitempty"`
	WBS2          string  `json:"WBS2,omitempty"`
	WBS2Name      string  `json:"WBS2Name,omitempty"`
	WBS3Level     string  `json:"WBS3Level,omitempty"`
	WBS2Locale    string  `json:"WBS2Locale,omitempty"`
	WBS3          string  `json:"WBS3,omitempty"`
	WBS3Name      string  `json:"WBS3Name,omitempty"`
	WBS3Locale    string  `json:"WBS3Locale,omitempty"`
	Org           string  `json:"Org,omitempty"`
	LaborCode     string  `json:"LaborCode,omitempty"`
	LCName        string  `json:"LCName,omitempty"`
	BillCategory  string  `json:"BillCategory,omitempty"`
	Locale        string  `json:"Locale,omitempty"`
	RegHrs        float64 `json:"RegHrs,omitempty"`
	OvtHrs        float64 `json:"OvtHrs,omitempty"`
	SpecialOvtHrs float64 `json:"SpecialOvtHrs,omitempty"`
	TransComment  string  `json:"TransComment,omitempty"`
}

// ListTimesheetBatches retrieves a list of timesheet batch summaries from the
// DataEntry tsLanding endpoint. Pass a *Query to filter, paginate, or sort
// results; nil retrieves defaults.
func (c *Client) ListTimesheetBatches(ctx context.Context, q *Query) ([]TimesheetLanding, error) {
	var results []TimesheetLanding
	if err := c.get(ctx, "DataEntry/tsLanding", q, &results); err != nil {
		return nil, fmt.Errorf("listing timesheet batches: %w", err)
	}
	return results, nil
}

// GetTimesheetControl retrieves the control records for a timesheet batch
// identified by the given batch key.
func (c *Client) GetTimesheetControl(ctx context.Context, batch string, q *Query) ([]TimesheetControl, error) {
	var results []TimesheetControl
	if err := c.get(ctx, "DataEntry/tsControl/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting timesheet control for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetTimesheetMaster retrieves the per-employee master records for a timesheet
// batch identified by the given batch key.
func (c *Client) GetTimesheetMaster(ctx context.Context, batch string, q *Query) ([]TimesheetMaster, error) {
	var results []TimesheetMaster
	if err := c.get(ctx, "DataEntry/tsMaster/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting timesheet master for batch %s: %w", batch, err)
	}
	return results, nil
}

// GetTimesheetDetail retrieves the individual time entry lines for a timesheet
// batch identified by the given batch key.
func (c *Client) GetTimesheetDetail(ctx context.Context, batch string, q *Query) ([]TimesheetDetail, error) {
	var results []TimesheetDetail
	if err := c.get(ctx, "DataEntry/tsDetail/"+batch, q, &results); err != nil {
		return nil, fmt.Errorf("getting timesheet detail for batch %s: %w", batch, err)
	}
	return results, nil
}

// CreateTimesheetBatch creates a new timesheet batch via the tsControl endpoint.
func (c *Client) CreateTimesheetBatch(ctx context.Context, control *TimesheetControl) ([]TimesheetControl, error) {
	var results []TimesheetControl
	if err := c.post(ctx, "DataEntry/tsControl", control, &results); err != nil {
		return nil, fmt.Errorf("creating timesheet batch: %w", err)
	}
	return results, nil
}

// UpdateTimesheetBatch updates an existing timesheet batch control record
// identified by the given batch key.
func (c *Client) UpdateTimesheetBatch(ctx context.Context, batch string, control *TimesheetControl) ([]TimesheetControl, error) {
	var results []TimesheetControl
	if err := c.put(ctx, "DataEntry/tsControl/"+batch, control, &results); err != nil {
		return nil, fmt.Errorf("updating timesheet batch %s: %w", batch, err)
	}
	return results, nil
}

// DeleteTimesheetBatch deletes a timesheet batch via the tsControl endpoint.
func (c *Client) DeleteTimesheetBatch(ctx context.Context, batch string) error {
	if err := c.delete(ctx, "DataEntry/tsControl/"+batch); err != nil {
		return fmt.Errorf("deleting timesheet batch %s: %w", batch, err)
	}
	return nil
}

// CreateTimesheetMaster creates a per-employee master record within a timesheet
// batch. The batch and employee parameters form the composite key.
func (c *Client) CreateTimesheetMaster(ctx context.Context, batch, employee string, master *TimesheetMaster) ([]TimesheetMaster, error) {
	var results []TimesheetMaster
	if err := c.post(ctx, "DataEntry/tsMaster/"+batch+"|"+employee, master, &results); err != nil {
		return nil, fmt.Errorf("creating timesheet master for batch %s employee %s: %w", batch, employee, err)
	}
	return results, nil
}

// UpdateTimesheetMaster updates a per-employee master record within a timesheet
// batch. The batch and employee parameters form the composite key.
func (c *Client) UpdateTimesheetMaster(ctx context.Context, batch, employee string, master *TimesheetMaster) ([]TimesheetMaster, error) {
	var results []TimesheetMaster
	if err := c.put(ctx, "DataEntry/tsMaster/"+batch+"|"+employee, master, &results); err != nil {
		return nil, fmt.Errorf("updating timesheet master for batch %s employee %s: %w", batch, employee, err)
	}
	return results, nil
}

// DeleteTimesheetMaster deletes a per-employee master record from a timesheet
// batch. The batch and employee parameters form the composite key.
func (c *Client) DeleteTimesheetMaster(ctx context.Context, batch, employee string) error {
	if err := c.delete(ctx, "DataEntry/tsMaster/"+batch+"|"+employee); err != nil {
		return fmt.Errorf("deleting timesheet master for batch %s employee %s: %w", batch, employee, err)
	}
	return nil
}
