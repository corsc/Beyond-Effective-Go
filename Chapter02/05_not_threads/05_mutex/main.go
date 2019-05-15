package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	result := int64(0)
	resultMutex := &sync.Mutex{}

	processors := runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < processors; i++ {
		go func() {
			for {
				resultMutex.Lock()
				result++
				resultMutex.Unlock()
			}
		}()
	}

	// give the goroutines a chance to run
	time.Sleep(1 * time.Second)

	fmt.Println("result: ", result)
}
