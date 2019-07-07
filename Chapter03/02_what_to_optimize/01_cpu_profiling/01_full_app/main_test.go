package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLoadGenerator(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return

		default:

		}

		req, err := http.NewRequest("GET", "http://0.0.0.0:8080", nil)
		require.NoError(t, err)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)
		_, _ = ioutil.ReadAll(resp.Body)
	}
}
