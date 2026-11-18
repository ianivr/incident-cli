package incident

import (
	"errors"
	"fmt"
)

// Status represents the status of an incident.
type Status int

// Status stages for incidents.
const (
	StatusUnknown Status = iota
	StatusOpen
	StatusInProgress
	StatusClosed
)

// String returns the lowercase name of the status.
func (s Status) String() string {
	switch s {
	case StatusOpen:
		return "open"
	case StatusInProgress:
		return "in_progress"
	case StatusClosed:
		return "closed"
	default:
		return fmt.Sprintf("Status(%d)", int(s))
	}
}

// ErrInvalidStatus is returned when an invalid status string is provided.
var ErrInvalidStatus = errors.New("invalid status")

// ParseStatus parses a string into a Status value. It returns an error if the string does not correspond to a valid status stage.
func ParseStatus(s string) (Status, error) {
	switch s {
	case "open":
		return StatusOpen, nil
	case "in_progress":
		return StatusInProgress, nil
	case "closed":
		return StatusClosed, nil
	default:
		return StatusUnknown, fmt.Errorf("%w: %q", ErrInvalidStatus, s)
	}
}

// IsValid checks if the Status value is valid (i.e., not StatusUnknown).
func (s Status) IsValid() bool {
	return s >= StatusOpen && s <= StatusClosed
}

// MarshalText implements the encoding.TextMarshaler interface for Status. It returns the lowercase name of the status as a byte slice.
func (s Status) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("%w: %q", ErrInvalidStatus, s)
	}
	return []byte(s.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Status. It parses a byte slice into a Status value.
func (s *Status) UnmarshalText(b []byte) error {
	parsedStatus, err := ParseStatus(string(b))
	if err != nil {
		return err
	}
	*s = parsedStatus
	return nil
}
