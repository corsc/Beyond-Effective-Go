package _1_preallocate

import (
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s ./Chapter03/05_patterns/11_preallocate/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Beyond-Effective-Go/Chapter03/05_patterns/11_preallocate
// cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
// BenchmarkExample-8   	     145	  85199634 ns/op
// BenchmarkFixed-8     	    1029	  11776375 ns/op

func BenchmarkExample(b *testing.B) {
	total := 1000000

	for i := 0; i < b.N; i++ {
		var data []string

		for x := 0; x < total; x++ {
			data = append(data, "x")
		}
	}
}

func BenchmarkFixed(b *testing.B) {
	total := 1000000

	for i := 0; i < b.N; i++ {
		data := make([]string, total)

		for x := 0; x < total; x++ {
			data[x] = "x"
		}
	}
}
