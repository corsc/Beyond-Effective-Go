package _1_preallocate

import (
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s ./Chapter03/05_patterns/11_preallocate/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Beyond-Effective-Go/Chapter03/05_patterns/11_preallocate
// BenchmarkExample-8   	     100	 120309432 ns/op
// BenchmarkFixed-8     	     200	  89336293 ns/op

func BenchmarkExample(b *testing.B) {
	total := 1000000
	var data []string

	for i := 0; i < b.N; i++ {
		for x := 0; x < total; x++ {
			data = append(data, "x")
		}
	}
}

func BenchmarkFixed(b *testing.B) {
	total := 1000000
	data := make([]string, total)

	for i := 0; i < b.N; i++ {
		for x := 0; x < total; x++ {
			data = append(data, "x")
		}
	}
}
