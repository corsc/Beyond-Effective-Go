package _2_fixed

func ProcessOrder(order Order) error {
	err := validateOrder(order)
	if err != nil {
		return err
	}

	receiptNo, err := chargeCustomer(order)
	if err != nil {
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
