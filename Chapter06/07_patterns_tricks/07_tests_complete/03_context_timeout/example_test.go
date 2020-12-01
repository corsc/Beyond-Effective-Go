package _3_context_timeout

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBadExample(t *testing.T) {
	resultErr := DoSomething(context.Background(), "ABC-123")
	assert.NoError(t, resultErr)
}

func TestGoodExample(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resultErr := DoSomething(ctx, "ABC-123")
	assert.NoError(t, resultErr)
}

func DoSomething(ctx context.Context, in string) error {
	// implementation removed
	return nil
}
