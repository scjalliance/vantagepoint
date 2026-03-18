package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestProjectJSONRoundtrip verifies that the Project struct marshals and
// unmarshals with the correct JSON field names matching the Vantagepoint API.
func TestProjectJSONRoundtrip(t *testing.T) {
	original := Project{
		WBS1:      "P001",
		WBS2:      "P001.01",
		WBS3:      "P001.01.001",
		WBSNumber: "P001.01.001",
		Name:      "Test Project",
		LongName:  "Test Project Long Name",
		SubLevel:  "wbs1",
		ProjMgr:   "MGR001",
		Org:       "ORG001",
		Fee:       50000.00,
		CLAddress: "ADDR001",
		StartDate: "2026-01-01T00:00:00",
		EndDate:   "2026-12-31T00:00:00",
		Status:    "Active",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling Project: %v", err)
	}

	// Verify the JSON uses the correct field names.
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"WBS1":      "P001",
		"WBS2":      "P001.01",
		"WBS3":      "P001.01.001",
		"WBSNumber": "P001.01.001",
		"Name":      "Test Project",
		"LongName":  "Test Project Long Name",
		"SubLevel":  "wbs1",
		"ProjMgr":   "MGR001",
		"Org":       "ORG001",
		"Fee":       50000.00,
		"CLAddress": "ADDR001",
		"StartDate": "2026-01-01T00:00:00",
		"EndDate":   "2026-12-31T00:00:00",
		"Status":    "Active",
	}

	for key, want := range expectedKeys {
		got, ok := raw[key]
		if !ok {
			t.Errorf("expected JSON key %q not found", key)
			continue
		}
		// json.Unmarshal turns numbers into float64.
		switch w := want.(type) {
		case float64:
			g, ok := got.(float64)
			if !ok || g != w {
				t.Errorf("key %q: got %v, want %v", key, got, want)
			}
		case string:
			g, ok := got.(string)
			if !ok || g != w {
				t.Errorf("key %q: got %v, want %v", key, got, want)
			}
		}
	}

	// Verify that old incorrect field names are NOT present.
	badKeys := []string{"ProjectManager", "OrgCode", "ContractAmount", "PrincipalCode", "SupervisorCode"}
	for _, key := range badKeys {
		if _, ok := raw[key]; ok {
			t.Errorf("unexpected old JSON key %q found in output", key)
		}
	}

	// Verify roundtrip: unmarshal back into a Project and compare.
	var decoded Project
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to Project: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestProjectJSONFromAPI verifies that a JSON payload using API field names
// correctly deserializes into the Project struct.
func TestProjectJSONFromAPI(t *testing.T) {
	apiJSON := `{
		"WBS1": "PRJ100",
		"WBS2": "PRJ100.10",
		"WBS3": "PRJ100.10.01",
		"SubLevel": "wbs3",
		"ProjMgr": "EMP042",
		"Org": "DIV05",
		"Fee": 125000.50,
		"CLAddress": "CA001",
		"Name": "Highway Bridge Design",
		"ChargeType": "Regular",
		"StartDate": "2026-03-01T00:00:00",
		"Closed": 0,
		"VersionID": 3,
		"ProjectCurrencyCode": "USD",
		"ProjectExchangeRate": 1.0
	}`

	var p Project
	if err := json.Unmarshal([]byte(apiJSON), &p); err != nil {
		t.Fatalf("unmarshaling API JSON: %v", err)
	}

	if p.WBS1 != "PRJ100" {
		t.Errorf("WBS1: got %q, want %q", p.WBS1, "PRJ100")
	}
	if p.WBS2 != "PRJ100.10" {
		t.Errorf("WBS2: got %q, want %q", p.WBS2, "PRJ100.10")
	}
	if p.WBS3 != "PRJ100.10.01" {
		t.Errorf("WBS3: got %q, want %q", p.WBS3, "PRJ100.10.01")
	}
	if p.SubLevel != "wbs3" {
		t.Errorf("SubLevel: got %q, want %q", p.SubLevel, "wbs3")
	}
	if p.ProjMgr != "EMP042" {
		t.Errorf("ProjMgr: got %q, want %q", p.ProjMgr, "EMP042")
	}
	if p.Org != "DIV05" {
		t.Errorf("Org: got %q, want %q", p.Org, "DIV05")
	}
	if p.Fee != 125000.50 {
		t.Errorf("Fee: got %v, want %v", p.Fee, 125000.50)
	}
	if p.CLAddress != "CA001" {
		t.Errorf("CLAddress: got %q, want %q", p.CLAddress, "CA001")
	}
	if p.VersionID != 3 {
		t.Errorf("VersionID: got %d, want %d", p.VersionID, 3)
	}
	if p.ProjectExchangeRate != 1.0 {
		t.Errorf("ProjectExchangeRate: got %v, want %v", p.ProjectExchangeRate, 1.0)
	}
}

