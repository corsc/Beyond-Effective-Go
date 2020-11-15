package _1_default_implementation

import (
	"errors"
	"fmt"
	"net/smtp"
)

func New(server, from string) *ReceiptSender {
	return &ReceiptSender{
		server: server,
		from:   from,
	}
}

type ReceiptSender struct {
	server string
	from   string
}

func (r *ReceiptSender) Send(to string, order Order) error {
	payload, err := r.buildReceipt(order)
	if err != nil {
		return err
	}

	err = smtp.SendMail(r.server, nil, r.from, []string{to}, payload)
	if err != nil {
		return fmt.Errorf("failed to send receipt with err: %w", err)
	}

	return nil
}

func (r *ReceiptSender) buildReceipt(order Order) ([]byte, error) {
	return nil, errors.New("not implemented")
}

type Order struct {
	// order details
}
