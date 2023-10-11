package _1_before

import (
	"errors"
	"time"
)

const maxWaitTime = 3 * time.Second

type Client struct{}

func (c *Client) WaitForMessage(data chan *Message) (*Message, error) {
	select {
	case msg := <-data:
		return msg, nil

	case <-time.After(maxWaitTime):
		return nil, errors.New("timed out")
	}
}

type Message struct {
	// details removed
}
