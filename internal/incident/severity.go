package incident

import (
	"errors"
	"fmt"
)

// Severity represents the severity level of an incident.
type Severity int

// Severity levels for incidents.
const (
	SeverityUnknown Severity = iota
	SeverityLow
	SeverityMedium
	SeverityHigh
	SeverityCritical
)

// String returns the lowercase name of the severity.
func (s Severity) String() string {
	switch s {
	case SeverityLow:
		return "low"
	case SeverityMedium:
		return "medium"
	case SeverityHigh:
		return "high"
	case SeverityCritical:
		return "critical"
	default:
		return fmt.Sprintf("Severity(%d)", int(s))
	}
}

// ErrInvalidSeverity is returned when an invalid severity string is provided.
var ErrInvalidSeverity = errors.New("invalid severity")

// ParseSeverity parses a string into a Severity value. It returns an error if the string does not correspond to a valid severity level.
func ParseSeverity(s string) (Severity, error) {
	switch s {
	case "low":
		return SeverityLow, nil
	case "medium":
		return SeverityMedium, nil
	case "high":
		return SeverityHigh, nil
	case "critical":
		return SeverityCritical, nil
	default:
		return SeverityUnknown, fmt.Errorf("%w: %q", ErrInvalidSeverity, s)
	}
}

// IsValid checks if the Severity value is valid (i.e., not SeverityUnknown).
func (s Severity) IsValid() bool {
	return s >= SeverityLow && s <= SeverityCritical
}

// MarshalText implements the encoding.TextMarshaler interface for Severity. It returns the lowercase name of the severity as a byte slice.
func (s Severity) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("%w: %q", ErrInvalidSeverity, s)
	}
	return []byte(s.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Severity. It parses a byte slice into a Severity value, returning an error if the input is invalid.
func (s *Severity) UnmarshalText(b []byte) error {
	sev, err := ParseSeverity(string(b))
	if err != nil {
		return err
	}
	*s = sev
	return nil
}
