package skip

import (
	"fmt"
	"os"
)

// Skipper is a subset of testing.T
type Skipper interface {
	// Skip the current test
	Skip(args ...interface{})
}

// IfNotSet will skip the current test when the supplied environment variable is not set
func IfNotSet(t Skipper, key string) {
	if os.Getenv(key) == "" {
		t.Skip(fmt.Sprintf("skiped as %s is not set", key))
	}
}
