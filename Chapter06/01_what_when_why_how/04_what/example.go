package _4_what

import (
	"errors"
)

type OrderManager struct{}

func (o *OrderManager) Process(order Order) error {
	err := o.validate(order)
	if err != nil {
		return err
	}

	receiptNo, err := o.chargeCustomer(order)
	if err != nil {
		return err
	}

	return o.sendReceipt(order, receiptNo)
}

func (o *OrderManager) validate(order Order) error {
	if order.CustomerName == "" {
		return errors.New("customer name cannot be empty")
	}

	if order.ShippingAddress == "" {
		return errors.New("shipping address cannot be empty")
	}

	return nil
}

func (o *OrderManager) chargeCustomer(order Order) (int, error) {
	return 0, nil
}

func (o *OrderManager) sendReceipt(order Order, receiptNo int) error {
	return nil
}

type Order struct {
	CustomerName    string
	ShippingAddress string
}
