package redis

import (
	"context"
	"time"

	"code.olapie.com/cache"
	"code.olapie.com/errors"
	"github.com/go-redis/redis"
)

type SimpleCache struct {
	c *redis.Client
}

var _ cache.SimpleCache = (*SimpleCache)(nil)

func NewSimpleCache(c *redis.Client) *SimpleCache {
	return &SimpleCache{
		c: c,
	}
}

func (c *SimpleCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.c.WithContext(ctx).Set(key, value, ttl).Err()
}

func (c *SimpleCache) Get(ctx context.Context, key string) ([]byte, error) {
	data, err := c.c.WithContext(ctx).Get(key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, errors.NotExist
	}
	return data, nil
}

func (c *SimpleCache) Delete(ctx context.Context, key string) error {
	return c.c.WithContext(ctx).Del(key).Err()
}
