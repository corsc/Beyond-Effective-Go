package _1_when_to_optimize

import (
	"bytes"
	"testing"
)

// $ go test -bench=. -benchtime=15s ./Chapter03/01_when_to_optimize/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/01_when_to_optimize
// BenchmarkCleanExample-8   	30000000	       767 ns/op
// BenchmarkFastExample-8    	100000000	       249 ns/op

var result string

func BenchmarkCleanExample(b *testing.B) {
	people := []Person{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Paul"},
		{ID: 3, Name: "George"},
		{ID: 4, Name: "Ringo"},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, person := range people {
			result = CleanExample(person)
		}
	}
}

func BenchmarkFastExample(b *testing.B) {
	people := []Person{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Paul"},
		{ID: 3, Name: "George"},
		{ID: 4, Name: "Ringo"},
	}
	buffer := &bytes.Buffer{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, person := range people {
			result = FastExample(buffer, person)
		}
	}
}
