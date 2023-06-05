package test

import (
	"errors"
	"navy/internal/common/nerrors"
	"navy/internal/port/domain"
	"testing"
)

func TestPort_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		given   domain.Port
		wantErr error
	}{
		{
			name: "Validate a valid port",
			given: domain.Port{
				ID:      "AEAJM",
				Name:    "Ajman",
				City:    "Ajman",
				Country: "United Arab Emirates",
				Coordinates: []float64{
					55.5136433,
					25.4052165,
				},
				Province: "Ajman",
				Timezone: "Asia/Dubai",
				Unlocs:   []string{"AEAJM"},
				Code:     "52000",
			},
			wantErr: nil,
		},
		{
			name: "Validate a port with empty Name",
			given: domain.Port{
				Name:    "", // empty
				City:    "Ajman",
				Country: "United Arab Emirates",
				Coordinates: []float64{
					55.5136433,
					25.4052165,
				},
				Province: "Ajman",
				Timezone: "Asia/Dubai",
				Unlocs:   []string{"AEAJM"},
				Code:     "52000",
			},
			wantErr: nerrors.ErrInvalidInput,
		},
		// TODO: add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := tt.given.Validate(); !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
