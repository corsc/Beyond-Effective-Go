package _2_after

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_WaitForMessage_HappyPath(t *testing.T) {
	// inputs
	data := make(chan *Message, 1)
	data <- &Message{}

	// call object under test
	client := &Client{}
	result, resultErr := client.WaitForMessage(data)

	// validation
	require.NoError(t, resultErr)
	assert.NotNil(t, result)
}

func TestClient_WaitForMessage_SadPath(t *testing.T) {
	// inputs (intentionally no data)
	data := make(chan *Message, 1)

	// call object under test
	client := &Client{
		waitTime: 10 * time.Millisecond,
	}
	result, resultErr := client.WaitForMessage(data)

	// validation
	require.Error(t, resultErr)
	assert.Nil(t, result)
}
