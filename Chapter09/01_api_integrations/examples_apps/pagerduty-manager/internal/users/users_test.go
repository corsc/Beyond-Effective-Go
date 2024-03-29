package users

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestManager_GetByEmail(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              *User
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(getHappyPathResponse))
			}),
			expected: &User{
				Name:  "Fred",
				Email: "fred@flintsones.com",
				Teams: []Team{
					{
						ID:      "AAA",
						Summary: "Team AAA",
					},
				},
			},
			expectAnErr: false,
		},
		{
			desc: "Sad path - no user found",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(`{}`))
			}),
			expected:    nil,
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
			result, resultErr := manager.GetByEmail(ctx, "fu@bar.com")

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}

func TestManager_Add(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              string
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusCreated)
				_, _ = resp.Write([]byte(addHappyPathResponse))
			}),
			expected:    "JOAN",
			expectAnErr: false,
		},
		{
			desc: "Sad path - system error",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusInternalServerError)
			}),
			expected:    "",
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

			user := &testUser{
				name:     "Joan",
				email:    "joan@ark.org",
				timeZone: "Europe/Paris",
				role:     "manager",
			}

			// call object under test
			manager := New(cfg, logger)
			result, resultErr := manager.Add(ctx, user, "Australia/Melbourne")

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}

type testUser struct {
	name     string
	email    string
	timeZone string
	role     string
}

func (t *testUser) GetName() string {
	return t.name
}

func (t *testUser) GetEmail() string {
	return t.email
}

func (t *testUser) GetTimeZone() string {
	return t.timeZone
}

func (t *testUser) GetUserRole() string {
	return t.role
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

var addHappyPathResponse = `
{
  "user": {
    "id": "JOAN",
    "name": "Joan",
    "email": "joan@ark.org"
  }
}
`
