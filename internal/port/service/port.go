package service

import (
	"context"
	"github.com/pkg/errors"
	"navy/internal/common/validate"
	"navy/internal/port/domain"
	"navy/internal/port/domain/mem"
)

type Service struct {
	repo domain.PortRepository
}

// NewService creates a new Service
func NewService(cfgs ...ServiceConfiguration) (*Service, error) {
	os := &Service{}
	// Apply all Configurations passed in

	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// Upsert updates or inserts a Port
func (s *Service) Upsert(ctx context.Context, p domain.Port) error {
	err := validate.Validate(p)
	if err != nil {
		return errors.Wrap(err, "invalid port")
	}

	return s.repo.Upsert(ctx, p)
}

func (s *Service) BatchUpsert(ctx context.Context, ps []domain.Port) error {
	for _, p := range ps {
		err := validate.Validate(p)
		if err != nil {
			return errors.Wrapf(err, "invalid port: %v", p)
		}
	}

	return s.repo.BatchUpsert(ctx, ps)
}

// ServiceConfiguration is a function that configures a Service
type ServiceConfiguration func(svc *Service) error

// WithRepository is a ServiceConfiguration that configures the Service to use a given repository
func WithRepository(repo domain.PortRepository) ServiceConfiguration {
	return func(svc *Service) error {
		svc.repo = repo
		return nil
	}
}

// WithMemRepository is a ServiceConfiguration that configures the Service to use an in-memory repository
func WithMemRepository() ServiceConfiguration {
	return func(svc *Service) error {
		svc.repo = mem.NewMemRepository()
		return nil
	}
}
