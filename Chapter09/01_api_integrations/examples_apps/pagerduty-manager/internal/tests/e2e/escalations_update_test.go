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

func TestE2EEscalations_Update(t *testing.T) {
	skip.IfNotSet(t, "E2E_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	escalationID := "PPJ9DAY"

	escalation := &testEscalation{
		teamName:     "Sage42",
		scheduleID:   "PJQM2NF",
		teamID:       "PJVN6XK",
		leadIDs:      []string{"PXJHUO9"},
		deptHeadsIDs: []string{"PDPIGEC"},
	}

	// call object under test
	manager := escalations.New(cfg, logger)
	resultErr := manager.Update(ctx, escalationID, escalation)

	// validation
	require.NoError(t, resultErr)
}
