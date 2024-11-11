package cache

import (
	"context"
	"time"
)

type RedisClient interface {
	Set(ctx context.Context, key string, v interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, v interface{}) error
	Delete(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
}
