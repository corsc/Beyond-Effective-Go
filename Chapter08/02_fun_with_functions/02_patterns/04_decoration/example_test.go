package _4_decoration

import (
	"io"
	"net/http"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	go StartServer()

	// give the server a chance to start
	runtime.Gosched()

	resp, err := http.DefaultClient.Get("http://0.0.0.0:8080")
	require.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello World!", string(body))
}
