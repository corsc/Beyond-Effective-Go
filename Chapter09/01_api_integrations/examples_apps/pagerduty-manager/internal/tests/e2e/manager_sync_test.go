package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/corsc/go-commons/testing/skip"
	pdmanager "github.com/corsc/pagerduty-manager"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestE2EManager_Sync(t *testing.T) {
	skip.IfNotSet(t, "E2E_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	// call object under test
	manager := pdmanager.New(cfg, logger)

	// parse the input file
	resultErr := manager.Parse(ctx)
	require.NoError(t, resultErr)

	// sync everything
	resultErr = manager.Sync(ctx)
	require.NoError(t, resultErr)
}
