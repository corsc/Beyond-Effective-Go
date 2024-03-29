package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/teams"

	"github.com/corsc/go-commons/testing/skip"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestE2ETeams_Get(t *testing.T) {
	skip.IfNotSet(t, "E2E_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	teamID := "PJVN6XK"

	// call object under test
	manager := teams.New(cfg, logger)
	result, resultErr := manager.Get(ctx, teamID)

	// validation
	require.NoError(t, resultErr)
	require.NotEmpty(t, result)
}
