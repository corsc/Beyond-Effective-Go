package _1_go_ast

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleParser(t *testing.T) {
	filename := "input_interface.go"

	results, err := simpleParser(filename)
	require.NoError(t, err)

	assert.NotNil(t, results)

	for _, thisResult := range results {
		fmt.Printf("Name: %s\n", thisResult)
	}
}
