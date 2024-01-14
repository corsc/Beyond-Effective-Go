package _6_test_resilience

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypicalImplementation(t *testing.T) {
	mockA := &MockDepA{}
	mockA.On("Do").Return(nil)

	mockB := &MockDepB{}
	mockB.On("Do").Return(errors.New("failed"))

	// not configured because we expect it not to be called
	mockC := &MockDepC{}

	unit := &Unit{
		a: mockA,
		b: mockB,
		c: mockC,
	}
	resultErr := unit.Do()

	assert.Error(t, resultErr)
}
