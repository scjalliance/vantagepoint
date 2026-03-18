package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestUnitByProjectControlJSONRoundtrip verifies that UnitByProjectControl can
// be marshaled to JSON and unmarshaled back without data loss, and that the
// JSON field names match the Vantagepoint API expectations.
func TestUnitByProjectControlJSONRoundtrip(t *testing.T) {
	original := UnitByProjectControl{
		Batch:       "UP-001",
		Description: "Monthly unit-by-project batch",
		Total:       1500.75,
		SumTotal:    1500.75,
		DiffTotal:   0.0,
		Period:      202603,
		Company:     "ACME",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshal to map: %v", err)
	}

	// Spot-check key JSON field names.
	checks := map[string]any{
		"Batch":       "UP-001",
		"Description": "Monthly unit-by-project batch",
		"Total":       1500.75,
		"SumTotal":    1500.75,
		"Period":      float64(202603),
		"Company":     "ACME",
	}

	for key, want := range checks {
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

	var decoded UnitByProjectControl
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal back to UnitByProjectControl: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestUnitByProjectDetailJSONRoundtrip verifies that UnitByProjectDetail can be
// marshaled to JSON and unmarshaled back without data loss, covering fields
// specific to the units-by-project variant including both RefNo and Unit.
func TestUnitByProjectDetailJSONRoundtrip(t *testing.T) {
	original := UnitByProjectDetail{
		Batch:    "UP-001",
		RefNo:    "REF-100",
		PKey:     "UP-001|REF-100|1",
		WBS1:     "PRJ-200",
		WBS2:     "PHASE-A",
		Unit:     "HR",
		UnitName: "Hours",
		Quantity: 40.0,
		CostRate: 125.50,
		Employee: "EMP042",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshal to map: %v", err)
	}

	// Spot-check key JSON field names.
	checks := map[string]any{
		"Batch":    "UP-001",
		"RefNo":    "REF-100",
		"PKey":     "UP-001|REF-100|1",
		"WBS1":     "PRJ-200",
		"WBS2":     "PHASE-A",
		"Unit":     "HR",
		"UnitName": "Hours",
		"Quantity": 40.0,
		"CostRate": 125.50,
		"Employee": "EMP042",
	}

	for key, want := range checks {
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

	var decoded UnitByProjectDetail
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal back to UnitByProjectDetail: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}
