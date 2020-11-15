package _2_short

import (
	"testing"
)

func TestShortExample(t *testing.T) {
	if testing.Short() {
		t.Skip("test skipped due to short mode")
	}

	// normal test implementation
}
