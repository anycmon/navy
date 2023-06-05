package mem

import (
	"context"
	"navy/internal/common/nerrors"
	"navy/internal/port/domain"
)

// MemRepository is an in-memory repository for Ports
type MemRepository struct {
	ports map[string]domain.Port
}

// NewMemRepository creates a new MemRepository
func NewMemRepository() *MemRepository {
	return &MemRepository{
		ports: make(map[string]domain.Port),
	}
}

// Upsert updates or inserts a Port
func (r *MemRepository) Upsert(ctx context.Context, p domain.Port) error {
	r.ports[p.Code] = p

	return nil
}

// BatchUpsert updates or inserts a slice of Ports
func (r *MemRepository) BatchUpsert(ctx context.Context, ps []domain.Port) error {
	for _, p := range ps {
		if err := r.Upsert(ctx, p); err != nil {
			return err
		}
	}

	return nil
}

// Get returns a Port by Code
func (r *MemRepository) Get(ctx context.Context, id string) (domain.Port, error) {
	p, ok := r.ports[id]
	if !ok {
		return domain.Port{}, nerrors.ErrNotFound
	}

	return p, nil
}

// Delete deletes a Port by Code
func (r *MemRepository) Delete(ctx context.Context, id string) error {
	delete(r.ports, id)

	return nil
}
