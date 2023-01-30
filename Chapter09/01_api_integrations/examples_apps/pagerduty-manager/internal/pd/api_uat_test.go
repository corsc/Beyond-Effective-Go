package pd

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/corsc/go-commons/testing/skip"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAPI_Get_UAT(t *testing.T) {
	skip.IfNotSet(t, "UAT_TEST")

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger, _ := zap.NewDevelopment()

	cfg := &testConfig{
		baseURL: "https://api.pagerduty.com",
	}

	uri := "/users"

	params := url.Values{}
	params.Set("query", url.QueryEscape("nasa.example"))
	params.Set("total", "false")
	params.Set("limit", "1")

	result := &getResponse{}

	// call object under test
	manager := New(cfg, logger)
	resultErr := manager.Get(ctx, uri, params, result)

	// validation
	require.NoError(t, resultErr)
	require.NotNil(t, result)
	assert.True(t, len(result.Users) > 0)
}
