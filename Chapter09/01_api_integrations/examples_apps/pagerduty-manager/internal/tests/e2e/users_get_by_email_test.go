package e2e

import (
	"context"
	"github.com/corsc/Beyond-Effective-Go/Chapter09/01_api_integrations/examples_apps/pagerduty-manager/internal/users"
	"testing"
	"time"

	"github.com/corsc/go-commons/testing/skip"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestE2EUsers_GetByEmail(t *testing.T) {
	skip.IfNotSet(t, "E2E_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	// call object under test
	manager := users.New(cfg, logger)
	results, resultErr := manager.GetByEmail(ctx, "corey.scott@sage42.com")

	// validation
	require.NoError(t, resultErr)
	require.NotEmpty(t, results)
}
