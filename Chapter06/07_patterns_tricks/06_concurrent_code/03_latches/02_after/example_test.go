package _3_latches

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConsumer_Consume(t *testing.T) {
	resultCh := make(chan *Order, 1)

	consumer := &Consumer{
		orderDropped: func(order *Order) {
			resultCh <- order
		},
	}

	ordersCh := make(chan *Order, 1)
	ordersCh <- newInvalidOrder()

	go consumer.Consume(ordersCh)

	select {
	case <-resultCh:
		assert.True(t, true, "Happy path")

	case <-time.After(1 * time.Second):
		assert.Fail(t, "test timed out")
	}
}

func newInvalidOrder() *Order {
	// implementation removed
	return nil
}
