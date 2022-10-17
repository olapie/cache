package mem

import (
	"context"
	"time"

	"code.olapie.com/cache"
	"code.olapie.com/errors"
	gocache "github.com/patrickmn/go-cache"
)

type SimpleCache struct {
	c *gocache.Cache
}

var _ cache.SimpleCache = (*SimpleCache)(nil)

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		c: gocache.New(time.Hour, 20*time.Minute),
	}
}

func (c *SimpleCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	c.c.Set(key, value, ttl)
	return nil
}

func (c *SimpleCache) Get(ctx context.Context, key string) ([]byte, error) {
	v, ok := c.c.Get(key)
	if !ok {
		return nil, errors.NotExist
	}
	if data, ok := v.([]byte); ok {
		return data, nil
	}
	return nil, errors.New("invalid data type")
}

func (c *SimpleCache) Delete(ctx context.Context, key string) error {
	c.c.Delete(key)
	return nil
}
