package _3_csp_usability

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheWithChannelsExample(t *testing.T) {
	cache := NewCacheUsingChannels()
	defer cache.shutdown()

	// store a value in the cache
	cache.Set("my-key", Person{Name: "Bob"})

	// retrieve a value from the cache
	resultCh := cache.Get("my-key")

	// wait for the result (or timeout)
	select {
	case result := <-resultCh:
		// happy path.  cache returned result
		expected := Person{Name: "Bob"}
		assert.Equal(t, expected, result)

	case <-time.After(1 * time.Second):
		// cache timed out
		assert.FailNow(t, "cache read timed out")
	}
}
