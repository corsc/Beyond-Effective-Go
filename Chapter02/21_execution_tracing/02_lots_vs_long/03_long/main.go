package main

import (
	"bytes"
	"math/rand"
	"os"
	"runtime/trace"
	"strconv"
	"sync"
	"time"
)

const (
	maxConcurrent = 5
	totalWork     = 1000000
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

	rand.Seed(time.Now().UnixNano())

	wg := &sync.WaitGroup{}

	for x := 0; x < maxConcurrent; x++ {
		wg.Add(1)
		go worker(wg, totalWork/maxConcurrent)
	}

	wg.Wait()
}

func worker(wg *sync.WaitGroup, total int) {
	defer wg.Done()

	for x := 0; x < total; x++ {
		first := rand.Int63()
		second := rand.Int63()

		buffer := &bytes.Buffer{}

		_, _ = buffer.WriteString(strconv.FormatInt(first, 10))
		_, _ = buffer.WriteString("+")
		_, _ = buffer.WriteString(strconv.FormatInt(second, 10))
		_, _ = buffer.WriteString("=")
		_, _ = buffer.WriteString(strconv.FormatInt(first+second, 10))
	}
}
