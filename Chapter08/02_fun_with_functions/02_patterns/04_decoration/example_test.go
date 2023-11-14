package _4_decoration

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestExample(t *testing.T) {
	go StartServer()

	resp, err := http.DefaultClient.Get("http://0.0.0.0:8080")
	require.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello World!", string(body))
}
