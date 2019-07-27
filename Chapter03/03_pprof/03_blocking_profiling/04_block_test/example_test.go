package main

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"sync"
	"testing"
	"time"
)

// To run this
func TestBlocking(t *testing.T) {
	dataCh := make(chan int)
	go publishData(dataCh)

	// start multiple consumers
	wg := &sync.WaitGroup{}
	for x := 0; x < 100; x++ {
		wg.Add(1)

		go consumeData(wg, dataCh)
	}

	// wait until all tasks are complete
	wg.Wait()
}

func publishData(dataCh chan int) {
	// slowly publish data
	for x := 0; x < 1000; x++ {
		dataCh <- rand.Int()

		time.Sleep(10 * time.Millisecond)
	}

	close(dataCh)
}

func consumeData(wg *sync.WaitGroup, dataCh chan int) {
	defer wg.Done()

	for value := range dataCh {
		fmt.Printf("Value: %d\n", value)
	}
}
