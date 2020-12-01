package _1_simple_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_Shutdown(t *testing.T) {
	server := &Server{}

	resultErr := server.Shutdown()
	assert.NoError(t, resultErr)
}

type Server struct{}

func (s *Server) Shutdown() error {
	// implementation  removed
	return nil
}
