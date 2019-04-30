package _6_csp_usability

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheWithM(t *testing.T) {
	cache := newCacheUsingMutex()

	// store a value in the cache
	cache.Set("my-key", Person{Name: "Bob"})

	// retrieve a value from the cache
	result := cache.Get("my-key")

	expected := Person{Name: "Bob"}
	assert.Equal(t, expected, result)
}
