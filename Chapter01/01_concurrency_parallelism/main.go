package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	wg := &sync.WaitGroup{}

	// start our separate processes
	wg.Add(1)
	go output(wg, 0)

	wg.Add(1)
	go output(wg, 1)

	// wait all goroutines are done
	wg.Wait()
}

func output(wg *sync.WaitGroup, value int) {
	defer wg.Done()

	for x := 0; x < 15; x++ {
		print(value)

		// inform the schedule we can be interrupted
		runtime.Gosched()
	}
}
