package domain

import "context"

// PortRepository is a repository for Ports
type PortRepository interface {
	// Upsert updates or inserts a Port
	Upsert(ctx context.Context, p Port) error

	// BatchUpsert updates or inserts a slice of Ports
	BatchUpsert(ctx context.Context, ps []Port) error

	// Get returns a Port by Code
	Get(ctx context.Context, code string) (Port, error)

	// Delete deletes a Port by Code
	Delete(ctx context.Context, code string) error
}
