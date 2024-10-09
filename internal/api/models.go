package api

import (
	"fmt"
	"time"
)

// Date is a wrapper for time.Time, implementing json.Unmarshaler.
//
// The date needs to be in format "YYYY-MM-DD" to be unmarshalled correctly.
type Date struct {
	time.Time
}

// UnmarshalJSON parsed the provided bytes into a time.Time instance.
//
// The date needs to be in format "YYYY-MM-DD" to be unmarshalled correctly.
func (d *Date) UnmarshalJSON(b []byte) error {
	data := b[len(`"`) : len(b)-len(`"`)]
	t, err := time.Parse("2006-01-02", string(data))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// String returns the string representation.
func (d *Date) String() string {
	return d.Format("2006-01-02")
}

// Status represents the status of the game server. This also includes a few global elements, such as announcements, server reset dates and leaderboards.
//
// Note: this struct is defined manually since oapi-codegen doesn't generate it.
type Status struct {
	Status     string       `json:"status"`
	Version    string       `json:"version"`
	LastReset  Date         `json:"resetDate"`
	Resets     ServerResets `json:"serverResets"`
	Statistics struct {
		Agents    int64 `json:"agents"`
		Ships     int64 `json:"ships"`
		Waypoints int64 `json:"waypoints"`
		Systems   int64 `json:"systems"`
	} `json:"stats"`
	Announcements []struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"announcements"`
}

// ServerResets represents the server's reset frequency and the next planned reset.
type ServerResets struct {
	Next      time.Time `json:"next"`
	Frequency string    `json:"frequency"`
}

// String implements the Stringer interface.
func (s *Status) String() string {
	return fmt.Sprintf("Status{'%s'}", s.Status)
}
