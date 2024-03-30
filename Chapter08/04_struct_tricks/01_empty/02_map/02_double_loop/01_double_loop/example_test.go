package _1_double_loop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StringCollection struct {
	data []string
}

func (s *StringCollection) Match(in []string) []string {
	var out []string

	for _, thisIn := range in {
		for _, thisItem := range s.data {
			if thisItem == thisIn {
				out = append(out, thisItem)
			}
		}
	}

	return out
}

func (s *StringCollection) Add(in []string) {
	for _, thisIn := range in {
		s.data = append(s.data, thisIn)
	}
}

func TestStringCollection(t *testing.T) {
	collection := &StringCollection{}
	collection.Add([]string{"A", "B", "C", "D"})

	result := collection.Match([]string{"C", "A"})
	assert.Equal(t, 2, len(result))
}
