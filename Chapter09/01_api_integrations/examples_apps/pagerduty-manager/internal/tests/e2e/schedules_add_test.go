package e2e

import (
	"context"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/schedules"
	"github.com/corsc/go-commons/testing/skip"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestE2ESchedules_Add(t *testing.T) {
	skip.IfNotSet(t, "E2E_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	timeZone := "Australia/Melbourne"

	schedule := &testSchedule{
		teamName:     "Sage42",
		teamID:       "PJVN6XK",
		responderIDs: []string{"PDPIGEC"},
		leadIDs:      []string{"PXJHUO9"},
	}

	// call object under test
	manager := schedules.New(cfg, logger)
	resultID, resultErr := manager.Add(ctx, schedule, timeZone)

	// validation
	require.NoError(t, resultErr)
	require.NotEmpty(t, resultID)
}

type testSchedule struct {
	teamName     string
	description  string
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

func (t *testSchedule) GetTeamName() string {
	return t.teamName
}

func (t *testSchedule) GetDescription() string {
	return t.description
}

func (t *testSchedule) GetTeamID() string {
	return t.teamID
}