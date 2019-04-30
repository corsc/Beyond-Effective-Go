package main

import (
	"runtime"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	semaphore := make(chan struct{}, runtime.NumCPU())

	for x := 0; x < 100000; x++ {
		wg.Add(1)
		go doWork(wg, semaphore)
	}

	wg.Wait()
}

// intentionally do something CPU intensive
func doWork(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	// reserve
	semaphore <- struct{}{}

	sum := 0

	for x := 0; x < 1000000; x++ {
		sum += x

		if x%100 == 0 {
			// yield the scheduler to switch to something else
			runtime.Gosched()
		}
	}

	// release
	<-semaphore
}
