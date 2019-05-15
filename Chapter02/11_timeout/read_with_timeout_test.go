package _1_timeout

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	timeout := 1 * time.Second

	closure := func() string {
		<-time.After(5 * time.Second)
		return "result"
	}

	result, resultErr := readWithTimeout(closure, timeout)
	assert.Equal(t, "", result)
	assert.Error(t, resultErr)
}
