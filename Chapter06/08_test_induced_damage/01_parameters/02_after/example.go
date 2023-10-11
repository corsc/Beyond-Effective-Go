package _2_after

import (
	"errors"
	"time"
)

const maxWaitTime = 3 * time.Second

func NewClient() *Client {
	return &Client{
		waitTime: maxWaitTime,
	}
}

type Client struct {
	waitTime time.Duration
}

func (c *Client) WaitForMessage(data chan *Message) (*Message, error) {
	select {
	case msg := <-data:
		return msg, nil

	case <-time.After(c.waitTime):
		return nil, errors.New("timed out")
	}
}

type Message struct {
	// details removed
}
