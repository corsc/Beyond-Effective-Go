package gcal

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalendarAPI_GetCalendars(t *testing.T) {
	// inputs
	userID := getTestVarFromEnv(t, "TEST_PD_USER_ID")
	userEmail := getTestVarFromEnv(t, "TEST_GC_USER_EMAIL")

	credentialsFile := "../../credentials.json"
	tokenFile := "../../token.json"
	users := map[string]string{
		userID: userEmail,
	}
	start := time.Now().Round(time.Hour)
	end := start.Add(30 * 25 * time.Hour)

	// call
	api := &CalendarAPI{}
	result, resultErr := api.GetCalendars(credentialsFile, tokenFile, users, start, end)

	// validate
	assert.NotNil(t, result)
	assert.Nil(t, resultErr)
}

func getTestVarFromEnv(t *testing.T, envVar string) string {
	out, found := os.LookupEnv(envVar)
	if !found {
		t.Skipf("test skipped due to lack of environment variable %s", envVar)
	}

	return out
}
