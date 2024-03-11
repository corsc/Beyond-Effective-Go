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
		scenarioDesc        string
		inputOrder          Order
		configureMockBank   func(bank *MockBank)
		configureMockSender func(sender *MockReceiptSender)
		expectedReceiptNo   string
		expectAnErr         bool
	}{
		{
			scenarioDesc: "Happy path (long version)",
			inputOrder: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
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
			expectedReceiptNo: "ABC-123",
			expectAnErr:       false,
		},
		{
			scenarioDesc: "Happy path (concise mocks)",
			inputOrder: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank:   happyPathBankCharge,
			configureMockSender: happyPathReceiptSend,
			expectedReceiptNo:   "ABC-123",
			expectAnErr:         false,
		},
		{
			scenarioDesc:        "Happy path (concise)",
			inputOrder:          validTestOrder,
			configureMockBank:   happyPathBankCharge,
			configureMockSender: happyPathReceiptSend,
			expectedReceiptNo:   testReceiptNo,
			expectAnErr:         false,
		},
		{
			scenarioDesc: "Sad path - customer name is missing",
			inputOrder: Order{
				CustomerName:  "", // name missing
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank:   happyPathBankCharge,
			configureMockSender: happyPathReceiptSend,
			expectedReceiptNo:   "",
			expectAnErr:         true,
		},
		{
			scenarioDesc: "Sad path - charge failed",
			inputOrder:   validTestOrder,
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("failed"))
			},
			configureMockSender: func(sender *MockReceiptSender) {},
			expectedReceiptNo:   "",
			expectAnErr:         true,
		},
		{
			scenarioDesc:      "Sad path - send receipt failed",
			inputOrder:        validTestOrder,
			configureMockBank: happyPathBankCharge,
			configureMockSender: func(sender *MockReceiptSender) {
				sender.On("SendReceipt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.New("failed"))
			},
			expectedReceiptNo: testReceiptNo,
			expectAnErr:       true,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.scenarioDesc, func(t *testing.T) {
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
			result, resultErr := orderManager.Process(ctx, scenario.inputOrder)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error. err: %s", resultErr)
			assert.Equal(t, scenario.expectedReceiptNo, result, "expected result")
		})
	}
}

var (
	validTestOrder = Order{
		CustomerName:  "Oscar",
		CustomerEmail: "me@example.com",
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
