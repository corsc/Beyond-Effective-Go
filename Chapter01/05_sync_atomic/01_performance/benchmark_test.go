package _1_performance

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var simpleTotal int64

// BenchmarkSimpleAdd-8   	1000	   1985808 ns/op
func BenchmarkSimpleAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 1000000; x++ {
			simpleTotal += int64(x)
		}
		fmt.Printf("total: %d\n", simpleTotal)
	}
}

var atomicTotal int64

// BenchmarkAtomicAdd-8	300	   	   5367699 ns/op
func BenchmarkAtomicAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 1000000; x++ {
			atomic.AddInt64(&atomicTotal, int64(x))
		}
		fmt.Printf("total: %d\n", atomicTotal)
	}
}

var mutexTotal int64
var mutex sync.Mutex

// BenchmarkMutexAdd-8	100	   	   14997051 ns/op
func BenchmarkMutexAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for x := 0; x < 1000000; x++ {
			mutex.Lock()
			mutexTotal += int64(x)
			mutex.Unlock()
		}
		fmt.Printf("total: %d\n", mutexTotal)
	}
}
