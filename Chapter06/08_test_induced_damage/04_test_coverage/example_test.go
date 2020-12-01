package _4_test_coverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserAsJSON(t *testing.T) {
	// inputs
	name := "Fred"
	address := "Bedrock"
	age := 35

	// call
	result, resultErr := UserAsJSON(name, address, age)

	// validation
	require.NoError(t, resultErr)
	assert.NotNil(t, result)
}
