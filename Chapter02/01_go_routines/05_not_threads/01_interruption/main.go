package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	result := int64(0)

	processors := runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	for i := 0; i < processors; i++ {
		go func() {
			for {
				atomic.AddInt64(&result, 1)
			}
		}()
	}

	// give the goroutines a chance to run
	time.Sleep(1 * time.Second)

	fmt.Println("result: ", atomic.LoadInt64(&result))
}
