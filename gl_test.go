package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestGLSummaryJSONRoundtrip verifies that the GLSummary struct marshals and
// unmarshals correctly with the expected JSON field names matching the
// Vantagepoint API.
func TestGLSummaryJSONRoundtrip(t *testing.T) {
	original := GLSummary{
		Account:             "5000",
		Name:                "Revenue",
		Period:              202601,
		Org:                 "ENG",
		Amount:              15000.50,
		CBAmount:            14500.25,
		CreditAmount:        15000.50,
		DebitAmount:         0,
		CBCreditAmount:      14500.25,
		CBDebitAmount:       0,
		TransType:           "L",
		Status:              "P",
		Type:                1,
		CashBasisAccount:    "5000CB",
		Detail:              "Y",
		FasAccount:          "5000F",
		AccountCurrencyCode: "USD",
		GlobalAccount:       "5000G",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling GLSummary: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"Account":             "5000",
		"Name":                "Revenue",
		"Period":              float64(202601),
		"Org":                 "ENG",
		"Amount":              15000.50,
		"CBAmount":            14500.25,
		"CreditAmount":        15000.50,
		"TransType":           "L",
		"Status":              "P",
		"Type":                float64(1),
		"CashBasisAccount":    "5000CB",
		"Detail":              "Y",
		"FasAccount":          "5000F",
		"AccountCurrencyCode": "USD",
		"GlobalAccount":       "5000G",
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

	var decoded GLSummary
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to GLSummary: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestLaborDetailJSONRoundtrip verifies that the LaborDetail struct marshals
// and unmarshals correctly, covering key fields across functional currency,
// project currency, billing currency, and employee currency amounts.
func TestLaborDetailJSONRoundtrip(t *testing.T) {
	original := LaborDetail{
		Period:                           202601,
		PostSeq:                          1,
		PKey:                             "LBR-001",
		WBS1:                             "PRJ-100",
		WBS2:                             "PHASE-A",
		WBS3:                             "TASK-01",
		LaborCode:                        "D01",
		Employee:                         "EMP042",
		TransType:                        "L",
		TransDate:                        "2026-01-15T00:00:00",
		Name:                             "John Smith",
		RegHrs:                           8.0,
		OvtHrs:                           2.0,
		SpecialOvtHrs:                    1.0,
		RegAmt:                           800.00,
		OvtAmt:                           300.00,
		SpecialOvtAmt:                    200.00,
		BillExt:                          1500.00,
		Rate:                             100.00,
		OvtPct:                           1.5,
		OvtRate:                          150.00,
		SpecialOvtPct:                    2.0,
		SpecialOvtRate:                   200.00,
		EmType:                           "R",
		Pool:                             1,
		Category:                         3,
		EmOrg:                            "ENG",
		PrOrg:                            "ENG",
		ChargeType:                       "D",
		RateType:                         "E",
		Comment:                          "Design review",
		BilledWBS1:                       "PRJ-100",
		BilledInvoice:                    "INV-2026-001",
		BilledPeriod:                     202602,
		RegAmtProjectCurrency:            800.00,
		RegAmtBillingCurrency:            850.00,
		RegAmtEmployeeCurrency:           800.00,
		EMCurrencyCode:                   "USD",
		Payrate:                          75.00,
		ModUser:                          "admin",
		AuthorizedBy:                     "EMP001",
		Rowversion:                       "AAAAAB==",
		RealizationAmountBillingCurrency: 1400.00,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling LaborDetail: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	// Spot-check key JSON field names.
	checks := map[string]any{
		"Period":                 float64(202601),
		"PostSeq":                float64(1),
		"PKey":                   "LBR-001",
		"WBS1":                   "PRJ-100",
		"Employee":               "EMP042",
		"RegHrs":                 8.0,
		"OvtHrs":                 2.0,
		"BillExt":                1500.00,
		"Rate":                   100.00,
		"EMCurrencyCode":         "USD",
		"RegAmtBillingCurrency":  850.00,
		"RegAmtEmployeeCurrency": 800.00,
		"Rowversion":             "AAAAAB==",
		"ModUser":                "admin",
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

	var decoded LaborDetail
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to LaborDetail: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestPSALedgerEntryJSONRoundtrip verifies that the PSALedgerEntry struct
// marshals and unmarshals correctly, covering key fields including multi-
// currency amounts, unit billing, and transfer fields.
func TestPSALedgerEntryJSONRoundtrip(t *testing.T) {
	original := PSALedgerEntry{
		Batch:                            "TS-2026-01",
		Period:                           202601,
		PostSeq:                          5,
		PKey:                             "PSA-001",
		WBS1:                             "PRJ-200",
		WBS2:                             "PHASE-B",
		WBS3:                             "TASK-02",
		Account:                          "5010",
		Org:                              "ARCH",
		TransType:                        "L",
		SubType:                          "R",
		RefNo:                            "REF-001",
		TransDate:                        "2026-01-20T00:00:00",
		Desc1:                            "Labor posting",
		Desc2:                            "Regular hours",
		Amount:                           1200.00,
		CBAmount:                         1150.00,
		BillExt:                          1800.00,
		ProjectCost:                      "Y",
		BillStatus:                       "B",
		Employee:                         "EMP042",
		Vendor:                           "VND-001",
		Line:                             1,
		PartialPayment:                   0,
		Discount:                         50.00,
		Voucher:                          "VCH-001",
		BilledWBS1:                       "PRJ-200",
		BilledPeriod:                     202602,
		UnitQuantity:                     8.0,
		UnitCostRate:                     150.00,
		UnitBillingRate:                  225.00,
		UnitBillExt:                      1800.00,
		XferWBS1:                         "PRJ-300",
		XferAccount:                      "5020",
		TaxCode:                          "TX01",
		TaxBasis:                         1200.00,
		TransactionAmount:                1200.00,
		TransactionCurrencyCode:          "USD",
		AmountProjectCurrency:            1200.00,
		AmountBillingCurrency:            1250.00,
		AutoEntryAmount:                  100.00,
		AmountSourceCurrency:             1200.00,
		PONumber:                         "PO-2026-100",
		LinkCompany:                      "ACME",
		LinkWBS1:                         "PRJ-200",
		Diary:                            "GL",
		DiaryNo:                          42,
		ModUser:                          "admin",
		AuthorizedBy:                     "EMP001",
		NonBill:                          "N",
		InvoiceStatus:                    "P",
		TransferredPeriod:                202603,
		RealizationAmountBillingCurrency: 1750.00,
		EmOrg:                            "ARCH",
		LastPSAExportDate:                "2026-02-01T00:00:00",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling PSALedgerEntry: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	// Spot-check key JSON field names.
	checks := map[string]any{
		"Batch":                   "TS-2026-01",
		"Period":                  float64(202601),
		"PKey":                    "PSA-001",
		"WBS1":                    "PRJ-200",
		"Account":                 "5010",
		"TransType":               "L",
		"Amount":                  1200.00,
		"BillExt":                 1800.00,
		"Employee":                "EMP042",
		"UnitQuantity":            8.0,
		"TransactionCurrencyCode": "USD",
		"AmountBillingCurrency":   1250.00,
		"DiaryNo":                 float64(42),
		"PONumber":                "PO-2026-100",
		"EmOrg":                   "ARCH",
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

	var decoded PSALedgerEntry
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to PSALedgerEntry: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestGLSummaryQueryValues verifies that GLSummaryQuery correctly encodes
// sumColumns and groupBy parameters alongside standard Query parameters.
func TestGLSummaryQueryValues(t *testing.T) {
	q := NewGLSummaryQuery()
	q.Limit(100)
	q.Filter("Period", "eq", "202601")
	q.SumColumns("Amount", "CBAmount", "CreditAmount")
	q.GroupBy("Account", "Org")

	vals := q.Values()

	// Verify sumColumns.
	if got := vals.Get("sumColumns"); got != "Amount,CBAmount,CreditAmount" {
		t.Errorf("sumColumns: got %q, want %q", got, "Amount,CBAmount,CreditAmount")
	}

	// Verify groupBy.
	if got := vals.Get("groupBy"); got != "Account,Org" {
		t.Errorf("groupBy: got %q, want %q", got, "Account,Org")
	}

	// Verify standard Query parameters still work.
	if got := vals.Get("limit"); got != "100" {
		t.Errorf("limit: got %q, want %q", got, "100")
	}

	// Verify filter was applied.
	if got := vals.Get("filterHash[0][name]"); got != "Period" {
		t.Errorf("filterHash[0][name]: got %q, want %q", got, "Period")
	}
	if got := vals.Get("filterHash[0][value]"); got != "202601" {
		t.Errorf("filterHash[0][value]: got %q, want %q", got, "202601")
	}
}

// TestGLSummaryQueryEmptyValues verifies that GLSummaryQuery omits sumColumns
// and groupBy when they are not set.
func TestGLSummaryQueryEmptyValues(t *testing.T) {
	q := NewGLSummaryQuery()
	vals := q.Values()

	if got := vals.Get("sumColumns"); got != "" {
		t.Errorf("sumColumns should be empty, got %q", got)
	}
	if got := vals.Get("groupBy"); got != "" {
		t.Errorf("groupBy should be empty, got %q", got)
	}
}
