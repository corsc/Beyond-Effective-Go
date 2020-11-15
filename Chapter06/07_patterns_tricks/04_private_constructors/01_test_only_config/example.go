package _1_test_only_config

import (
	"time"
)

const sendEmailTimeout = 3 * time.Second

func New() *OrderManager {
	return newOrderManager(sendEmailTimeout)
}

func newOrderManager(sendTimeout time.Duration) *OrderManager {
	return &OrderManager{
		sendTimeout: sendTimeout,
	}
}

type OrderManager struct {
	sendTimeout time.Duration
}
