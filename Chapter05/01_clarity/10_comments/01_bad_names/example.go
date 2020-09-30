package _1_bad_names

func ProcessOrder(o Order) error {
	// validator order
	err := v(o)
	if err != nil {
		return err
	}

	// charge customer
	r, err := c(o)
	if err != nil {
		return err
	}

	// send receipt
	err = s(o, r)
	if err != nil {
		return err
	}

	return i(o)
}

func v(order Order) error {
	return nil
}

func c(order Order) (int, error) {
	return 0, nil
}

func s(order Order, receiptNo int) error {
	return nil
}

func i(order Order) error {
	return nil
}

type Order struct {
	// order details
}
