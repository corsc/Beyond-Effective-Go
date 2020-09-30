package _2_channel_based

import (
	"fmt"
	"time"
)

func Example() {
	futureA := doWorkFuture(123)
	futureB := doWorkFuture(456)

	// other code

	valueA := <-futureA
	valueB := <-futureB

	// Output: Receive future value: 1123
	// Receive future value: 1456
	fmt.Printf("Receive future value: %d\n", valueA)
	fmt.Printf("Receive future value: %d\n", valueB)
}

func doWorkFuture(in int64) <-chan int64 {
	result := make(chan int64, 1)

	// start goroutine to calculate/fetch the value asynchronously
	go doTimeConsumingWork(result, in)

	return result
}

func doTimeConsumingWork(result chan<- int64, in int64) {
	// do something that takes a while
	<-time.After(1 * time.Second)

	// assign the result of the calculation/fetch
	result <- 1000 + in
}
