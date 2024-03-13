package pd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAPI_Post(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusCreated)
				_, _ = resp.Write([]byte(postHappyPathResponse))
			}),
			expectAnErr: false,
		},
		{
			desc: "Sad path - bad response",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusInternalServerError)
			}),
			expectAnErr: true,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// inputs
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			logger, _ := zap.NewDevelopment()

			// mocks
			testServer := httptest.NewServer(scenario.configureMockResponse)
			defer testServer.Close()

			cfg := &testConfig{
				baseURL: testServer.URL,
			}

			// call object under test
			manager := New(cfg, logger)

			uri := "/users"

			reqDTO := &newUserRequest{User: userFormat{Name: "John"}}

			respDTO := &newUserResponse{}

			resultErr := manager.Post(ctx, uri, reqDTO, respDTO)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
		})
	}
}

type newUserRequest struct {
	User userFormat `json:"user"`
}

type newUserResponse struct {
	User userFormat `json:"user"`
}

type userFormat struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	TimeZone string `json:"time_zone"`
	Role     string `json:"role"`
}

var postHappyPathResponse = `
{
  "user": {
    "name": "Fred",
    "email": "fred@flintsones.com"
  }
}
`
