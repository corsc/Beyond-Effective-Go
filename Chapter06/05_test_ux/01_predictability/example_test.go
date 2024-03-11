package _1_predictability

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		// scenarios removed
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
