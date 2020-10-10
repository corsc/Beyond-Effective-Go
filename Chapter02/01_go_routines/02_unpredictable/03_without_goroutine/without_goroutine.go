package _3_without_goroutine

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewLoader() *Loader {
	loader := &Loader{}

	loader.initializeDBPool()

	return loader
}

type Loader struct {
	// database connection pool
	db *sql.DB
}

func (l *Loader) LoadAll() error {
	results, err := l.db.Query("SELECT user FROM user")
	if err != nil {
		return err
	}

	var username string
	for results.Next() {
		err := results.Scan(&username)
		if err != nil {
			return err
		}

		fmt.Printf("User: %#v\n", username)
	}

	return nil
}

func (l *Loader) initializeDBPool() {
	var err error
	l.db, err = sql.Open("mysql", "root@tcp(localhost:3306)/mysql?autocommit=true")
	if err != nil {
		panic(err)
	}
}
