package _2_thread_safe_slice

import (
	"runtime"
	"sync/atomic"
)

type ThreadSafeSlice struct {
	data     []interface{}
	reserved int64
	done     int64
}

func (t *ThreadSafeSlice) Put(value interface{}) {
	// reserve an index
	index := atomic.AddInt64(&t.reserved, 1) - 1

	// store the data
	t.data[index] = value

	// wait and "eventually" update the done counter
	for !atomic.CompareAndSwapInt64(&t.done, index, index+1) {
		// yield and let other goroutines run
		runtime.Gosched()
	}
}

func (t *ThreadSafeSlice) GetAll() []interface{} {
	currentDone := atomic.LoadInt64(&t.done)
	return t.data[:currentDone]
}
