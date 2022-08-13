package main

import (
	"os"
	"runtime/trace"
)

func main() {
	// create a file to hold the trace data
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

	// rest of the program goes here
}
