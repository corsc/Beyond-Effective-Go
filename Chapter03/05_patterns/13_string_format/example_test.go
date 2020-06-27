package _3_string_format

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

// $ go test -run=Bench. -bench=. -benchtime=10s -benchmem ./Chapter03/05_patterns/13_string_format/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Advanced-Go-Programming/Chapter03/05_patterns/13_string_format
// BenchmarkExample-8   	100000000	       159 ns/op	      24 B/op	       2 allocs/op
// BenchmarkStrconv-8   	100000000	       117 ns/op	      31 B/op	       2 allocs/op
// BenchmarkAppend-8    	100000000	       103 ns/op	      19 B/op	       2 allocs/op
// BenchmarkBuilder-8   	100000000	       115 ns/op	      31 B/op	       2 allocs/op

var result string

func BenchmarkExample(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%d ns", time.Since(start).Nanoseconds())
	}
}

func BenchmarkStrconv(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = strconv.FormatInt(time.Since(start).Nanoseconds(), 10) + " ns"
	}
}

func BenchmarkAppend(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var dest []byte
		dest = strconv.AppendInt(dest, time.Since(start).Nanoseconds(), 10)
		dest = append([]byte(" ns"))
		result = string(dest)
	}
}

func BenchmarkBuilder(b *testing.B) {
	start := time.Now()

	b.ResetTimer()

	builder := &strings.Builder{}

	for i := 0; i < b.N; i++ {
		builder.WriteString(strconv.FormatInt(time.Since(start).Nanoseconds(), 10))
		builder.WriteString(" ns")
		result = string(builder.String())

		builder.Reset()
	}
}
