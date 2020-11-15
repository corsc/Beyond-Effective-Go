package _2_private_constructor

import (
	"database/sql"
	"errors"
)

var dbPool *sql.DB

func New() *Repository {
	return newRepository(dbPool)
}

func newRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) LoadByID(id int) (*Order, error) {
	return nil, errors.New("not implemented")
}

type Order struct {
	// order details
}
