package _9_assert_interface

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ AnInterface = &AnObject{}

func TestMyObject_implements(t *testing.T) {
	assert.Implements(t, (*AnInterface)(nil), &AnObject{})
}

type AnInterface interface {
	Name() string
}

type AnObject struct{}

func (a *AnObject) Name() string {
	return ""
}
