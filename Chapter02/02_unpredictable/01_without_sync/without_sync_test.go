package _1_without_sync

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAll(t *testing.T) {
	dsn, isSet := os.LookupEnv("ADV_GO_DNS")
	if !isSet {
		t.Skip("database not configured")
	}

	loader, err := NewLoader(dsn)

	require.NoError(t, err)
	loader.LoadAll()
}
