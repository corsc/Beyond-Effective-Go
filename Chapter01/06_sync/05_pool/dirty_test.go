package _5_pool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsePooledAccount(t *testing.T) {
	t.Skip("broken test, used as example")

	assert.Equal(t, 5, usePooledAccount())
	assert.Equal(t, 5, usePooledAccount())
}
