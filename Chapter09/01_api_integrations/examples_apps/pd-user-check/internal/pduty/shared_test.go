package pduty

import (
	"os"
	"testing"
)

func getTestVarFromEnv(t *testing.T, envVar string) string {
	out, found := os.LookupEnv(envVar)
	if !found {
		t.Skipf("test skipped due to lack of environment variable %s", envVar)
	}

	return out
}
