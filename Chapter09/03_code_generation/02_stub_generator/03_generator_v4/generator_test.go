package _3_generator_v4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStubGenerator_Generate(t *testing.T) {
	filename := "testdata/input_interface.go"

	generator := &StubGenerator{}
	result, err := generator.Generate(filename, "UserLoader")
	require.NoError(t, err)

	// check we’ve generated what we expected to
	assert.Equal(t, expectedStubResult, result)
}

var expectedStubResult = `
package loader

type StubUserLoader struct {}

func (s *StubUserLoader) LoadByID(ctx context.Context, userID int64) (*user.User, error) { 
	return &user.User{}, nil
}
`
