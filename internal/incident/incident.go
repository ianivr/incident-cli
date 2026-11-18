package incident

import (
	"errors"
	"fmt"
	"time"
)

// Incident represents an incident with its details.
type Incident struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Severity    Severity  `json:"severity"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ClosedAt    time.Time `json:"closed_at,omitzero"`
}

// ErrInvalidTitle is returned when an incident is created with an empty title.
var ErrInvalidTitle = errors.New("title cannot be empty")

// ErrAlreadyClosed is returned when an attempt is made to close an already closed incident.
var ErrAlreadyClosed = errors.New("incident is already closed")

// New creates a new Incident instance with the provided title, description, and severity. It validates its input and returns an error for invalid incidents.
func New(title, description string, s Severity) (*Incident, error) {
	if title == "" {
		return nil, fmt.Errorf("%w: %q", ErrInvalidTitle, title)
	}
	if !s.IsValid() {
		return nil, fmt.Errorf("%w: %q", ErrInvalidSeverity, s)
	}

	timeNow := time.Now()

	return &Incident{
		Title:       title,
		Description: description,
		Severity:    s,
		Status:      StatusOpen,
		CreatedAt:   timeNow,
		UpdatedAt:   timeNow,
	}, nil
}

// Close marks the incident as closed and records the time of closure. It returns an error if the incident is already closed.
func (inc *Incident) Close() error {
	if inc.Status == StatusClosed {
		return ErrAlreadyClosed
	}

	timeNow := time.Now()

	inc.Status = StatusClosed
	inc.ClosedAt = timeNow
	inc.UpdatedAt = timeNow

	return nil
}
