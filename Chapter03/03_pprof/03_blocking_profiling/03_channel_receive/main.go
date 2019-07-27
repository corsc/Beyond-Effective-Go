package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func main() {
	// start blocking profile
	runtime.SetBlockProfileRate(1)

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

	fmt.Println("All done!")

	// start the default mux to host the profiling
	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
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
