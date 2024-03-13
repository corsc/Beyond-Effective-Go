package _1_mocks

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
			scenarioDesc: "simple usage example",
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
			scenarioDesc: "different responses based on inputs",
			inputOrder: Order{
				CustomerName:  "May",
				CustomerEmail: "me@example.com",
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
			expectedReceiptNo: "ABC-123",
			expectAnErr:       false,
		},
		{
			scenarioDesc: "different responses based order",
			inputOrder: Order{
				CustomerName:  "Oscar",
				CustomerEmail: "me@example.com",
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

func TestDecorate(t *testing.T) {
	scenarios := []struct {
		desc            string
		in              *Receipt
		configDecorator func(decorator *MockReceiptDecorator)
		expectedResult  *Receipt
		expectAnErr     bool
	}{
		{
			desc: "RunFunc example",
			in:   &Receipt{},
			configDecorator: func(decorator *MockReceiptDecorator) {
				decorator.On("Decorate", mock.Anything).
					Run(func(args mock.Arguments) {
						receipt, ok := args.Get(0).(*Receipt)
						assert.True(t, ok)

						receipt.ID = "ABC-123"
					}).
					Return(nil)
			},
			expectedResult: &Receipt{
				ID: "ABC-123",
			},
			expectAnErr: false,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// inputs

			// mocks
			thisMock := &MockReceiptDecorator{}
			scenario.configDecorator(thisMock)

			resultErr := thisMock.Decorate(scenario.in)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error: %t, err: '%s'", scenario.expectAnErr, resultErr)
			assert.Equal(t, scenario.expectedResult, scenario.in)
		})
	}

}
