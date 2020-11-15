package _2_internal_init

import (
	"time"
)

const sendEmailTimeout = 3 * time.Second

func New() *OrderManager {
	return newOrderManager(sendEmailTimeout)
}

func newOrderManager(sendTimeout time.Duration) *OrderManager {
	return &OrderManager{
		sendTimeout:    sendTimeout,
		emailTemplates: map[string]string{},
	}
}

type OrderManager struct {
	sendTimeout    time.Duration
	emailTemplates map[string]string
	sender         ReceiptSender
}

//go:generate mockery --name=ReceiptSender --case underscore --testonly --inpackage
type ReceiptSender interface {
	Send(to, body string) error
}
