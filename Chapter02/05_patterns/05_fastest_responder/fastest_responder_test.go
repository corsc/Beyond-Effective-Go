package _5_fastest_responder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternetIsAccessible(t *testing.T) {
	assert.True(t, InternetIsAccessible())
}
