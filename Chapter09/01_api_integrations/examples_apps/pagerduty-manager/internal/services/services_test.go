package services

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

func TestManager_Get(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              *Service
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(getHappyPathResponse))
			}),
			expected: &Service{
				ID:          "BOOK",
				Name:        "The Booking Policy",
				Description: "",
			},
			expectAnErr: false,
		},
		{
			desc: "Sad path - no such service",
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
			result, resultErr := manager.Get(ctx, "BOOK")

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}

func TestManager_GetByName(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              *Service
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(listHappyPathResponse))
			}),
			expected: &Service{
				ID:          "BOOK",
				Name:        "The Booking Policy",
				Description: "",
			},
			expectAnErr: false,
		},
		{
			desc: "Sad path - no such service",
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
			result, resultErr := manager.GetByName(ctx, "BOOK")

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
			expected:    "BOOK",
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

			newService := &testService{
				name:        "A",
				description: "B",
			}

			newTeam := &testTeam{
				teamID:             "C",
				escalationPolicyID: "D",
			}

			// call object under test
			manager := New(cfg, logger)
			result, resultErr := manager.Add(ctx, newService, newTeam)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result)
		})
	}
}

type testService struct {
	name        string
	description string
}

func (t *testService) GetName() string {
	return t.name
}

func (t *testService) GetDescription() string {
	return t.description
}

type testTeam struct {
	teamID             string
	escalationPolicyID string
}

func (t *testTeam) GetEscalationPolicyID() string {
	return t.escalationPolicyID
}

func (t *testTeam) GetTeamID() string {
	return t.teamID
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
 "service": {
   "id": "BOOK",
   "name": "The Booking Policy"
 }
}
`

var listHappyPathResponse = `
{
 "services": [
   {
     "id": "BOOK",
     "name": "The Booking Policy"
   }
 ]
}
`

var addHappyPathResponse = `
{
 "service": {
   "id": "BOOK",
   "name": "The Booking Policy"
 }
}
`
