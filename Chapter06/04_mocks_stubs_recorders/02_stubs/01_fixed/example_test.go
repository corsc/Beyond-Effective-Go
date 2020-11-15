package _1_fixed

import (
	"context"
	"errors"
)

type Bank interface {
	Charge(ctx context.Context, customerName string, amount int64) (string, error)
}

type StubErrorBank struct{}

func (s *StubErrorBank) Charge(ctx context.Context, customerName string, amount int64) (string, error) {
	return "", errors.New("charge failed")
}
