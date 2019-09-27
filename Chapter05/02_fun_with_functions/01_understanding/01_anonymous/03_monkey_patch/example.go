package _3_monkey_patch

import (
	"database/sql"
	"sync"
)

var (
	db     *sql.DB
	dbInit sync.Once
)

// GetDB will return a database connection
var GetDB = func() *sql.DB {
	// prevent the accidental creation of multiple db connection pools
	dbInit.Do(func() {
		db, _ = sql.Open("mysql", "user:password@/dbname")
	})

	return db
}
