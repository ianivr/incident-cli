package incident

import (
	"encoding/json"
	"testing"
)

func TestIncidentJSONRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		build func(t *testing.T) *Incident
	}{
		{
			name: "new open incident",
			build: func(t *testing.T) *Incident {
				return mustNew(t, "New Incident", "", SeverityMedium)
			},
		},
		{
			name: "with description and id",
			build: func(t *testing.T) *Incident {
				inc := mustNew(t, "DB down", "primary replica unreachable", SeverityCritical)
				inc.ID = 42
				return inc
			},
		},
		{
			name: "closed incident",
			build: func(t *testing.T) *Incident {
				inc := mustNew(t, "Disk full", "", SeverityLow)
				if err := inc.Close(); err != nil {
					t.Fatalf("Close: %v", err)
				}
				return inc
			},
		},
		{
			name: "in progress incident",
			build: func(t *testing.T) *Incident {
				inc := mustNew(t, "Latency spike", "", SeverityHigh)
				inc.Status = StatusInProgress
				return inc
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertJSONRoundTrip(t, tt.build(t))
		})
	}
}

func mustNew(t *testing.T, title, description string, sev Severity) *Incident {
	t.Helper()
	inc, err := New(title, description, sev)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return inc
}

func assertJSONRoundTrip(t *testing.T, want *Incident) {
	t.Helper()

	data, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var got Incident
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if got.ID != want.ID {
		t.Errorf("ID = %v, want %v", got.ID, want.ID)
	}
	if got.Title != want.Title {
		t.Errorf("Title = %q, want %q", got.Title, want.Title)
	}
	if got.Description != want.Description {
		t.Errorf("Description = %q, want %q", got.Description, want.Description)
	}
	if got.Severity != want.Severity {
		t.Errorf("Severity = %q, want %q", got.Severity, want.Severity)
	}
	if got.Status != want.Status {
		t.Errorf("Status = %q, want %q", got.Status, want.Status)
	}
	if !got.CreatedAt.Equal(want.CreatedAt) {
		t.Errorf("CreatedAt = %v, want %v", got.CreatedAt, want.CreatedAt)
	}
	if !got.UpdatedAt.Equal(want.UpdatedAt) {
		t.Errorf("UpdatedAt = %v, want %v", got.UpdatedAt, want.UpdatedAt)
	}
	if !got.ClosedAt.Equal(want.ClosedAt) {
		t.Errorf("ClosedAt = %v, want %v", got.ClosedAt, want.ClosedAt)
	}
}
