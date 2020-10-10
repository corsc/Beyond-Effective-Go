package _2_thread_safe_slice

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThreadSafeSlice(t *testing.T) {
	tsSlice := &ThreadSafeSlice{
		data: make([]interface{}, 20),
	}

	wg := &sync.WaitGroup{}

	// create some reads and some writes
	for x := 0; x < 10; x++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			tsSlice.Put(time.Now().UnixNano())
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			_ = tsSlice.GetAll()
		}()
	}

	wg.Wait()

	// hack to ensure Put()'s have finished
	<-time.After(100 * time.Millisecond)

	results := tsSlice.GetAll()
	assert.Equal(t, 10, len(results))
}
