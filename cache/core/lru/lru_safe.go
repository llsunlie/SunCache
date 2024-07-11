package lru

import (
	"SunCache/cache/core"
	"sync"
)

type SafeCache struct {
	mu       sync.Mutex
	lru      *Cache
	maxBytes int64
}

func NewSafeCache(maxBytes int64) (safeCache *SafeCache) {
	safeCache = &SafeCache{
		maxBytes: maxBytes,
		lru:      NewCache(maxBytes),
	}
	return
}

func (c *SafeCache) Add(key string, value core.Value) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = NewCache(c.maxBytes)
	}
	c.lru.add(key, value)
}

func (c *SafeCache) Get(key string) (value core.Value, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.get(key); ok {
		return v, ok
	}
	return
}

func (c *SafeCache) UseBytes() (count int64) {
	return c.lru.useBytes
}

func (c *SafeCache) MaxBytes() (count int64) {
	return c.maxBytes
}
