package _1_simple

import (
	"sync"
)

var (
	instance   *Cache
	initConfig sync.Once
)

func GetCache() *Cache {
	initConfig.Do(func() {
		instance = &Cache{}
	})

	return instance
}

type Cache struct {
	// not implemented
}
