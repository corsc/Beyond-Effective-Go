package fakepkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeTestForFtest(t *testing.T) {
	// we are using this test as "test data" in the ftests package
	assert.True(t, true)
}
