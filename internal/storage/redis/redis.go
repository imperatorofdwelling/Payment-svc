package redis

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient(c config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       0,
		Protocol: c.Protocol,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return rdb, nil
}
