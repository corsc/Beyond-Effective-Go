package _1_typical_mistake

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTypicalMistake(t *testing.T) {
	objectUnderTest := &ConcurrentObject{}

	go objectUnderTest.DoWork()

	<-time.After(1 * time.Second)

	expected := 1234
	assert.Equal(t, expected, objectUnderTest.GetResult())
}

type ConcurrentObject struct{}

func (s *ConcurrentObject) DoWork() {
	// implementation removed
}

func (s *ConcurrentObject) GetResult() int {
	// implementation removed
	return 1234
}
