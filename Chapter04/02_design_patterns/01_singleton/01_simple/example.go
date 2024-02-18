package _1_simple

import (
	"sync"
	"time"
)

var (
	instance   *Cache
	initConfig sync.Once
)

func GetCache() *Cache {
	initConfig.Do(func() {
		instance = &Cache{
			// we need some internals to prove they are the same
			items:     map[string]interface{}{},
			createdAt: time.Now(),
		}
	})

	return instance
}

type Cache struct {
	items     map[string]interface{}
	createdAt time.Time
}
