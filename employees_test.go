package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestEmployeeJSONRoundtrip verifies that the Employee struct marshals and
// unmarshals with the correct JSON field names matching the Vantagepoint API.
// Key field names verified: EMail (not Email), PreferredName (not PreferName),
// MobilePhone (not CellPhone), Org (not OrgCode), Status as string (not IsActive
// bool), FAX, ZIP.
func TestEmployeeJSONRoundtrip(t *testing.T) {
	original := Employee{
		Employee:      "EMP001",
		FirstName:     "Jane",
		LastName:      "Doe",
		PreferredName: "Janie",
		EMail:         "jane.doe@example.com",
		MobilePhone:   "555-0100",
		Org:           "ENG",
		Status:        "A",
		FAX:           "555-0199",
		ZIP:           "98101",
		PayRate:       85000.00,
		HoursPerDay:   8.0,
		HasPhoto:      1,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling Employee: %v", err)
	}

	// Verify the JSON uses the correct VP field names.
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"Employee":      "EMP001",
		"FirstName":     "Jane",
		"LastName":      "Doe",
		"PreferredName": "Janie",
		"EMail":         "jane.doe@example.com",
		"MobilePhone":   "555-0100",
		"Org":           "ENG",
		"Status":        "A",
		"FAX":           "555-0199",
		"ZIP":           "98101",
		"PayRate":       85000.00,
		"HoursPerDay":   8.0,
		"HasPhoto":      float64(1),
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
	badKeys := []string{"Email", "PreferName", "CellPhone", "OrgCode", "IsActive"}
	for _, key := range badKeys {
		if _, ok := raw[key]; ok {
			t.Errorf("unexpected old JSON key %q found in output", key)
		}
	}

	// Verify roundtrip: unmarshal back into an Employee and compare.
	var decoded Employee
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to Employee: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestEmployeeJSONFromAPI verifies that a JSON payload using API field names
// correctly deserializes into the Employee struct.
func TestEmployeeJSONFromAPI(t *testing.T) {
	apiJSON := `{
		"Employee": "EMP042",
		"FirstName": "John",
		"LastName": "Smith",
		"PreferredName": "Johnny",
		"EMail": "john.smith@example.com",
		"MobilePhone": "555-0200",
		"Org": "ARCH",
		"Status": "A",
		"FAX": "555-0299",
		"ZIP": "10001",
		"PayRate": 95000.50,
		"HoursPerDay": 7.5,
		"Supervisor": "EMP001",
		"SupervisorName": "Jane Doe",
		"HireDate": "2020-06-15T00:00:00",
		"HasPhoto": 0,
		"BillingCategory": 2,
		"UtilizationRatio": 0.85
	}`

	var e Employee
	if err := json.Unmarshal([]byte(apiJSON), &e); err != nil {
		t.Fatalf("unmarshaling API JSON: %v", err)
	}

	if e.Employee != "EMP042" {
		t.Errorf("Employee: got %q, want %q", e.Employee, "EMP042")
	}
	if e.PreferredName != "Johnny" {
		t.Errorf("PreferredName: got %q, want %q", e.PreferredName, "Johnny")
	}
	if e.EMail != "john.smith@example.com" {
		t.Errorf("EMail: got %q, want %q", e.EMail, "john.smith@example.com")
	}
	if e.MobilePhone != "555-0200" {
		t.Errorf("MobilePhone: got %q, want %q", e.MobilePhone, "555-0200")
	}
	if e.Org != "ARCH" {
		t.Errorf("Org: got %q, want %q", e.Org, "ARCH")
	}
	if e.Status != "A" {
		t.Errorf("Status: got %q, want %q", e.Status, "A")
	}
	if e.FAX != "555-0299" {
		t.Errorf("FAX: got %q, want %q", e.FAX, "555-0299")
	}
	if e.ZIP != "10001" {
		t.Errorf("ZIP: got %q, want %q", e.ZIP, "10001")
	}
	if e.PayRate != 95000.50 {
		t.Errorf("PayRate: got %v, want %v", e.PayRate, 95000.50)
	}
	if e.UtilizationRatio != 0.85 {
		t.Errorf("UtilizationRatio: got %v, want %v", e.UtilizationRatio, 0.85)
	}
}

// TestEmployeeOmitsEmptyFields verifies that zero-value fields are omitted
// from JSON output due to the omitempty tags.
func TestEmployeeOmitsEmptyFields(t *testing.T) {
	e := Employee{
		Employee: "EMP001",
		LastName: "Doe",
	}

	data, err := json.Marshal(e)
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
	if raw["Employee"] != "EMP001" {
		t.Errorf("Employee: got %v, want %q", raw["Employee"], "EMP001")
	}
	if raw["LastName"] != "Doe" {
		t.Errorf("LastName: got %v, want %q", raw["LastName"], "Doe")
	}
}

// TestEmployeeStatusIsString verifies that Status is a string field (e.g.,
// "A", "I", "T") rather than a boolean IsActive field.
func TestEmployeeStatusIsString(t *testing.T) {
	for _, status := range []string{"A", "I", "T"} {
		e := Employee{Employee: "EMP001", Status: status}
		data, err := json.Marshal(e)
		if err != nil {
			t.Fatalf("marshaling with Status=%q: %v", status, err)
		}

		var raw map[string]any
		if err := json.Unmarshal(data, &raw); err != nil {
			t.Fatalf("unmarshaling to map: %v", err)
		}

		got, ok := raw["Status"].(string)
		if !ok {
			t.Errorf("Status should be a string, got %T", raw["Status"])
		}
		if got != status {
			t.Errorf("Status: got %q, want %q", got, status)
		}
	}
}

func TestEmployeeProjectJSONRoundtrip(t *testing.T) {
	original := EmployeeProject{
		RecordID:        "604A000A6867428E8C151530C3B73003",
		WBS1:            "2003005.00",
		Employee:        "00201",
		Role:            "8",
		RoleDescription: "Design Lead",
		StartDate:       "2003-06-18T14:10:00.753",
	}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling EmployeeProject: %v", err)
	}
	var decoded EmployeeProject
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling EmployeeProject: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}
