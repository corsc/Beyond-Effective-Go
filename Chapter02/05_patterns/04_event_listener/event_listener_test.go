package _4_event_listener

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventListener(t *testing.T) {
	clock := &Clock{}

	// add some listeners
	listenerCh1 := make(chan int64, 1)
	listenerCh2 := make(chan int64)
	listenerCh3 := make(chan int64, 1)

	clock.AddListener(listenerCh1)
	clock.AddListener(listenerCh2)
	clock.AddListener(listenerCh3)

	// create event
	clock.onTick(1234)

	// validate listeners received event
	require.Equal(t, 1, len(listenerCh1))
	assert.Equal(t, int64(1234), <-listenerCh1)

	require.Equal(t, 1, len(listenerCh3))
	assert.Equal(t, int64(1234), <-listenerCh3)

	// unbuffered channel, write was skipped
	assert.Equal(t, 0, len(listenerCh2))

	// remove listener
	clock.RemoveListener(listenerCh1)
	assert.Equal(t, 2, len(clock.listeners))

	// create event
	clock.onTick(5678)

	// validate listeners received event
	require.Equal(t, 1, len(listenerCh3))
	assert.Equal(t, int64(5678), <-listenerCh3)

	// removed listener, should not received events
	assert.Equal(t, 0, len(listenerCh1))

	// unbuffered channel, write was skipped
	assert.Equal(t, 0, len(listenerCh2))
}
