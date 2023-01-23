package main

import "sync"

type CVEStatsCache struct {
	cache map[string]CVEStats
	lock  sync.RWMutex
}

func NewStatsCache() CVEStatsCache {
	return CVEStatsCache{cache: make(map[string]CVEStats), lock: sync.RWMutex{}}
}

func (c *CVEStatsCache) Set(key string, value CVEStats) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[key] = value
}

func (c *CVEStatsCache) Get(key string) (CVEStats, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	v, ok := c.cache[key]
	return v, ok
}
