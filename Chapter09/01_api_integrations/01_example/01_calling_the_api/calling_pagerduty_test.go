package _1_calling_the_api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUsersAPI_GetUsers(t *testing.T) {
	scenarios := []struct {
		desc           string
		mockHandler    func(resp http.ResponseWriter, req *http.Request)
		expectedResult []*User
		expectErr      bool
	}{
		{
			desc: "Happy Path - all set",
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
			expectErr: false,
		},
		{
			desc: "Happy Path - 1 Set",
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
			expectErr: false,
		},
		{
			desc: "Sad Path - bad auth",
			mockHandler: func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusUnauthorized)
			},
			expectedResult: nil,
			expectErr:      true,
		},
		{
			desc: "Sad Path - 5xx response",
			mockHandler: func(resp http.ResponseWriter, req *http.Request) {
				resp.WriteHeader(http.StatusInternalServerError)
			},
			expectedResult: nil,
			expectErr:      true,
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
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error: %t, err: '%s'", scenario.expectErr, resultErr)
			assert.Equal(t, result, scenario.expectedResult)
		})
	}
}
