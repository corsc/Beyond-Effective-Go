package _3_once

import (
	"database/sql"
	"reflect"
	"sync"
)

type User struct {
	ID    int    `sql:"id"`
	Name  string `sql:"name"`
	Email string `sql:"email"`
}

type UserDAO struct {
	insertSQL     string
	insertSQLOnce sync.Once
}

func (u *UserDAO) Save(db *sql.DB, user User) error {
	query := u.generateInsertSQL()

	_, err := db.Exec(query, user.ID, user.Name, user.Email)
	return err
}

func (u *UserDAO) SaveFixed(db *sql.DB, user User) error {
	u.insertSQLOnce.Do(func() {
		u.insertSQL = u.generateInsertSQL()
	})

	_, err := db.Exec(u.insertSQL, user.ID, user.Name, user.Email)
	return err
}

// build select query from the struct tags
func (u *UserDAO) generateInsertSQL() string {
	out := "INSERT INTO user ("

	typ := reflect.TypeOf(User{})

	for x := 0; x < typ.NumField(); x++ {
		if x > 0 {
			out += ","
		}
		out += typ.Field(x).Tag.Get("sql")
	}

	out += ") VALUES ("

	for x := 0; x < typ.NumField(); x++ {
		if x > 0 {
			out += ","
		}
		out += "?"
	}

	out += ")"

	return out
}
