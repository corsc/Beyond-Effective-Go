package _2_template_example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleTemplate(t *testing.T) {
	result, err := SimpleTemplate()
	require.NoError(t, err)

	expected := `Good morning Craig.
The current time is: 2000-02-01T02:03:04Z`

	assert.Equal(t, expected, result)
}
