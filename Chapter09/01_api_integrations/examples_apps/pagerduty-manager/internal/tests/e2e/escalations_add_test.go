package e2e

import (
	"context"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/escalations"
	"github.com/corsc/go-commons/testing/skip"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestE2EEscalations_Add(t *testing.T) {
	skip.IfNotSet(t, "E2E_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	escalation := &testEscalation{
		teamName:     "Sage42",
		scheduleID:   "PJQM2NF",
		teamID:       "PJVN6XK",
		leadIDs:      []string{"PXJHUO9"},
		deptHeadsIDs: []string{"PDPIGEC"},
	}

	// call object under test
	manager := escalations.New(cfg, logger)
	resultID, resultErr := manager.Add(ctx, escalation)

	// validation
	require.NoError(t, resultErr)
	require.NotEmpty(t, resultID)
}

type testEscalation struct {
	teamName     string
	description  string
	scheduleID   string
	teamID       string
	leadIDs      []string
	deptHeadsIDs []string
}

func (t *testEscalation) GetTeamName() string {
	return t.teamName
}

func (t *testEscalation) GetDescription() string {
	return t.description
}

func (t *testEscalation) GetScheduleID() string {
	return t.scheduleID
}

func (t *testEscalation) GetTeamID() string {
	return t.teamID
}

func (t *testEscalation) GetLeadIDs() []string {
	return t.leadIDs
}

func (t *testEscalation) GetDeptHeadsIDs() []string {
	return t.deptHeadsIDs
}
