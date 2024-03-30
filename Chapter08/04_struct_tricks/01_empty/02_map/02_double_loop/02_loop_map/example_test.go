package _2_loop_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StringCollection struct {
	data  []string
	index map[string]struct{}
}

func (s *StringCollection) Match(in []string) []string {
	var out []string

	for _, thisIn := range in {
		_, found := s.index[thisIn]
		if found {
			out = append(out, thisIn)
		}
	}

	return out
}

func (s *StringCollection) Add(in []string) {
	for _, thisIn := range in {
		s.data = append(s.data, thisIn)
		s.index[thisIn] = struct{}{}
	}
}

func TestStringCollection(t *testing.T) {
	collection := &StringCollection{
		index: map[string]struct{}{},
	}
	collection.Add([]string{"A", "B", "C", "D"})

	result := collection.Match([]string{"C", "A"})
	assert.Equal(t, 2, len(result))
}
