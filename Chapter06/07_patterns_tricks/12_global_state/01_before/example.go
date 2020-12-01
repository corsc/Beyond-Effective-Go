package _1_before

import (
	"context"
	"database/sql"
	"errors"
)

var db *sql.DB

type OrderRepository struct{}

func (o *OrderRepository) LoadByID(ctx context.Context, id int) (Order, error) {
	_ = db.QueryRowContext(ctx, "SELECT * FROM order WHERE ID = ?", id)

	// implementation removed
	return Order{}, errors.New("not implemented")
}

type Order struct {
	// details removed
}
