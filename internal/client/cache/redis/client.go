package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/escoutdoor/bookit/internal/client/cache"
	"github.com/escoutdoor/bookit/internal/config"
	"github.com/redis/go-redis/v9"
)

type client struct {
	redisClient *redis.Client
	cfg         *config.RedisConfig
}

var _ cache.RedisClient = (*client)(nil)

func NewClient(ctx context.Context, cfg *config.RedisConfig) (*client, error) {
	const op = "redis.NewClient"

	redisConn, err := NewConn(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: new redis connection: %s", op, err)
	}

	return &client{
		redisClient: redisConn,
		cfg:         cfg,
	}, nil
}

func (cl *client) Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error {
	_, err := cl.redisClient.Set(
		ctx,
		key,
		v,
		expiration,
	).Result()
	if err != nil {
		return err
	}

	return nil
}

func (cl *client) Get(ctx context.Context, key string, v interface{}) error {
	err := cl.redisClient.Get(ctx, key).Scan(v)
	if err != nil {
		return err
	}

	return nil
}

func (cl *client) Delete(ctx context.Context, key string) error {
	err := cl.redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cl *client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := cl.redisClient.Expire(ctx, key, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}
