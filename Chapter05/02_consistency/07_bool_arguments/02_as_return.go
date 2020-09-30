package _7_bool_arguments

import (
	"errors"
	"fmt"
)

func AsReturnExample() error {
	orderID := "QQQ111"
	productCode := "ABC123"
	additionalQty := 3

	newQty, exists, err := addToOrderQuantity(orderID, productCode, additionalQty)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("failed to load order with id %s", orderID)
	}

	// debug
	fmt.Printf("New Quanity: %d\n", newQty)

	return nil
}

func addToOrderQuantity(orderID string, productCode string, qty int) (int, bool, error) {
	return 0, false, errors.New("not implemented")
}
