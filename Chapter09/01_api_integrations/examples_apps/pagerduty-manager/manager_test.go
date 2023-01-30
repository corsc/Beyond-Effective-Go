package pdmanager

import (
	"context"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/stretchr/testify/require"
)

func TestManager_Parse(t *testing.T) {
	scenarios := []struct {
		desc      string
		in        string
		expectErr bool
	}{
		{
			desc:      "happy path",
			in:        "./test_data/simple.json",
			expectErr: false,
		},
		{
			desc:      "sad path - empty file",
			in:        "./test_data/empty.json",
			expectErr: true,
		},
		{
			desc:      "sad path - invalid file",
			in:        "./test_data/invalid.json",
			expectErr: true,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// inputs
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			cfg := &testConfig{
				filename: scenario.in,
			}

			logger, _ := zap.NewDevelopment()

			// call object under test
			manager := New(cfg, logger)
			resultErr := manager.Parse(ctx)

			// validation
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error. err: %s", resultErr)
		})
	}
}

type testConfig struct {
	filename string
}

func (t *testConfig) BaseURL() string {
	return ""
}

func (t *testConfig) AuthToken() string {
	return ""
}

func (t *testConfig) Debug() bool {
	return true
}

func (t *testConfig) Filename() string {
	return t.filename
}
