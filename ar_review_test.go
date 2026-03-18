package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestARBalanceJSONRoundtrip verifies that the ARBalance struct marshals and
// unmarshals correctly with the expected JSON field names matching the
// Vantagepoint API.
func TestARBalanceJSONRoundtrip(t *testing.T) {
	original := ARBalance{
		WBS1:            "PRJ-100",
		WBS1Name:        "Project Alpha",
		Invoice:         "INV-2026-001",
		InvoiceDate:     "2026-01-15T00:00:00",
		Balance:         5000.50,
		OriginalAmt:     7500.00,
		BillingClientID: "CLI-042",
		BillingClient:   "Acme Corp",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling ARBalance: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"WBS1":            "PRJ-100",
		"WBS1Name":        "Project Alpha",
		"Invoice":         "INV-2026-001",
		"InvoiceDate":     "2026-01-15T00:00:00",
		"Balance":         5000.50,
		"OriginalAmt":     7500.00,
		"BillingClientID": "CLI-042",
		"BillingClient":   "Acme Corp",
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

	var decoded ARBalance
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to ARBalance: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestARBalanceDetailCRJSONRoundtrip verifies that the ARBalanceDetailCR
// struct marshals and unmarshals correctly with the expected JSON field names
// matching the Vantagepoint API.
func TestARBalanceDetailCRJSONRoundtrip(t *testing.T) {
	original := ARBalanceDetailCR{
		WBS1:            "PRJ-200",
		WBS2:            "PHASE-A",
		WBS3:            "TASK-01",
		Invoice:         "INV-2026-005",
		CashReceiptDate: "2026-02-01T00:00:00",
		Account:         "1200",
		TransType:       "CR",
		SubType:         "P",
		Amount:          3500.75,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling ARBalanceDetailCR: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"WBS1":            "PRJ-200",
		"WBS2":            "PHASE-A",
		"WBS3":            "TASK-01",
		"Invoice":         "INV-2026-005",
		"CashReceiptDate": "2026-02-01T00:00:00",
		"Account":         "1200",
		"TransType":       "CR",
		"SubType":         "P",
		"Amount":          3500.75,
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

	var decoded ARBalanceDetailCR
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to ARBalanceDetailCR: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}
