package _7_test_concurrency

import (
	"testing"
)

func TestConcurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping concurrent test because of short mode")
	}

	// rest of the test
}
