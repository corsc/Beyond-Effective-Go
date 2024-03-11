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
		scenarioDesc        string
		inputOrder          Order
		configureMockBank   func(bank *MockBank)
		configureMockSender func(sender *MockReceiptSender)
		expectedReceiptNo   string
		expectAnErr         bool
	}{
		{
			scenarioDesc: "Happy path",
			inputOrder: Order{
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
			expectedReceiptNo: "ABC-123",
			expectAnErr:       false,
		},
		{
			scenarioDesc: "Sad path - customer name is missing",
			inputOrder: Order{
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
			expectedReceiptNo: "",
			expectAnErr:       true,
		},
		{
			scenarioDesc: "Sad path - charge failed",
			inputOrder: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
				Amount:        123,
			},
			configureMockBank: func(bank *MockBank) {
				bank.On("Charge", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("failed"))
			},
			configureMockSender: func(sender *MockReceiptSender) {},
			expectedReceiptNo:   "",
			expectAnErr:         true,
		},
		{
			scenarioDesc: "Sad path - send receipt failed",
			inputOrder: Order{
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
			expectedReceiptNo: "ABC-123",
			expectAnErr:       true,
		},
		{
			scenarioDesc: "Sad path - ensure no receipt is sent without a successful charge",
			inputOrder: Order{
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
			expectedReceiptNo: "",
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
