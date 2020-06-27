package main

import (
	"runtime"
	"sync"
)

func main() {
	result := make(chan int, 100)

	wg := &sync.WaitGroup{}

	// start our separate processes
	wg.Add(1)
	go output(wg, 0, result)

	wg.Add(1)
	go output(wg, 1, result)

	// wait all goroutines are done
	wg.Wait()

	close(result)

	for value := range result {
		print(value)
	}
}

func output(wg *sync.WaitGroup, value int, result chan int) {
	defer wg.Done()

	for x := 0; x < 15; x++ {
		result <- value

		// inform the schedule we can be interrupted
		runtime.Gosched()
	}
}
