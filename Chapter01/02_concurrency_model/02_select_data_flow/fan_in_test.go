package _2_select_data_flow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFanInExample(t *testing.T) {
	stopCh := make(chan struct{})
	go func() {
		defer close(stopCh)
		<-time.After(1 * time.Second)
	}()

	inputA := make(chan int, 3)
	inputA <- 1
	inputA <- 2
	inputA <- 3

	inputB := make(chan int, 3)
	inputB <- 1
	inputB <- 2
	inputB <- 3

	output := make(chan int, 10)

	FanInExample(stopCh, inputA, inputB, output)

	assert.Equal(t, 6, len(output))
}
