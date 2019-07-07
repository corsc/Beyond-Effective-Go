package _2_test_based

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/02_what_to_optimize/01_cpu_profiling/shared"
	"github.com/stretchr/testify/require"
)

func BenchmarkLoadGenerator(b *testing.B) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	helloRequest := &shared.HelloRequest{
		Name: fmt.Sprintf("User %d", time.Now().UnixNano()),
	}

	for i := 0; i < b.N; i++ {
		buffer.Reset()
		err := encoder.Encode(helloRequest)
		require.NoError(b, err)

		req := httptest.NewRequest("GET", "/", buffer)

		resp := httptest.NewRecorder()
		shared.CardShuffler(resp, req)

		require.Equal(b, http.StatusOK, resp.Code)
		ioutil.ReadAll(resp.Body)
	}
}

func TestLoadGenerator(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	helloRequest := &shared.HelloRequest{
		Name: fmt.Sprintf("User %d", time.Now().UnixNano()),
	}

	for {
		select {
		case <-ctx.Done():
			return

		default:

		}

		buffer.Reset()
		err := encoder.Encode(helloRequest)
		require.NoError(t, err)

		req := httptest.NewRequest("GET", "/", buffer)

		resp := httptest.NewRecorder()
		shared.CardShuffler(resp, req)

		require.Equal(t, http.StatusOK, resp.Code)
		ioutil.ReadAll(resp.Body)
	}
}
