package pduty

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduleAPI_GetSchedule(t *testing.T) {
	// inputs
	apiKey := getTestVarFromEnv(t, "TEST_PD_API_KEY")
	scheduleID := getTestVarFromEnv(t, "TEST_PD_SCHEDULE_ID")
	start := time.Now().Round(time.Hour)
	end := start.Add(30 * 25 * time.Hour)

	// call
	api := &ScheduleAPI{}
	result, resultErr := api.GetSchedule(apiKey, scheduleID, start, end)

	// validate
	assert.NotNil(t, result)
	assert.Nil(t, resultErr)
}
