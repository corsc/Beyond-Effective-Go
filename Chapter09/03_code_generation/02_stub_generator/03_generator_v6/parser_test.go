package _3_generator_v6

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleParser(t *testing.T) {
	filename := "testdata/input_interface.go"

	generator := &StubGenerator{}
	results, err := generator.parseSource(filename)
	require.NoError(t, err)

	assert.NotNil(t, results)
	require.Equal(t, 1, len(results))

	resultInterface := results[0]
	assert.Equal(t, "loader", resultInterface.PackageName)
	assert.Equal(t, "UserLoader", resultInterface.Name)
	require.Equal(t, 1, len(resultInterface.Methods))

	method := resultInterface.Methods[0]
	assert.Equal(t, "LoadByID", method.Name)
	assert.Equal(t, 2, len(method.Inputs))
	assert.Equal(t, 2, len(method.Outputs))
}
