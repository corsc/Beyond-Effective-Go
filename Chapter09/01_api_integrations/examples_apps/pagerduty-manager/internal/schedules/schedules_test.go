package schedules

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
		expected              *Schedule
		expectErr             bool
	}{
		{
			desc: "happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(getHappyPathResponse))
			}),
			expected: &Schedule{
				ID:          "A",
				Name:        "B",
				Description: "C",
				TimeZone:    "D",
				Teams:       []*team{{ID: "E"}},
				ScheduleLayers: []*scheduleLayer{
					{
						ID:                        "AA",
						Name:                      "BB",
						Start:                     time.Date(2020, 1, 1, 1, 23, 0, 0, time.UTC),
						RotationVirtualStart:      time.Date(2020, 1, 1, 1, 23, 0, 0, time.UTC),
						RotationTurnLengthSeconds: 123,
					},
				},
			},
			expectErr: false,
		},
		{
			desc: "sad path - no such service",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(`{}`))
			}),
			expected:  nil,
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
			result, resultErr := manager.Get(ctx, "A")

			// validation
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}

func TestManager_GetByName(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              *Schedule
		expectErr             bool
	}{
		{
			desc: "happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(listHappyPathResponse))
			}),
			expected: &Schedule{
				ID:          "A",
				Name:        "B",
				Description: "C",
				TimeZone:    "D",
				Teams:       []*team{{ID: "E"}},
			},
			expectErr: false,
		},
		{
			desc: "sad path - no such escalation policy",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(`{}`))
			}),
			expected:  nil,
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
			result, resultErr := manager.GetByName(ctx, "BOOK")

			// validation
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}

func TestManager_Add(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expected              string
		expectErr             bool
	}{
		{
			desc: "happy path",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusCreated)
				_, _ = resp.Write([]byte(addHappyPathResponse))
			}),
			expected:  "A",
			expectErr: false,
		},
		{
			desc: "sad path - system error",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusInternalServerError)
			}),
			expected:  "",
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

			newSchedule := &testSchedule{
				name:         "A",
				description:  "B",
				timeZone:     "Australia/Melbourne",
				teamID:       "D",
				responderIDs: []string{"E"},
				leadIDs:      []string{"F"},
			}

			// call object under test
			manager := New(cfg, logger)
			result, resultErr := manager.Add(ctx, newSchedule, "Australia/Melbourne")

			// validation
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result)
		})
	}
}

func TestManager_Update(t *testing.T) {
	scenarios := []struct {
		desc                  string
		configureMockResponse http.HandlerFunc
		expectErr             bool
	}{
		{
			desc: "happy path",
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
			expectErr: false,
		},
		{
			desc: "sad path - system error",
			configureMockResponse: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusInternalServerError)
			}),
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

			newSchedule := &testSchedule{
				name:         "A",
				description:  "B",
				timeZone:     "Australia/Melbourne",
				teamID:       "D",
				responderIDs: []string{"E"},
				leadIDs:      []string{"F"},
			}

			// call object under test
			manager := New(cfg, logger)
			resultErr := manager.Update(ctx, "FU", newSchedule, "Australia/Melbourne")

			// validation
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error. err: %s", resultErr)
		})
	}
}

type testSchedule struct {
	name         string
	description  string
	timeZone     string
	teamID       string
	responderIDs []string
	leadIDs      []string
}

func (t *testSchedule) GetResponderIDs() []string {
	return t.responderIDs
}

func (t *testSchedule) GetLeadIDs() []string {
	return t.leadIDs
}

func (t *testSchedule) GetTimeZone() string {
	return t.timeZone
}

func (t *testSchedule) GetTeamName() string {
	return t.name
}

func (t *testSchedule) GetDescription() string {
	return t.description
}

func (t *testSchedule) GetTeamID() string {
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
  "schedule": {
	"id": "A",
	"name": "B",
	"description": "C",
	"time_zone": "D",
	"teams": [
	  {
		"ID": "E"
	  }
	],
	"schedule_layers": [
	  {
		"id": "AA",
		"name": "BB",
		"start": "2020-01-01T01:23:00Z",
		"rotation_virtual_start": "2020-01-01T01:23:00Z",
		"rotation_turn_length_seconds": 123
	  }
	]
  }
}
`

var listHappyPathResponse = `
{
 "schedules": [
   {
     "id": "A",
     "name": "B",
     "description": "C",
     "time_zone": "D",
     "teams": [
        {
          "ID": "E"
        }
     ]
   }
 ]
}
`

var addHappyPathResponse = `
{
 "schedule": {
   "id": "A",
   "name": "B",
   "description": "C"
 }
}
`

var updateHappyPathResponse = `
{
 "schedule": {
   "id": "A",
   "name": "B",
   "description": "C"
 }
}
`
