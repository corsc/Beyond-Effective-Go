package _1_calling_the_api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsersAPI_GetUsers(t *testing.T) {
	scenarios := []struct {
		desc           string
		mockHandler    func(resp http.ResponseWriter, req *http.Request)
		expectedResult []*User
		expectAnErr    bool
	}{
		{
			desc: "Happy path - all set",
			mockHandler: func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(`
{
	"users": [
		{
			"id": "666",
			"name": "Joan",
			"email": "joan@work.com",
			"name": "Joan",
			"contact_methods": [
				{
					"type": "email_contact_method",
					"address": "joan@work.com"
				},
				{
					"type": "phone_contact_method",
					"address": "+61234567890"
				},
				{
					"type": "sms_contact_method",
					"address": "+61234567890"
				},
				{
					"type": "push_notification_contact_method",
					"address": "joan@phone"
				}
			]
		}
	]
}
`))
			},
			expectedResult: []*User{
				{
					ID:         "666",
					Name:       "Joan",
					Email:      "joan@work.com",
					EmailIsSet: true,
					PhoneIsSet: true,
					SMSIsSet:   true,
					PushIsSet:  true,
				},
			},
			expectAnErr: false,
		},
		{
			desc: "Happy path - 1 Set",
			mockHandler: func(resp http.ResponseWriter, req *http.Request) {
				_, _ = resp.Write([]byte(`
{
	"users": [
		{
			"id": "666",
			"name": "Joan",
			"email": "joan@work.com",
			"name": "Joan",
			"contact_methods": [
				{
					"type": "email_contact_method",
					"address": "joan@work.com"
				}
			]
		}
	]
}
`))
			},
			expectedResult: []*User{
				{
					ID:         "666",
					Name:       "Joan",
					Email:      "joan@work.com",
					EmailIsSet: true,
					PhoneIsSet: false,
					SMSIsSet:   false,
					PushIsSet:  false,
				},
			},
			expectAnErr: false,
		},
		{
			desc: "Sad path - bad auth",
			mockHandler: func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusUnauthorized)
			},
			expectedResult: nil,
			expectAnErr:    true,
		},
		{
			desc: "Sad path - 5xx response",
			mockHandler: func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusInternalServerError)
			},
			expectedResult: nil,
			expectAnErr:    true,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// inputs
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			// mocks
			testServer := httptest.NewServer(http.HandlerFunc(scenario.mockHandler))
			defer testServer.Close()

			// call object under test
			objectUnderTest := NewUsersAPI(testServer.URL, "-secret-")
			result, resultErr := objectUnderTest.GetUsers(ctx)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error: %t, err: '%s'", scenario.expectAnErr, resultErr)
			assert.Equal(t, result, scenario.expectedResult)
		})
	}
}
