package _1_waitgroup

import (
	"sync"
)

func Example() bool {
	wg := &sync.WaitGroup{}

	// start some goroutines
	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// do something
		}()
	}

	// wait for all goroutines to complete
	wg.Wait()

	// all goroutines have finished
	return true
}
