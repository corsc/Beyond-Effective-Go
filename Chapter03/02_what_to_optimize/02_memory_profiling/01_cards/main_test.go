package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLoadGenerator(t *testing.T) {
	//t.Skip("test skipped because it should only be used for load generation/profiling")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return

		default:

		}

		wg := &sync.WaitGroup{}

		for x := 0; x < 20; x++ {
			wg.Add(1)

			go func() {
				defer wg.Done()

				req, err := http.NewRequest("GET", "http://0.0.0.0:8888", nil)
				require.NoError(t, err)

				resp, err := (&http.Client{}).Do(req)
				require.NoError(t, err)

				require.Equal(t, http.StatusOK, resp.StatusCode)
				_, _ = ioutil.ReadAll(resp.Body)
			}()
		}

		wg.Wait()
	}
}
