package _1_mutexes

import (
	"sync"
)

type ConcurrentMap struct {
	Data map[string]string
	sync.Mutex
}

func Usage() {
	cmap := &ConcurrentMap{}

	cmap.Lock()
	cmap.Data["a"] = "b"
	cmap.Unlock()
}
