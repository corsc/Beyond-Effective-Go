package _3_cross_streams

import (
	"sync"
)

func Example() {
	outputCh := make(chan Data)
	mutex := &sync.Mutex{}

	// example starts here
	mutex.Lock()

	outputCh <- Data{}

	mutex.Unlock()
}

type Data struct {
	// some fields
}
