package _3_recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUpper(t *testing.T) {
	// variables
	in := []string{"Hello", "World"}
	expected := []string{"HELLO", "WORLD"}

	// call
	result := forEach(in, toUpper)

	// validate
	assert.Equal(t, expected, result)
}

func TestReverse(t *testing.T) {
	// variables
	in := []string{"Hello", "World"}
	expected := []string{"olleH", "dlroW"}

	// call
	result := forEach(in, reverse)

	// validate
	assert.Equal(t, expected, result)
}
