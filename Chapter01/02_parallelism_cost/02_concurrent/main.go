package _2_concurrent

import (
	"sync"
)

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

	errorCh := make(chan error, 2)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sendReceipt(order, receiptNo, errorCh)

	wg.Add(1)
	go informWarehouse(order, errorCh)

	wg.Wait()

	for resultErr := range errorCh {
		// return the first detected error
		if resultErr != nil {
			return err
		}
	}

	return nil
}

func validateOrder(order Order) error {
	return nil
}

func chargeCustomer(order Order) (int, error) {
	return 0, nil
}

func sendReceipt(order Order, receiptNo int, errorCh chan<- error) {
	errorCh <- nil
}

func informWarehouse(order Order, errorCh chan<- error) {
	errorCh <- nil
}

type Order struct {
	// order details
}
