package _1_slice_only

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Collection struct {
	data []string
}

func (c *Collection) AddIfDifferent(in string) {
	for _, thisItem := range c.data {
		if thisItem == in {
			// all ready exists, skip add
			return
		}
	}

	c.data = append(c.data, in)
}

func TestCollection(t *testing.T) {
	collection := &Collection{}

	collection.AddIfDifferent("A")
	collection.AddIfDifferent("B")
	collection.AddIfDifferent("A")

	assert.Equal(t, 2, len(collection.data))
}
