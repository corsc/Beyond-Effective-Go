package _4_optimization

import (
	"testing"
)

// $ go test -run=Bench. -bench=. ./Chapter03/04_benchmarking/04_optimization/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Advanced-Go-Programming/Chapter03/04_benchmarking/04_optimization
// BenchmarkAdd-8        	2000000000	         0.28 ns/op
// BenchmarkAddFixed-8   	2000000000	         0.28 ns/op

// $ go test -gcflags=-N -bench=. ./Chapter03/04_benchmarking/04_optimization/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Advanced-Go-Programming/Chapter03/04_benchmarking/04_optimization
// BenchmarkAdd-8        	2000000000	         1.84 ns/op
// BenchmarkAddFixed-8   	2000000000	         1.96 ns/op

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}

var result int

func BenchmarkAddFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = add(1, 2)
	}
}

func add(i int, i2 int) int {
	return i + i2
}
