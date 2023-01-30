package pduty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAPI_GetUsers(t *testing.T) {
	// inputs
	apiKey := getTestVarFromEnv(t, "TEST_PD_API_KEY")
	userID := getTestVarFromEnv(t, "TEST_PD_USER_ID")

	entries := []*ScheduleEntry{
		{
			User: &User{
				ID: userID,
			},
		},
	}

	// call
	api := &UserAPI{}
	result, resultErr := api.GetUsers(apiKey, entries)

	// validate
	assert.NotNil(t, result)
	assert.Nil(t, resultErr)
}
