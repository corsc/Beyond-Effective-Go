package _2_read_from_nil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromNilChannel(t *testing.T) {
	expected := []string{
		"default",
		"read from channel",
		"default",
		"read from channel",
	}

	results := readFromNilChannel()
	assert.Equal(t, expected, results)
}
