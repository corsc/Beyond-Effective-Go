package _2_channels

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer_Shutdown(t *testing.T) {
	server := &Server{}

	errCh := server.Shutdown()

	select {
	case resultErr := <-errCh:
		assert.NoError(t, resultErr)

	case <-time.After(1 * time.Second):
		assert.Fail(t, "test timed out")
	}
}

type Server struct{}

func (s *Server) Shutdown() chan error {
	errorCh := make(chan error)

	go func() {
		// implementation  removed

		errorCh <- nil
	}()

	return errorCh
}
