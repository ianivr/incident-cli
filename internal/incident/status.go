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
