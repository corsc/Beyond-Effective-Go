// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package _3_public_mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockBank is an autogenerated mock type for the Bank type
type MockBank struct {
	mock.Mock
}

// Charge provides a mock function with given fields: ctx, customerName, amount
func (_m *MockBank) Charge(ctx context.Context, customerName string, amount int64) (string, error) {
	ret := _m.Called(ctx, customerName, amount)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) string); ok {
		r0 = rf(ctx, customerName, amount)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, customerName, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
