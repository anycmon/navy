package service

import (
	"context"
	"navy/internal/common/nerrors"
	"navy/internal/port/domain"
	"navy/internal/port/domain/mem"
	"testing"
)

func TestService_Upsert(t *testing.T) {
	t.Parallel()

	type fields struct {
		repo domain.PortRepository
	}
	type args struct {
		p domain.Port
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Upsert a valid port",
			fields: fields{
				repo: mem.NewMemRepository(),
			},
			args: args{
				p: domain.Port{
					ID:          "NYCPT",
					Name:        "Test",
					City:        "New York",
					Country:     "USA",
					Coordinates: []float64{1, 2},
					Province:    "NY",
					Timezone:    "America/New_York",
					Unlocs:      []string{"NYCPT"},
					Code:        "NYCPT",
				},
			},
			wantErr: nil,
		},
		{
			name: "Upsert a port with empty Name",
			fields: fields{
				repo: mem.NewMemRepository(),
			},
			args: args{
				p: domain.Port{
					ID:          "NYCPT",
					Name:        "",
					City:        "New York",
					Country:     "USA",
					Coordinates: []float64{1, 2},
					Province:    "NY",
					Timezone:    "America/New_York",
					Unlocs:      []string{"NYCPT"},
					Code:        "NYCPT",
				},
			},
			wantErr: nerrors.ErrInvalidInput,
		},
		// TODO: Add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: tt.fields.repo,
			}
			if err := s.Upsert(context.Background(), tt.args.p); (err != nil) != (tt.wantErr != nil) {
				t.Errorf("Upsert() nerrors = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
