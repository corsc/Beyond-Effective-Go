package _1_skipping_tests

import (
	"os/exec"
	"testing"
)

func TestSkipWithHelper(t *testing.T) {
	if !mySQLIsInstalled() {
		t.Skip("test skipped as MySQL is not installed")
	}

	// rest of the test
}

func mySQLIsInstalled() bool {
	if _, err := exec.LookPath("mysql"); err == nil {
		return true
	}

	return false
}
