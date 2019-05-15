package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	go func() {
		<-time.After(1 * time.Second)
	}()

	profile := pprof.Lookup("goroutine")
	profile.WriteTo(os.Stdout, 1)
}
