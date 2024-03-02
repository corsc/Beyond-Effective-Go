package _3_builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Option 1: Creating a compile-time assertion
var _ File = &fileV1{}

func TestFileV1_implementsFile(t *testing.T) {
	// Option 2: using a test and the Testify assert library
	assert.Implements(t, (*File)(nil), new(fileV1))
}
