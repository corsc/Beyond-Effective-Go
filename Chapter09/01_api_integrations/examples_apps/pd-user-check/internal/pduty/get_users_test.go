package pduty

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserAPI_GetUsers(t *testing.T) {
	// inputs
	apiKey := getTestVarFromEnv(t, "TEST_PD_API_KEY")
	search := "corey"

	// call
	api := &UsersAPI{}
	results, resultErr := api.GetUsers(apiKey, search)

	// validate
	require.Nil(t, resultErr)
	require.NotNil(t, results)

	assert.Equal(t, 2, len(results))
}