// TestProjectOmitsEmptyFields verifies that zero-value fields are omitted
// from JSON output due to the omitempty tags.
func TestProjectOmitsEmptyFields(t *testing.T) {
	p := Project{
		WBS1: "P001",
		Name: "Minimal",
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshaling: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	if len(raw) != 2 {
		t.Errorf("expected 2 JSON keys, got %d: %v", len(raw), raw)
	}
	if raw["WBS1"] != "P001" {
		t.Errorf("WBS1: got %v, want %q", raw["WBS1"], "P001")
	}
	if raw["Name"] != "Minimal" {
		t.Errorf("Name: got %v, want %q", raw["Name"], "Minimal")
	}
}

// TestProjectEmployeeJSONRoundtrip verifies that ProjectEmployee marshals and
// unmarshals with the correct JSON field names.
func TestProjectEmployeeJSONRoundtrip(t *testing.T) {
	original := ProjectEmployee{
		RecordID:   "REC001",
		WBS1:       "P001",
		Employee:   "EMP001",
		Role:       "PM",
		TeamStatus: "Active",
		CRMHours:   120.5,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling ProjectEmployee: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"RecordID":   "REC001",
		"WBS1":       "P001",
		"Employee":   "EMP001",
		"Role":       "PM",
		"TeamStatus": "Active",
		"CRMHours":   120.5,
	}

	for key, want := range expectedKeys {
		got, ok := raw[key]
		if !ok {
			t.Errorf("expected JSON key %q not found", key)
			continue
		}
		switch w := want.(type) {
		case float64:
			g, ok := got.(float64)
			if !ok || g != w {
				t.Errorf("key %q: got %v, want %v", key, got, want)
			}
		case string:
			g, ok := got.(string)
			if !ok || g != w {
				t.Errorf("key %q: got %v, want %v", key, got, want)
			}
		}
	}

	var decoded ProjectEmployee
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to ProjectEmployee: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestProjectRevenueJSONRoundtrip verifies that ProjectRevenue marshals and
// unmarshals with the correct JSON field names.
func TestProjectRevenueJSONRoundtrip(t *testing.T) {
	original := ProjectRevenue{
		RevAllocID:     "REV001",
		WBS1:           "P001",
		Description:    "Phase 1 Revenue",
		RevenueDate:    "2026-06-15T00:00:00",
		RevenueAmt:     75000.00,
		PercentRevenue: 50.0,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling ProjectRevenue: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"RevAllocID":     "REV001",
		"WBS1":           "P001",
		"Description":    "Phase 1 Revenue",
		"RevenueDate":    "2026-06-15T00:00:00",
		"RevenueAmt":     75000.00,
		"PercentRevenue": 50.0,
	}

	for key, want := range expectedKeys {
		got, ok := raw[key]
		if !ok {
			t.Errorf("expected JSON key %q not found", key)
			continue
		}
		switch w := want.(type) {
		case float64:
			g, ok := got.(float64)
			if !ok || g != w {
				t.Errorf("key %q: got %v, want %v", key, got, want)
			}
		case string:
			g, ok := got.(string)
			if !ok || g != w {
				t.Errorf("key %q: got %v, want %v", key, got, want)
			}
		}
	}

	var decoded ProjectRevenue
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to ProjectRevenue: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}
