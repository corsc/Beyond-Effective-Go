package pd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAPI_Get(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              *getResponse
		expectErr             bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(getHappyPathResponse))
			}),
			expected: &getResponse{
				Users: []*User{
					{
						Name:  "Fred",
						Email: "fred@flintsones.com",
						Teams: []Team{
							{
								ID:      "AAA",
								Summary: "Team AAA",
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			desc: "Sad path - bad response",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(`not JSON`))
			}),
			expected:  &getResponse{},
			expectErr: true,
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

			params := url.Values{}
			params.Set("query", url.QueryEscape("fu@bar.com"))
			params.Set("total", "false")
			params.Set("limit", "1")

			result := &getResponse{}

			resultErr := manager.Get(ctx, uri, params, result)

			// validation
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}

type getResponse struct {
	Users []*User `json:"users"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Teams []Team `json:"teams"`
}

type Team struct {
	ID      string `json:"id"`
	Summary string `json:"summary"`
}

type testConfig struct {
	baseURL string
}

func (t *testConfig) AuthToken() string {
	return os.Getenv("PD_TOKEN")
}

func (t *testConfig) Debug() bool {
	return true
}

func (t *testConfig) BaseURL() string {
	return t.baseURL
}

var getHappyPathResponse = `
{
  "users": [
    {
      "name": "Fred",
      "email": "fred@flintsones.com",
      "teams": [
        {
          "id": "AAA",
          "summary": "Team AAA"
        }
      ]
    }
  ]
}
`
