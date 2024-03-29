package _1_skipping_tests

import (
	"os"
	"testing"
)

func TestSkipWithEnvVar(t *testing.T) {
	if os.Getenv("MYSQL_HOST") == "" {
		t.Skip("test skipped as MYSQL_HOST is not set")
	}

	// rest of the test
}

func TestExampleUAT(t *testing.T) {
	if os.Getenv("RUN_UAT") == "" {
		t.Skip("UAT skipped")
	}

	// rest of the test
}
