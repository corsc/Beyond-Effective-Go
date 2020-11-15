package _1_predictability

import (
	"context"
	"errors"
)

func NewOrderManager(bank Bank, sender ReceiptSender) *OrderManager {
	return &OrderManager{
		bank:   bank,
		sender: sender,
	}
}

type OrderManager struct {
	bank   Bank
	sender ReceiptSender
}

func (o *OrderManager) Process(ctx context.Context, order Order) (string, error) {
	err := o.validate(order)
	if err != nil {
		return "", err
	}

	receiptNo, err := o.chargeCustomer(ctx, order)
	if err != nil {
		return "", err
	}

	err = o.sendReceipt(ctx, order, receiptNo)
	if err != nil {
		return receiptNo, err
	}

	return receiptNo, nil
}

func (o *OrderManager) validate(order Order) error {
	if order.CustomerName == "" {
		return errors.New("customer name cannot be empty")
	}

	if order.CustomerEmail == "" {
		return errors.New("customer email cannot be empty")
	}

	if order.Amount <= 0 {
		return errors.New("order amount must be greater than 0")
	}

	return nil
}

func (o *OrderManager) chargeCustomer(ctx context.Context, order Order) (string, error) {
	return o.bank.Charge(ctx, order.CustomerName, order.Amount)
}

func (o *OrderManager) sendReceipt(ctx context.Context, order Order, receiptNo string) error {
	return o.sender.SendReceipt(ctx, order.CustomerEmail, order.Amount, receiptNo)
}

type Order struct {
	CustomerName  string
	CustomerEmail string
	Amount        int64
}

//go:generate mockery --name=Bank --case underscore --testonly --inpackage
type Bank interface {
	Charge(ctx context.Context, customerName string, amount int64) (string, error)
}

//go:generate mockery --name=ReceiptSender --case underscore --testonly --inpackage
type ReceiptSender interface {
	SendReceipt(ctx context.Context, customerEmail string, amount int64, receiptNo string) error
}
