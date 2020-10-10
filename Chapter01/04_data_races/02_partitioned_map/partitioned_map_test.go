package _2_partitioned_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPartitionedMap_basics(t *testing.T) {
	pMap := NewPartitionedMap(10)

	// retrieve value, should return nil and false
	result, resultFound := pMap.Get("foo")
	assert.Nil(t, result)
	assert.False(t, resultFound)

	// set value
	pMap.Set("foo", "bar")

	// retrieve value, should return the value and true
	result, resultFound = pMap.Get("foo")
	assert.Equal(t, "bar", result)
	assert.True(t, resultFound)
}
