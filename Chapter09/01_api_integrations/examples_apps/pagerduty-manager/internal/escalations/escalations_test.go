package escalations

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
		expected              *EscalationPolicy
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(getHappyPathResponse))
			}),
			expected: &EscalationPolicy{
				ID:          "A",
				Name:        "B",
				Description: "C",
			},
			expectAnErr: false,
		},
		{
			desc: "Sad path - no such escalation",
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
			result, resultErr := manager.Get(ctx, "A")

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
		expected              *EscalationPolicy
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(listHappyPathResponse))
			}),
			expected: &EscalationPolicy{
				ID:          "A",
				Name:        "B",
				Description: "C",
			},
			expectAnErr: false,
		},
		{
			desc: "Sad path - no such escalation policy",
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
			result, resultErr := manager.GetByName(ctx, "A")

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
			expected:    "A",
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

			newEscalation := &testEscalation{
				name:        "B",
				description: "C",
				teamID:      "E",
				scheduleID:  "F",
				leadIDs:     []string{"G"},
				deptHeadIDs: []string{"H"},
			}

			// call object under test
			manager := New(cfg, logger)
			result, resultErr := manager.Add(ctx, newEscalation)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result)
		})
	}
}

func TestManager_Update(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              string
		expectAnErr           bool
	}{
		{
			desc: "Happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				if req.Method == http.MethodGet {
					resp.WriteHeader(http.StatusOK)
					_, _ = resp.Write([]byte(getHappyPathResponse))
				}

				if req.Method == http.MethodPut {
					resp.WriteHeader(http.StatusOK)
					_, _ = resp.Write([]byte(updateHappyPathResponse))
				}
			}),
			expected:    "A",
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

			escalationID := "D"

			newEscalation := &testEscalation{
				name:        "B",
				description: "C",
				teamID:      "E",
				scheduleID:  "F",
				leadIDs:     []string{"G"},
				deptHeadIDs: []string{"H"},
			}

			// call object under test
			manager := New(cfg, logger)
			resultErr := manager.Update(ctx, escalationID, newEscalation)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
		})
	}
}

type testEscalation struct {
	name        string
	description string
	teamID      string
	scheduleID  string
	leadIDs     []string
	deptHeadIDs []string
}

func (t *testEscalation) GetScheduleID() string {
	return t.scheduleID
}

func (t *testEscalation) GetLeadIDs() []string {
	return t.leadIDs
}

func (t *testEscalation) GetDeptHeadsIDs() []string {
	return t.deptHeadIDs
}

func (t *testEscalation) GetTeamName() string {
	return t.name
}

func (t *testEscalation) GetDescription() string {
	return t.description
}

func (t *testEscalation) GetTeamID() string {
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
 "escalation_policy": {
   "id": "A",
   "name": "B",
   "description": "C"
 }
}
`

var listHappyPathResponse = `
{
 "escalation_policies": [
   {
     "id": "A",
     "name": "B",
     "description": "C"
   }
 ]
}
`

var addHappyPathResponse = `
{
 "escalation_policy": {
   "id": "A",
   "name": "B",
   "description": "C"
 }
}
`

var updateHappyPathResponse = `
{
 "escalation_policy": {
   "id": "A",
   "name": "B",
   "description": "C"
 }
}
`
