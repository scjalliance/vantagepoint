package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestUnitControlJSONRoundtrip verifies that UnitControl can be marshaled
// to JSON and unmarshaled back without data loss.
func TestUnitControlJSONRoundtrip(t *testing.T) {
	original := UnitControl{
		Batch:       "UN-001",
		Description: "Weekly unit batch",
		Total:       150.0,
		SumTotal:    150.0,
		DiffTotal:   0.0,
		Period:      202506,
		Company:     "ACME",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded UnitControl
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestUnitDetailJSONRoundtrip verifies that UnitDetail can be marshaled
// to JSON and unmarshaled back without data loss.
func TestUnitDetailJSONRoundtrip(t *testing.T) {
	original := UnitDetail{
		Batch:       "UN-001",
		Unit:        "UNIT-A",
		PKey:        "UN-001|UNIT-A|1",
		WBS1:        "PRJ-001",
		WBS2:        "PRJ-001.100",
		Quantity:    25.5,
		CostRate:    75.00,
		BillingRate: 125.00,
		TableNo:     "TBL-01",
		Account:     "5000",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded UnitDetail
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}
