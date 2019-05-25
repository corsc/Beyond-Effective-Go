package _5_select_data_flow

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFanOutExample(t *testing.T) {
	input := make(chan int, 3)
	input <- 1
	input <- 2
	input <- 3
	go func() {
		defer close(input)
		<-time.After(1 * time.Second)
	}()

	outputA := make(chan int, 10)
	outputB := make(chan int, 10)

	FanOutExample(input, outputA, outputB)

	assert.Equal(t, 3, len(outputA)+len(outputB))
}
