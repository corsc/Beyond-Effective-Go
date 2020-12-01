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

func TestExampleUATTest(t *testing.T) {
	if os.Getenv("UAT") == "" {
		t.Skip("UAT test skipped")
	}

	// rest of the test
}
