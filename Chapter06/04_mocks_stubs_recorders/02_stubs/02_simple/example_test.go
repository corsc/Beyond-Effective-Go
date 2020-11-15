package _2_simple

import (
	"context"
)

type Bank interface {
	Charge(ctx context.Context, customerName string, amount int64) (string, error)
}

type StubErrorBank struct {
	receiptNo string
	err       error
}

func (s *StubErrorBank) Charge(ctx context.Context, customerName string, amount int64) (string, error) {
	return s.receiptNo, s.err
}
