package cache

import (
	"context"

	"github.com/puzpuzpuz/xsync/v3"
)

type mapCache struct {
	xmap *xsync.MapOf[string, bool]
}

func NewMapCache() Cache {
	return &mapCache{
		xmap: xsync.NewMapOf[string, bool](),
	}
}

func (c *mapCache) Add(_ context.Context, key string) error {
	c.xmap.Store(key, true)
	return nil
}

func (c *mapCache) Remove(_ context.Context, key string) error {
	c.xmap.Delete(key)
	return nil
}

func (c *mapCache) Exists(_ context.Context, key ...string) (bool, error) {
	for _, v := range key {
		_, ok := c.xmap.Load(v)
		if ok {
			return true, nil
		}
	}
	return false, nil
}

func (c *mapCache) Size(_ context.Context) (int64, error) {
	return int64(c.xmap.Size()), nil
}
