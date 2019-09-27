package _2_slice_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Collection struct {
	data  []string
	index map[string]struct{}
}

func (c *Collection) AddIfDifferent(in string) {
	_, exists := c.index[in]
	if exists {
		// all ready exists, skip add
		return
	}

	c.data = append(c.data, in)
	c.index[in] = struct{}{}
}

func TestCollection(t *testing.T) {
	collection := &Collection{
		index: map[string]struct{}{},
	}

	collection.AddIfDifferent("A")
	collection.AddIfDifferent("B")
	collection.AddIfDifferent("A")

	assert.Equal(t, 2, len(collection.data))
}
