package _2_string_concat

import (
	"strings"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s -benchmem ./Chapter03/05_patterns/12_string_concat/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Advanced-Go-Programming/Chapter03/05_patterns/12_string_concat
// BenchmarkExample-8                	    2000	   7335893 ns/op	53164007 B/op	    9999 allocs/op
// BenchmarkFixed-8                  	  500000	     25875 ns/op	   48504 B/op	      17 allocs/op
// BenchmarkFixedWithPreallocate-8   	 1000000	     21235 ns/op	   10240 B/op	       1 allocs/op

func BenchmarkExample(b *testing.B) {
	total := 10000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := ""

		for x := 0; x < total; x++ {
			result += "x"
		}
	}
}

func BenchmarkFixed(b *testing.B) {
	total := 10000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := &strings.Builder{}

		for x := 0; x < total; x++ {
			result.WriteString("x")
		}
		_ = result.String()
	}
}

func BenchmarkFixedWithPreallocate(b *testing.B) {
	total := 10000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := &strings.Builder{}
		result.Grow(total)

		for x := 0; x < total; x++ {
			result.WriteString("x")
		}
		_ = result.String()
	}
}
