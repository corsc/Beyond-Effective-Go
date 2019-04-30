package _6_csp_usability

import (
	"sync"
)

func newCacheUsingMutex() *CacheUsingMutex {
	// create a cache
	return &CacheUsingMutex{
		cache:      make(map[string]Person),
		cacheMutex: &sync.RWMutex{},
	}
}

type CacheUsingMutex struct {
	cache      map[string]Person
	cacheMutex *sync.RWMutex
}

func (c *CacheUsingMutex) Get(key string) Person {
	c.cacheMutex.RLock()
	defer c.cacheMutex.RUnlock()

	return c.cache[key]
}

func (c *CacheUsingMutex) Set(key string, value Person) {
	c.cacheMutex.Lock()
	defer c.cacheMutex.Unlock()

	c.cache[key] = value
}
