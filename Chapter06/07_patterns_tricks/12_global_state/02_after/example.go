package _2_after

import (
	"context"
	"database/sql"
	"errors"
)

var db *sql.DB

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

type OrderRepository struct {
	db *sql.DB
}

func (o *OrderRepository) LoadByID(ctx context.Context, id int) (Order, error) {
	_ = o.db.QueryRowContext(ctx, "SELECT * FROM order WHERE ID = ?", id)

	// implementation removed
	return Order{}, errors.New("not implemented")
}

type Order struct {
	// details removed
}
