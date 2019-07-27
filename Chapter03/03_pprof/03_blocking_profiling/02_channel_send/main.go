package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

func main() {
	// start blocking profile
	runtime.SetBlockProfileRate(1)

	// create a semaphore channel to limit the parallelism
	semaphore := make(chan struct{}, 5)

	wg := &sync.WaitGroup{}

	// start lots of tasks
	for x := 0; x < 100; x++ {
		wg.Add(1)
		go performTask(wg, semaphore)
	}

	// wait until all tasks are complete
	wg.Wait()

	fmt.Println("All done!")

	// start the default mux to host the profiling
	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}

func performTask(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer func() {
		<-semaphore

		wg.Done()
	}()

	// wait for semaphore
	semaphore <- struct{}{}

	// do some task that takes a long time
	done := time.After(1 * time.Second)
	x := 0

	for {
		// waste time but use CPU
		x++

		select {
		case <-done:
			return

		default:
			x++
		}
	}
}
