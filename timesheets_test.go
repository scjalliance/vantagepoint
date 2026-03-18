package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestTimesheetControlRoundtrip verifies that TimesheetControl can be marshaled
// to JSON and unmarshaled back without data loss.
func TestTimesheetControlRoundtrip(t *testing.T) {
	original := TimesheetControl{
		Batch:                  "TS-001",
		Description:            "Weekly timesheet batch",
		Recurring:              "N",
		StartDate:              "2025-06-01",
		EndDate:                "2025-06-07",
		RegHrsTotal:            40.0,
		SumRegHrsTotal:         40.0,
		DiffRegHrsTotal:        0.0,
		OvtHrsTotal:            5.5,
		SumOvtHrsTotal:         5.5,
		DiffOvtHrsTotal:        0.0,
		SpecialOvtHrsTotal:     2.0,
		SumSpecialOvtHrsTotal:  2.0,
		DiffSpecialOvtHrsTotal: 0.0,
		Selected:               "Y",
		Posted:                 "N",
		Creator:                "admin",
		Period:                 202506,
		Company:                "ACME",
		Diary:                  "TS",
		DiaryNo:                42,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded TimesheetControl
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestTimesheetLandingRoundtrip verifies that TimesheetLanding can be marshaled
// to JSON and unmarshaled back without data loss.
func TestTimesheetLandingRoundtrip(t *testing.T) {
	original := TimesheetLanding{
		Batch:               "TS-001",
		Description:         "June 2025 timesheets",
		Recurring:           "N",
		EndDate:             "2025-06-30",
		Creator:             "admin",
		Period:              "06/2025",
		RawPeriod:           202506,
		Company:             "ACME",
		TransType:           "TS",
		ApprovalStatus:      "A",
		ApprovalStatusDesc:  "Approved",
		AutoPost:            "N",
		PostComment:         "Batch posted",
		PostDate:            "2025-07-01",
		PostLogReportName:   "PostLog.rpt",
		PostLogReportPath:   "/reports/PostLog.rpt",
		PostSeq:             1,
		PostStatus:          "P",
		PostStatusDesc:      "Posted",
		PostUser:            "admin",
		RecurSchedule:       "",
		Reverse:             "N",
		Reversed:            "N",
		TransListReportPath: "/reports/TransList.rpt",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded TimesheetLanding
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestTimesheetDetailRoundtrip verifies that TimesheetDetail can be marshaled
// to JSON and unmarshaled back without data loss.
func TestTimesheetDetailRoundtrip(t *testing.T) {
	original := TimesheetDetail{
		Batch:         "TS-001",
		Employee:      "EMP100",
		PKey:          "TS-001|EMP100|1",
		Seq:           1,
		TransDate:     "2025-06-02",
		WBS1:          "PRJ-001",
		WBS1Name:      "Main Project",
		ClientName:    "Acme Corp",
		WBS2Level:     "Phase",
		WBS1Locale:    "US",
		WBS2:          "PRJ-001.100",
		WBS2Name:      "Design Phase",
		WBS3Level:     "Task",
		WBS2Locale:    "US",
		WBS3:          "PRJ-001.100.01",
		WBS3Name:      "Schematic Design",
		WBS3Locale:    "US",
		Org:           "ARCH",
		LaborCode:     "DES",
		LCName:        "Design",
		BillCategory:  "T&M",
		Locale:        "US",
		RegHrs:        8.0,
		OvtHrs:        1.5,
		SpecialOvtHrs: 0.0,
		TransComment:  "Schematic design review",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded TimesheetDetail
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestTimesheetMasterRoundtrip verifies that TimesheetMaster can be marshaled
// to JSON and unmarshaled back without data loss.
func TestTimesheetMasterRoundtrip(t *testing.T) {
	original := TimesheetMaster{
		Batch:             "TS-001",
		Employee:          "EMP100",
		EmplName:          "Jane Doe",
		BillingCategory:   "Professional",
		LocaleMethod:      "E",
		Locale:            "US",
		Posted:            "N",
		Seq:               1,
		Status:            "S",
		StatusDescription: "Submitted",
		AuthorizedByName:  "John Smith",
		AuthorizedBy:      "EMP200",
		RejectReason:      "",
		ModUser:           "admin",
		ModUserEmployee:   "EMP001",
		ModDate:           "2025-06-03",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded TimesheetMaster
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestTimesheetDetailOmitsZeroValues verifies that zero-valued fields are
// omitted from JSON output due to the omitempty tag.
func TestTimesheetDetailOmitsZeroValues(t *testing.T) {
	detail := TimesheetDetail{
		Batch:    "TS-001",
		Employee: "EMP100",
		RegHrs:   8.0,
	}

	data, err := json.Marshal(detail)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshal to map: %v", err)
	}

	// Fields with zero values should be omitted.
	for _, key := range []string{"OvtHrs", "SpecialOvtHrs", "Seq", "WBS1", "TransComment"} {
		if _, ok := raw[key]; ok {
			t.Errorf("expected field %q to be omitted, but it was present", key)
		}
	}

	// Non-zero fields should be present.
	for _, key := range []string{"Batch", "Employee", "RegHrs"} {
		if _, ok := raw[key]; !ok {
			t.Errorf("expected field %q to be present, but it was omitted", key)
		}
	}
}

// TestTimesheetControlFromJSON verifies that TimesheetControl can be decoded
// from a JSON object with a subset of fields, simulating a real API response.
func TestTimesheetControlFromJSON(t *testing.T) {
	input := `{
		"Batch": "TS-042",
		"Description": "Week of June 2",
		"Period": 202506,
		"RegHrsTotal": 37.5,
		"OvtHrsTotal": 3.0,
		"Posted": "N"
	}`

	var ctrl TimesheetControl
	if err := json.Unmarshal([]byte(input), &ctrl); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if ctrl.Batch != "TS-042" {
		t.Errorf("Batch: got %q, want %q", ctrl.Batch, "TS-042")
	}
	if ctrl.Period != 202506 {
		t.Errorf("Period: got %d, want %d", ctrl.Period, 202506)
	}
	if ctrl.RegHrsTotal != 37.5 {
		t.Errorf("RegHrsTotal: got %f, want %f", ctrl.RegHrsTotal, 37.5)
	}
	if ctrl.OvtHrsTotal != 3.0 {
		t.Errorf("OvtHrsTotal: got %f, want %f", ctrl.OvtHrsTotal, 3.0)
	}
	if ctrl.Posted != "N" {
		t.Errorf("Posted: got %q, want %q", ctrl.Posted, "N")
	}
	// Unset fields should be zero values.
	if ctrl.StartDate != "" {
		t.Errorf("StartDate: got %q, want empty", ctrl.StartDate)
	}
	if ctrl.DiaryNo != 0 {
		t.Errorf("DiaryNo: got %d, want 0", ctrl.DiaryNo)
	}
}
