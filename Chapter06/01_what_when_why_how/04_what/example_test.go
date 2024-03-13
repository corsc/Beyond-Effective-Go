package _4_what

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrderManager_Process_happyPath(t *testing.T) {
	// inputs
	order := Order{
		CustomerName:    "Oscar",
		ShippingAddress: "123 Sesame Street",
	}
	var expectedErr error = nil

	// call object under test
	orderManager := &OrderManager{}
	resultErr := orderManager.Process(order)

	// validation
	require.Equal(t, expectedErr, resultErr)
}

func TestOrderManager_Process_sadPath_nameValidation(t *testing.T) {
	// inputs
	order := Order{
		CustomerName: "",
	}
	expectedErr := errors.New("customer name cannot be empty")

	// call object under test
	orderManager := &OrderManager{}
	resultErr := orderManager.Process(order)

	// validation
	require.Equal(t, expectedErr, resultErr)
}

func TestOrderManager_Process_sadPath_nameValidationImproved(t *testing.T) {
	// inputs
	order := Order{
		CustomerName:    "",
		ShippingAddress: "123 Sesame Street",
	}
	expectAnErr := true

	// call object under test
	orderManager := &OrderManager{}
	resultErr := orderManager.Process(order)

	// validation
	require.Equal(t, expectAnErr, resultErr != nil)
}

func TestOrderManager_Process_sadPath_shippingAddressValidation(t *testing.T) {
	// inputs
	order := Order{
		CustomerName:    "Oscar",
		ShippingAddress: "",
	}
	expectedErr := errors.New("shipping address cannot be empty")

	// call object under test
	orderManager := &OrderManager{}
	resultErr := orderManager.Process(order)

	// validation
	require.Equal(t, expectedErr, resultErr)
}
