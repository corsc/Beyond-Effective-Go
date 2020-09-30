package _1_double_loop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Filter struct {
	data []string
}

func (f *Filter) Match(in []string) []string {
	var out []string

	for _, thisIn := range in {
		for _, thisItem := range f.data {
			if thisItem == thisIn {
				out = append(out, thisItem)
			}
		}
	}

	return out
}

func (f *Filter) Add(in []string) {
	for _, thisIn := range in {
		f.data = append(f.data, thisIn)
	}
}

func TestFilter(t *testing.T) {
	filter := &Filter{}
	filter.Add([]string{"A", "B", "C", "D"})

	result := filter.Match([]string{"C", "A"})
	assert.Equal(t, 2, len(result))
}
