package redis

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"navy/internal/port/domain"

	"github.com/redis/go-redis/v9"
)

// RedisRepository is a repository that uses Redis as a backend
type RedisRepository struct {
	rdb *redis.Client
}

// Upsert updates or inserts a Port
func (r RedisRepository) Upsert(ctx context.Context, p domain.Port) error {
	if resp := r.rdb.Set(ctx, p.ID, p, 0); resp != nil {
		return errors.Wrapf(resp.Err(), "failed to set port %s", p.ID)
	}

	return nil
}

// BatchUpsert updates or inserts a list of Ports
func (r RedisRepository) BatchUpsert(ctx context.Context, ps []domain.Port) error {
	p := r.rdb.Pipeline()
	for _, port := range ps {
		data, err := json.Marshal(port)
		if err != nil {
			return errors.Wrapf(err, "failed to marshal port %s", port.ID)
		}

		if resp := p.Set(ctx, port.ID, data, 0); resp.Err() != nil {
			return errors.Wrapf(resp.Err(), "failed to set port %s", port.ID)
		}
	}
	_, err := p.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to exec pipeline")
	}

	return nil
}

// Get gets a Port by its code
func (r RedisRepository) Get(ctx context.Context, id string) (domain.Port, error) {
	resp := r.rdb.Get(ctx, id)
	if resp.Err() != nil {
		return domain.Port{}, errors.Wrapf(resp.Err(), "failed to get port %s", id)
	}

	port := domain.Port{}
	if err := json.Unmarshal([]byte(resp.Val()), &port); err != nil {
		return domain.Port{}, errors.Wrapf(err, "failed to unmarshal port %s", id)
	}

	return port, nil
}

// Delete deletes a Port by its code
func (r RedisRepository) Delete(ctx context.Context, id string) error {
	resp := r.rdb.Del(ctx, id)
	if resp.Err() != nil {
		return errors.Wrapf(resp.Err(), "failed to delete port %s", id)
	}

	return nil
}

// NewRedisRepository creates a new RedisRepository
func NewRedisRepository(address string) *RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisRepository{
		rdb: rdb,
	}
}
