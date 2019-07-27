package _4_defer

import (
	"sync"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s -benchmem ./Chapter03/05_patterns/14_defer/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/14_defer
// BenchmarkExample-8   	300000000	        46.4 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFixed-8     	1000000000	        16.0 ns/op	       0 B/op	       0 allocs/op

func BenchmarkExample(b *testing.B) {
	mutex := &sync.Mutex{}

	for i := 0; i < b.N; i++ {
		doWorkWithDefer(mutex)
	}
}

func doWorkWithDefer(mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()

	// do other things
}

func BenchmarkFixed(b *testing.B) {
	mutex := &sync.Mutex{}

	for i := 0; i < b.N; i++ {
		doWorkWithoutDefer(mutex)
	}
}

func doWorkWithoutDefer(mutex *sync.Mutex) {
	mutex.Lock()

	// do other things

	mutex.Unlock()
}
