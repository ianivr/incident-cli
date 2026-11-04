package incident

import (
	"errors"
	"testing"
)

func TestParseSeverity(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Severity
		wantErr error
	}{
		{name: "low is valid", input: "low", want: SeverityLow, wantErr: nil},
		{name: "medium is valid", input: "medium", want: SeverityMedium, wantErr: nil},
		{name: "high is valid", input: "high", want: SeverityHigh, wantErr: nil},
		{name: "critical is valid", input: "critical", want: SeverityCritical, wantErr: nil},
		{name: "HIGH is invalid", input: "HIGH", want: SeverityUnknown, wantErr: ErrInvalidSeverity},
		{name: "empty string is invalid", input: "", want: SeverityUnknown, wantErr: ErrInvalidSeverity},
		{name: "surrounding spaces are invalid", input: " low ", want: SeverityUnknown, wantErr: ErrInvalidSeverity},
		{name: "unknown word is invalid", input: "urgent", want: SeverityUnknown, wantErr: ErrInvalidSeverity},
		{name: "zero value name is invalid", input: "Severity(0)", want: SeverityUnknown, wantErr: ErrInvalidSeverity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSeverity(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseSeverity(%q) error = %v, want %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseSeverity(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
