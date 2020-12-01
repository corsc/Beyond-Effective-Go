package _8_parallel_benchmarks

import (
	"testing"
)

func BenchmarkDoSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoSomething()
	}
}

func BenchmarkDoSomethingParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DoSomething()
		}
	})
}

func DoSomething() {
	// implementation removed
}
