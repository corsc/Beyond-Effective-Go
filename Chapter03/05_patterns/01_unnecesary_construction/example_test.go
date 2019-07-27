package _1_unnecesary_construction

import (
	"errors"
	"fmt"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s ./Chapter03/05_patterns/01_unnecesary_construction/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/01_unnecesary_construction
// BenchmarkExample-8   	100000000	       191 ns/op
// BenchmarkFixed-8     	10000000000	         0.28 ns/op

func log(debug bool, msg string) {
	if !debug {
		return
	}

	println(msg)
}

func BenchmarkExample(b *testing.B) {
	debug := false
	err := errors.New("something failed")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		log(debug, fmt.Sprintf("error was: %s", err))
	}
}

func logFixed(debug bool, msg string, args ...interface{}) {
	if !debug {
		return
	}

	fmt.Printf(msg, args...)
}

func BenchmarkFixed(b *testing.B) {
	debug := false
	err := errors.New("something failed")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logFixed(debug, "error was: %s", err)
	}
}
