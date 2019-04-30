package _3_once

import (
	"database/sql"
	"sync"
)

type Repository struct {
	db     *sql.DB
	dbInit sync.Once
}

func (r *Repository) Dial() error {
	var err error

	// prevent the accidental creation of multiple db connection pools
	r.dbInit.Do(func() {
		r.db, err = sql.Open("mysql", "user:password@/dbname")
	})

	return err
}
