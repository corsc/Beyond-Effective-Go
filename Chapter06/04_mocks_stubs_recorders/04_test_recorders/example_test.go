package _4_test_recorders

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Logger interface {
	Error(message string, args ...interface{})
}

type LogRecorder struct {
	Usage []string
}

func (r *LogRecorder) Error(message string, args ...interface{}) {
	r.Usage = append(r.Usage, fmt.Sprintf(message, args...))
}

func (r *LogRecorder) AssertContains(t *testing.T, contains interface{}, msgAndArgs ...interface{}) {
	assert.Contains(t, r.Usage, contains, msgAndArgs...)
}
