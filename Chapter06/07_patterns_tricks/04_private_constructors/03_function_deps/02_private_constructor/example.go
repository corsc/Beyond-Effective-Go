package _2_private_constructor

import (
	"errors"
	"fmt"
	"net/smtp"
)

type sendMail func(addr string, a smtp.Auth, from string, to []string, msg []byte) error

func New(server, from string) *ReceiptSender {
	return newReceiptSender(server, from, smtp.SendMail)
}

func newReceiptSender(server, from string, sendFunc sendMail) *ReceiptSender {
	return &ReceiptSender{
		server:   server,
		from:     from,
		sendFunc: sendFunc,
	}
}

type ReceiptSender struct {
	server   string
	from     string
	sendFunc sendMail
}

func (r *ReceiptSender) Send(to string, order Order) error {
	payload, err := r.buildReceipt(order)
	if err != nil {
		return err
	}

	err = r.sendFunc(r.server, nil, r.from, []string{to}, payload)
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
