package _2_table_driven_tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestOrderManager_Process(t *testing.T) {
	scenarios := []struct {
		desc                string
		in                  Order
		configureMockBank   func(bank *MockBank)
		configureMockSender func(sender *MockReceiptSender)
		expected            string
		expectErr           bool
	}{
		{
			desc: "Happy path",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("ABC-123", nil)
			},
			configureMockSender: func(sender *MockReceiptSender) {
				sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			expected:  "ABC-123",
			expectErr: false,
		},
		{
			desc: "Sad path - customer name is missing",
			in: Order{
				CustomerName:  "", // name missing
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("ABC-123", nil)
			},
			configureMockSender: func(sender *MockReceiptSender) {
				sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			expected:  "",
			expectErr: true,
		},
		{
			desc: "Sad path - charge failed",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("failed"))
			},
			configureMockSender: func(sender *MockReceiptSender) {},
			expected:            "",
			expectErr:           true,
		},
		{
			desc: "Sad path - send receipt failed",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("ABC-123", nil)
			},
			configureMockSender: func(sender *MockReceiptSender) {
				sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("failed"))
			},
			expected:  "ABC-123",
			expectErr: true,
		},
		{
			desc: "Sad path - ensure no receipt is sent without a successful charge",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("failed"))
			},
			configureMockSender: func(sender *MockReceiptSender) {
				sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						assert.FailNow(t, "receipt should not be sent as charge failed")
					}).Return(errors.New("failed"))
			},
			expected:  "",
			expectErr: true,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// inputs
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			// mocks
			mockBank := &MockBank{}
			scenario.configureMockBank(mockBank)

			mockReceiptSender := &MockReceiptSender{}
			scenario.configureMockSender(mockReceiptSender)

			// call object under test
			orderManager := NewOrderManager(mockBank, mockReceiptSender)
			result, resultErr := orderManager.Process(ctx, scenario.in)

			// validation
			require.Equal(t, scenario.expectErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expected, result, "expected result")
		})
	}
}
