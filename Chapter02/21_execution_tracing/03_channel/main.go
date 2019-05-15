package main

import (
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	// create file to hold the trace data
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// start/stop the tracer
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	dataCh := make(chan struct{})

	go send(wg, dataCh)
	go receive(wg, dataCh)

	wg.Wait()
}

func send(wg *sync.WaitGroup, dataCh chan<- struct{}) {
	defer wg.Done()

	<-time.After(10 * time.Microsecond)

	dataCh <- struct{}{}
}

func receive(wg *sync.WaitGroup, dataCh <-chan struct{}) {
	defer wg.Done()

	<-dataCh
}
