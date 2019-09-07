package _2_strict

import (
	"sync"
)

var (
	instance   Cache
	initConfig sync.Once
)

func GetCache() Cache {
	initConfig.Do(func() {
		instance = &myCache{}
	})

	return instance
}

type Cache interface {
	Get(key string) string
	Put(key, value string)
}

type myCache struct {
	// not implemented
}

func (m myCache) Get(key string) string {
	// implementation removed
	return ""
}

func (m myCache) Put(key, value string) {
	// implementation removed
}
