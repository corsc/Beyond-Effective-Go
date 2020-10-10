package _2_batching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBatchedSender_Send(t *testing.T) {
	var data []Item
	for x := 0; x < batchSize; x++ {
		data = append(data, Item{ID: x})
	}

	resultCh := make(chan [batchSize]Item, 10)

	bs := &BatchedSender{}
	for _, item := range data {
		bs.Send(resultCh, item)
	}

	assert.Equal(t, 1, len(resultCh))
}
