package _2_loop_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Filter struct {
	data  []string
	index map[string]struct{}
}

func (f *Filter) Match(in []string) []string {
	var out []string

	for _, thisIn := range in {
		_, found := f.index[thisIn]
		if found {
			out = append(out, thisIn)
		}
	}

	return out
}

func (f *Filter) Add(in []string) {
	for _, thisIn := range in {
		f.data = append(f.data, thisIn)
		f.index[thisIn] = struct{}{}
	}
}

func TestFilter(t *testing.T) {
	filter := &Filter{
		index: map[string]struct{}{},
	}
	filter.Add([]string{"A", "B", "C", "D"})

	result := filter.Match([]string{"C", "A"})
	assert.Equal(t, 2, len(result))
}
