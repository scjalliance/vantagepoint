package vantagepoint

import (
	"encoding/json"
	"testing"
)

// TestFirmJSONRoundtrip verifies that the Firm struct marshals and unmarshals
// with the correct JSON field names matching the Vantagepoint API.
// Key field names verified: ClientID (not FirmCode), WebSite (not Website),
// PrimaryEmail (not Email), HasHierarchy, Client field.
func TestFirmJSONRoundtrip(t *testing.T) {
	original := Firm{
		ClientID:      "abc-123-def",
		Client:        "ACME",
		Name:          "Acme Corporation",
		Status:        "A",
		WebSite:       "https://acme.example.com",
		PrimaryEmail:  "info@acme.example.com",
		PrimaryPhone:  "555-0100",
		PrimaryCity:   "Seattle",
		PrimaryState:  "WA",
		PrimaryZip:    "98101",
		Employees:     250,
		AnnualRevenue: 5000000.50,
		HasHierarchy:  1,
		ClientInd:     "Y",
		VendorInd:     "N",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling Firm: %v", err)
	}

	// Verify the JSON uses the correct VP field names.
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshaling to map: %v", err)
	}

	expectedKeys := map[string]any{
		"ClientID":      "abc-123-def",
		"Client":        "ACME",
		"Name":          "Acme Corporation",
		"Status":        "A",
		"WebSite":       "https://acme.example.com",
		"PrimaryEmail":  "info@acme.example.com",
		"PrimaryPhone":  "555-0100",
		"PrimaryCity":   "Seattle",
		"PrimaryState":  "WA",
		"PrimaryZip":    "98101",
		"Employees":     float64(250),
		"AnnualRevenue": 5000000.50,
		"HasHierarchy":  float64(1),
		"ClientInd":     "Y",
		"VendorInd":     "N",
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
	badKeys := []string{"FirmCode", "Website", "Email"}
	for _, key := range badKeys {
		if _, ok := raw[key]; ok {
			t.Errorf("unexpected old JSON key %q found in output", key)
		}
	}

	// Verify roundtrip: unmarshal back into a Firm and compare.
	var decoded Firm
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling back to Firm: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}

// TestFirmJSONFromAPI verifies that a JSON payload using API field names
// correctly deserializes into the Firm struct.
func TestFirmJSONFromAPI(t *testing.T) {
	apiJSON := `{
		"ClientID": "guid-456-xyz",
		"Client": "GLOBEX",
		"Name": "Globex Corporation",
		"Status": "A",
		"WebSite": "https://globex.example.com",
		"PrimaryEmail": "contact@globex.example.com",
		"PrimaryPhone": "555-0200",
		"PrimaryAddress1": "100 Main St",
		"PrimaryCity": "Portland",
		"PrimaryState": "OR",
		"PrimaryZip": "97201",
		"PrimaryCountry": "US",
		"Employees": 500,
		"AnnualRevenue": 10000000.00,
		"HasHierarchy": 1,
		"ClientInd": "Y",
		"VendorInd": "Y",
		"Org": "WEST",
		"Owner": "EMP001",
		"Category": "ARCH",
		"CreateDate": "2024-01-15T00:00:00",
		"ModDate": "2025-06-01T00:00:00"
	}`

	var f Firm
	if err := json.Unmarshal([]byte(apiJSON), &f); err != nil {
		t.Fatalf("unmarshaling API JSON: %v", err)
	}

	if f.ClientID != "guid-456-xyz" {
		t.Errorf("ClientID: got %q, want %q", f.ClientID, "guid-456-xyz")
	}
	if f.Client != "GLOBEX" {
		t.Errorf("Client: got %q, want %q", f.Client, "GLOBEX")
	}
	if f.WebSite != "https://globex.example.com" {
		t.Errorf("WebSite: got %q, want %q", f.WebSite, "https://globex.example.com")
	}
	if f.PrimaryEmail != "contact@globex.example.com" {
		t.Errorf("PrimaryEmail: got %q, want %q", f.PrimaryEmail, "contact@globex.example.com")
	}
	if f.HasHierarchy != 1 {
		t.Errorf("HasHierarchy: got %d, want %d", f.HasHierarchy, 1)
	}
	if f.Employees != 500 {
		t.Errorf("Employees: got %d, want %d", f.Employees, 500)
	}
	if f.AnnualRevenue != 10000000.00 {
		t.Errorf("AnnualRevenue: got %v, want %v", f.AnnualRevenue, 10000000.00)
	}
	if f.Org != "WEST" {
		t.Errorf("Org: got %q, want %q", f.Org, "WEST")
	}
}

// TestFirmOmitsEmptyFields verifies that zero-value fields are omitted
// from JSON output due to the omitempty tags.
func TestFirmOmitsEmptyFields(t *testing.T) {
	f := Firm{
		ClientID: "abc-123",
		Name:     "Test Corp",
	}

	data, err := json.Marshal(f)
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
	if raw["ClientID"] != "abc-123" {
		t.Errorf("ClientID: got %v, want %q", raw["ClientID"], "abc-123")
	}
	if raw["Name"] != "Test Corp" {
		t.Errorf("Name: got %v, want %q", raw["Name"], "Test Corp")
	}
}

// TestFirmHasHierarchyIsInt verifies that HasHierarchy is an int field
// (0 or 1) rather than a boolean.
func TestFirmHasHierarchyIsInt(t *testing.T) {
	for _, val := range []int{0, 1} {
		f := Firm{ClientID: "test-id", HasHierarchy: val}
		data, err := json.Marshal(f)
		if err != nil {
			t.Fatalf("marshaling with HasHierarchy=%d: %v", val, err)
		}

		var raw map[string]any
		if err := json.Unmarshal(data, &raw); err != nil {
			t.Fatalf("unmarshaling to map: %v", err)
		}

		if val == 0 {
			// omitempty should omit zero-value int
			if _, ok := raw["HasHierarchy"]; ok {
				t.Errorf("HasHierarchy=0 should be omitted by omitempty")
			}
			continue
		}

		got, ok := raw["HasHierarchy"].(float64)
		if !ok {
			t.Errorf("HasHierarchy should be a number, got %T", raw["HasHierarchy"])
		}
		if int(got) != val {
			t.Errorf("HasHierarchy: got %v, want %d", got, val)
		}
	}
}

func TestFirmProjectJSONRoundtrip(t *testing.T) {
	original := FirmProject{
		ClientID:  "CLIENTID123",
		WBS1:      "0300000.00",
		Role:      "sysOwner",
		RoleName:  "Owner",
		PRName:    "Anderson Clinic Expansion",
		PRStatus:  "A",
		ClientInd: "Y",
		VendorInd: "N",
	}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshaling FirmProject: %v", err)
	}
	var decoded FirmProject
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshaling FirmProject: %v", err)
	}
	if decoded != original {
		t.Errorf("roundtrip mismatch:\n  got:  %+v\n  want: %+v", decoded, original)
	}
}
