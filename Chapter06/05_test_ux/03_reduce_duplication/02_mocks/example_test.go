package _2_mocks

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
			desc: "happy path (long version)",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@home.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return(testReceiptNo, nil)
			},
			configureMockSender: func(sender *MockReceiptSender) {
				sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			expected:  "ABC-123",
			expectErr: false,
		},
		{
			desc: "happy path (concise mocks)",
			in: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@home.com",
				Amount:        123,
			},
			configureMockBank:   happyPathBankCharge,
			configureMockSender: happyPathReceiptSend,
			expected:            "ABC-123",
			expectErr:           false,
		},
		{
			desc:                "happy path (concise)",
			in:                  validTestOrder,
			configureMockBank:   happyPathBankCharge,
			configureMockSender: happyPathReceiptSend,
			expected:            testReceiptNo,
			expectErr:           false,
		},
		{
			desc: "sad path - customer name is missing",
			in: Order{
				CustomerName:  "", // name missing
				CustomerEmail: "me@home.com",
				Amount:        123,
			},
			configureMockBank:   happyPathBankCharge,
			configureMockSender: happyPathReceiptSend,
			expected:            "",
			expectErr:           true,
		},
		{
			desc: "sad path - charge failed",
			in:   validTestOrder,
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("failed"))
			},
			configureMockSender: func(sender *MockReceiptSender) {},
			expected:            "",
			expectErr:           true,
		},
		{
			desc:              "sad path - send receipt failed",
			in:                validTestOrder,
			configureMockBank: happyPathBankCharge,
			configureMockSender: func(sender *MockReceiptSender) {
				sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("failed"))
			},
			expected:  testReceiptNo,
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

var (
	validTestOrder = Order{
		CustomerName:  "Oscar",
		CustomerEmail: "me@home.com",
		Amount:        123,
	}

	testReceiptNo = "ABC-123"
)

func happyPathBankCharge(bank *MockBank) {
	bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
		Return(testReceiptNo, nil)
}

func happyPathReceiptSend(sender *MockReceiptSender) {
	sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
}
