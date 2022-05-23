package _3_reply_channel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkerGroup(t *testing.T) {
	// create a worker group
	totalWorkers := 3
	group := NewWorkers(totalWorkers)

	// create some tasks
	taskA := Task{
		inputA:  1,
		inputB:  2,
		replyCh: make(chan int, 1),
	}

	taskB := Task{
		inputA:  2,
		inputB:  3,
		replyCh: make(chan int, 1),
	}

	// submit tasks for processing by worker group
	group.SubmitTask(taskA)
	group.SubmitTask(taskB)

	// validate the results
	select {
	case result := <-taskA.replyCh:
		assert.Equal(t, 3, result)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout waiting for result")
	}

	select {
	case result := <-taskB.replyCh:
		assert.Equal(t, 5, result)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "timeout waiting for result")
	}
}
