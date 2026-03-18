package vantagepoint

import "time"

// DateOnly is a date without time information, formatted as YYYY-MM-DD for JSON.
type DateOnly struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface for DateOnly.
// Zero values are marshaled as JSON null.
func (d DateOnly) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.Format("2006-01-02") + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for DateOnly.
// It accepts multiple date formats that Vantagepoint may return, including
// "2006-01-02", "2006-01-02T15:04:05", "2006-01-02T15:04:05.000",
// "2006-01-02T15:04:05Z", and RFC 3339.
func (d *DateOnly) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `""` {
		d.Time = time.Time{}
		return nil
	}
	// Remove quotes
	s = s[1 : len(s)-1]
	// Try multiple date formats VP might return
	for _, layout := range []string{
		"2006-01-02",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05Z",
	} {
		t, err := time.Parse(layout, s)
		if err == nil {
			d.Time = t
			return nil
		}
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
