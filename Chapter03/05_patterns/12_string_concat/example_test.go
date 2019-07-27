package _2_string_concat

import (
	"strings"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s -benchmem ./Chapter03/05_patterns/12_string_concat/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/12_string_concat
// BenchmarkExample-8                	     100	 511073978 ns/op	5039953206 B/op	   10277 allocs/op
// BenchmarkFixed-8                  	  500000	     72859 ns/op	   58836 B/op	       0 allocs/op
// BenchmarkFixedWithPreallocate-8   	  300000	     46165 ns/op	   50205 B/op	       0 allocs/op

func BenchmarkExample(b *testing.B) {
	total := 10000
	result := ""

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for x := 0; x < total; x++ {
			result += "x"
		}
	}
}

func BenchmarkFixed(b *testing.B) {
	total := 10000

	result := &strings.Builder{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for x := 0; x < total; x++ {
			result.WriteString("x")
		}
		_ = result.String()
	}
}

func BenchmarkFixedWithPreallocate(b *testing.B) {
	total := 10000

	result := &strings.Builder{}
	result.Grow(total)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for x := 0; x < total; x++ {
			result.WriteString("x")
		}
		_ = result.String()
	}
}
