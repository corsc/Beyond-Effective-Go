package _2_switch_outputs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer_Shutdown(t *testing.T) {
	objectUnderTest := &ConcurrentObject{}

	resultCh := make(chan int, 1)

	go objectUnderTest.DoWork(resultCh)

	select {
	case result := <-resultCh:
		expected := 1234

		assert.Equal(t, expected, result)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "test timed out")
	}
}

type ConcurrentObject struct{}

func (s *ConcurrentObject) DoWork(result chan int) {
	// implementation removed

	result <- 1234
}
