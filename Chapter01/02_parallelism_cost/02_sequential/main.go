package _2_sequential

func ProcessOrder(order Order) error {
	err := validateOrder(order)
	if err != nil {
		// validation failed
		return err
	}

	receiptNo, err := chargeCustomer(order)
	if err != nil {
		// payment failed
		return err
	}

	err = sendReceipt(order, receiptNo)
	if err != nil {
		return err
	}

	return informWarehouse(order)
}

func validateOrder(order Order) error {
	return nil
}

func chargeCustomer(order Order) (int, error) {
	return 0, nil
}

func sendReceipt(order Order, receiptNo int) error {
	return nil
}

func informWarehouse(order Order) error {
	return nil
}

type Order struct {
	// order details
}
