// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package _1_mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockReceiptSender is an autogenerated mock type for the ReceiptSender type
type MockReceiptSender struct {
	mock.Mock
}

// SendReceipt provides a mock function with given fields: ctx, customerEmail, amount, receiptNo
func (_m *MockReceiptSender) SendReceipt(ctx context.Context, customerEmail string, amount int64, receiptNo string) error {
	ret := _m.Called(ctx, customerEmail, amount, receiptNo)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, string) error); ok {
		r0 = rf(ctx, customerEmail, amount, receiptNo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
