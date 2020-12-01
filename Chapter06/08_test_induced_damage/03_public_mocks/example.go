package _3_public_mocks

import (
	"context"
)

//go:generate mockery --name=Bank --case underscore --testonly --inpackage
//go:generate mockery --name=Bank --case underscore --testonly --outpkg "example" --output "./internal/example"
type Bank interface {
	Charge(ctx context.Context, customerName string, amount int64) (string, error)
}
