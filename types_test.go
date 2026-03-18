package vantagepoint

import (
	"encoding/json"
	"testing"
	"time"
)

// TestDateOnlyUnmarshalJSON verifies that DateOnly correctly unmarshals from
// the various date formats returned by the Vantagepoint API.
func TestDateOnlyUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:  "date only",
			input: `"2023-06-15"`,
			want:  time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "datetime without millis",
			input: `"2023-06-15T00:00:00"`,
			want:  time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "VP format with millis",
			input: `"2017-12-16T00:00:00.000"`,
			want:  time.Date(2017, 12, 16, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "datetime with Z suffix",
			input: `"2023-06-15T00:00:00Z"`,
			want:  time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "RFC3339 with timezone offset",
			input: `"2023-06-15T00:00:00-07:00"`,
			want:  time.Date(2023, 6, 15, 0, 0, 0, 0, time.FixedZone("", -7*3600)),
		},
		{
			name:  "null value",
			input: `null`,
			want:  time.Time{},
		},
		{
			name:  "empty string",
			input: `""`,
			want:  time.Time{},
		},
		{
			name:    "invalid format",
			input:   `"not-a-date"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d DateOnly
			err := json.Unmarshal([]byte(tt.input), &d)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !d.Time.Equal(tt.want) {
				t.Errorf("got %v, want %v", d.Time, tt.want)
			}
		})
	}
}

// TestDateOnlyMarshalJSON verifies that DateOnly correctly marshals to the
// canonical "2006-01-02" format, and zero values become null.
func TestDateOnlyMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		date DateOnly
		want string
	}{
		{
			name: "normal date",
			date: DateOnly{Time: time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)},
			want: `"2023-06-15"`,
		},
		{
			name: "zero value",
			date: DateOnly{},
			want: `null`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.date)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if string(got) != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
