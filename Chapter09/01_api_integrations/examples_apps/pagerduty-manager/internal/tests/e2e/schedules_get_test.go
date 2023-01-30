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

func TestE2ESchedules_Get(t *testing.T) {
	skip.IfNotSet(t, "E2E_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	scheduleID := "PJQM2NF"

	// call object under test
	manager := schedules.New(cfg, logger)
	result, resultErr := manager.Get(ctx, scheduleID)

	// validation
	require.NoError(t, resultErr)
	require.NotEmpty(t, result)
}
