package _1_simple

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton(t *testing.T) {
	firstCall := GetCache()
	secondCall := GetCache()

	assert.Equal(t, firstCall, secondCall)
}
