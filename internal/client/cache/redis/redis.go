package redis

import (
	"context"
	"fmt"

	"github.com/escoutdoor/bookit/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewConn(ctx context.Context, cfg *config.RedisConfig) (*redis.Client, error) {
	const op = "redis.NewConn"
	url := generateURL(cfg)

	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("%s: parse url: %s", op, err)
	}

	cl := redis.NewClient(opts)
	_, err = cl.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%s: ping redis: %s", op, err)
	}

	return cl, nil
}

func generateURL(cfg *config.RedisConfig) string {
	return fmt.Sprintf("redis://%s:%d/%d", cfg.Host, cfg.Port, *cfg.DB)
}
