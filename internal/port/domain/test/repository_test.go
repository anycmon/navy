package test

import (
	"context"
	"github.com/stretchr/testify/require"
	"navy/internal/port/domain"
	"navy/internal/port/domain/mem"
	"navy/internal/port/domain/redis"
	"os"
	"testing"
)

// As all repositories implement the same interface, we can test them all with the same tests.
func TestRepositories(t *testing.T) {
	t.Parallel()

	for name, repo := range repositories() {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testUpsertRepository(t, repo)
			testBatchUpsertRepository(t, repo)
		})
	}
}

func testBatchUpsertRepository(t *testing.T, repo domain.PortRepository) {
	t.Helper()
	ctx := context.Background()

	for _, tc := range []struct {
		name    string
		ports   []domain.Port
		wantErr error
	}{
		{
			name: "BatchUpsert few valid ports",
			ports: []domain.Port{
				{
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
				{
					ID:          "NYCPT2",
					Name:        "Test",
					City:        "New York2",
					Country:     "USA",
					Coordinates: []float64{1, 2},
					Province:    "NY",
					Timezone:    "America/New_York",
					Unlocs:      []string{"NYCPT"},
					Code:        "NYCPT",
				},

				{
					ID:          "NYCPT3",
					Name:        "Test",
					City:        "New York3",
					Country:     "USA",
					Coordinates: []float64{1, 2},
					Province:    "NY",
					Timezone:    "America/New_York",
					Unlocs:      []string{"NYCPT"},
					Code:        "NYCPT",
				},
			},
		},
		// TODO: add more test cases
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.BatchUpsert(ctx, tc.ports)
			require.Equal(t, tc.wantErr, err)

			for _, p := range tc.ports {
				assertPortInRepo(ctx, t, repo, p)
			}

			t.Cleanup(func() {
				for _, p := range tc.ports {
					err := repo.Delete(ctx, p.ID)
					require.NoError(t, err)
				}
			})
		})
	}

}

func testUpsertRepository(t *testing.T, repo domain.PortRepository) {
	t.Helper()
	ctx := context.Background()

	for _, tc := range []struct {
		name    string
		port    domain.Port
		wantErr error
	}{
		{
			name: "Upsert",
			port: domain.Port{
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
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Upsert(ctx, tc.port)
			require.Equal(t, tc.wantErr, err)

			assertPortInRepo(ctx, t, repo, tc.port)
		})

		t.Cleanup(func() {
			err := repo.Delete(ctx, tc.port.ID)
			require.NoError(t, err)
		})
	}
}

func assertPortInRepo(ctx context.Context, t *testing.T, repo domain.PortRepository, p domain.Port) {
	port, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	require.Equal(t, p, port)
}

func repositories() map[string]domain.PortRepository {
	return map[string]domain.PortRepository{
		"mem":   mem.NewMemRepository(),
		"redis": redis.NewRedisRepository(os.Getenv("TEST_REDIS_URL")),
	}
}
