package _1_function_based

import (
	"fmt"
	"sync"
	"time"
)

func Example() {
	futureA := doWorkFuture(123)
	futureB := doWorkFuture(456)

	// other code

	valueA := futureA()
	valueB := futureB()

	// Output: Receive future value: 1123
	// Receive future value: 1456
	fmt.Printf("Receive future value: %d\n", valueA)
	fmt.Printf("Receive future value: %d\n", valueB)
}

func doWorkFuture(in int64) func() int64 {
	var result int64
	mutex := &sync.Mutex{}
	mutex.Lock()

	// start goroutine to calculate/fetch the value asynchronously
	go doTimeConsumingWork(mutex, &result, in)

	return func() int64 {
		// this function cannot return until the value has been calculated
		mutex.Lock()
		defer mutex.Unlock()

		return result
	}
}

func doTimeConsumingWork(mutex *sync.Mutex, result *int64, in int64) {
	defer mutex.Unlock()

	// do something that takes a while
	<-time.After(1 * time.Second)

	// assign the result of the calculation/fetch
	*result = int64(1000 + in)
}
