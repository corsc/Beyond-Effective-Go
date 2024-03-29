package _2_higher_order

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

func TestAddPeriod(t *testing.T) {
	// variables
	in := []string{"Hello", "World"}
	expected := []string{"Hello.", "World."}

	// call
	result := forEach(in, addPeriod)

	// validate
	assert.Equal(t, expected, result)
}
