// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package _6_test_resilience

import (
	"github.com/stretchr/testify/mock"
)

// MockDepC is an autogenerated mock type for the DepC type
type MockDepC struct {
	mock.Mock
}

// Do provides a mock function with given fields:
func (_m *MockDepC) Do() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
