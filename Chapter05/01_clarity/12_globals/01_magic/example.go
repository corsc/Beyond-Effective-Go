package _1_magic

import (
	"context"
	"database/sql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "user:password@localhost/db")
	if err != nil {
		panic(err)
	}
}

func LoadUser(ctx context.Context, ID int) (*User, error) {
	row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", ID)

	out := &User{}
	err := row.Scan(out.Name)
	if err != nil {
		return nil, err
	}

	return out, nil
}

type User struct {
	Name string
}
