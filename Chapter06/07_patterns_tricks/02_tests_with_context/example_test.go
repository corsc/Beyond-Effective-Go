package _2_tests_with_context

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTempting(t *testing.T) {
	resultErr := do(context.Background())

	require.NoError(t, resultErr)
}

func TestImproved(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resultErr := do(ctx)

	require.NoError(t, resultErr)
}

func do(ctx context.Context) error {
	return nil
}
