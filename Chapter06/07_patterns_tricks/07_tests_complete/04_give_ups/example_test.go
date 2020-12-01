package _4_give_ups

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	server := &Server{}
	go server.Do()

	for {
		if server.IsDone() {
			assert.True(t, true, "test passed")
			return
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func TestServerImproved(t *testing.T) {
	server := &Server{}
	go server.Do()

	for attempt := 0; attempt < 5; attempt++ {
		if server.IsDone() {
			assert.True(t, true, "test passed")
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	assert.Fail(t, "test failed after maximum attempts")
}

type Server struct{}

func (s *Server) Do() {
	// implementation removed
}

func (s *Server) IsDone() bool {
	// implementation removed
	return true
}
