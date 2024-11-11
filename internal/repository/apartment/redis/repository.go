package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	cachecl "github.com/escoutdoor/bookit/internal/client/cache"
	"github.com/escoutdoor/bookit/internal/repository"
	aparterr "github.com/escoutdoor/bookit/internal/repository/apartment"
	"github.com/escoutdoor/bookit/internal/repository/apartment/redis/converter"
	repomodel "github.com/escoutdoor/bookit/internal/repository/apartment/redis/model"
	"github.com/redis/go-redis/v9"

	"github.com/escoutdoor/bookit/internal/model"
	"github.com/google/uuid"
)

type KeyGenerator interface {
	Generate(id uuid.UUID) string
}

type DefaultKeyGenerator struct{}

func (kg *DefaultKeyGenerator) Generate(id uuid.UUID) string {
	return fmt.Sprintf("apartment:%s", id)
}

type cache struct {
	client       cachecl.RedisClient
	keyGenerator KeyGenerator
	ttl          time.Duration
}

var _ repository.ApartmentCache = (*cache)(nil)

func NewApartmentCache(
	cl cachecl.RedisClient,
	keyGenerator KeyGenerator,
	ttl time.Duration,
) *cache {
	// for testing
	// fmt.Printf("key generated: %s\n", keyGenerator.Generate(uuid.New()))
	return &cache{
		client:       cl,
		keyGenerator: keyGenerator,
		ttl:          ttl,
	}
}

func (r *cache) GetByID(ctx context.Context, id uuid.UUID) (*model.Apartment, error) {
	const op = "ApartmentCache.GetByID"
	key := r.keyGenerator.Generate(id)

	var apartment repomodel.Apartment
	err := r.client.Get(ctx, key, &apartment)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, aparterr.ErrNotFound
		}

		return nil, fmt.Errorf("%s: get: %s", op, err)
	}

	return converter.ToApartmentFromRepository(&apartment), nil
}

func (r *cache) Set(ctx context.Context, in *model.Apartment) error {
	const op = "ApartmentCache.Set"
	apartment := converter.ToRepositoryFromApartment(in)
	key := r.keyGenerator.Generate(apartment.ID)

	err := r.client.Set(ctx, key, apartment, r.ttl)
	if err != nil {
		return fmt.Errorf("%s: set cache: %s", op, err)
	}

	return nil
}

func (r *cache) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "ApartmentCache.Delete"
	key := r.keyGenerator.Generate(id)

	err := r.client.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}

func (r *cache) Expire(ctx context.Context, id uuid.UUID) error {
	const op = "ApartmentCache.Expire"
	key := r.keyGenerator.Generate(id)

	err := r.client.Expire(ctx, key, r.ttl)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}
