package api

import (
	"sync"

	"vollkorntomate.de/cvedash-api/internal/data"
)

type CVEStatsCache struct {
	cache map[string]data.CVEStats
	lock  sync.RWMutex
}

func NewStatsCache() CVEStatsCache {
	return CVEStatsCache{cache: make(map[string]data.CVEStats), lock: sync.RWMutex{}}
}

func (c *CVEStatsCache) Set(key string, value data.CVEStats) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[key] = value
}

func (c *CVEStatsCache) Get(key string) (data.CVEStats, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	v, ok := c.cache[key]
	return v, ok
}

func (c *CVEStatsCache) GetAll() map[string]data.CVEStats {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.cache
}
