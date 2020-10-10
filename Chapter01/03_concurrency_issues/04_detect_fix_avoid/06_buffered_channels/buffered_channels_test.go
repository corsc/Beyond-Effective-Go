package _6_buffered_channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelStats(t *testing.T) {
	bufferedCh := make(chan struct{}, 10)

	// add some data
	bufferedCh <- struct{}{}
	bufferedCh <- struct{}{}
	bufferedCh <- struct{}{}

	// validate how much data is in the channel
	assert.Equal(t, 3, len(bufferedCh))
}
