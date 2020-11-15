package _1_mocks

import (
	"context"
	"errors"
	"fmt"
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
			desc: "simple usage example",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@home.com",
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
			desc: "different responses based on inputs",
			in: Order{
				CustomerName:  "May",
				CustomerEmail: "me@home.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, "John", mock.Anything).
					Return("", errors.New("insufficient funds"))

				bank.On("Charge", mock.Anything, "May", mock.Anything).
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
			desc: "different responses based order",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@home.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("timeout")).Once()

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
			desc: "RunFunc",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@home.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						customerName := args.String(1)
						fmt.Printf("supplied customer name was: %s", customerName)
					}).
					Return("", errors.New("timeout")).Once()

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
