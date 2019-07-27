package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

var (
	counter int
	mutex   = &sync.Mutex{}
)

func main() {
	// start mutex profile
	runtime.SetMutexProfileFraction(5)

	wg := &sync.WaitGroup{}

	// lots of tasks
	for x := 0; x < 1000; x++ {
		wg.Add(1)

		go performTask(wg)
	}

	// wait until all tasks are complete
	wg.Wait()

	fmt.Println("All done!")

	// start the default mux to host the profiling
	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}

func performTask(wg *sync.WaitGroup) {
	defer wg.Done()

	for x := 0; x < 100000; x++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}
