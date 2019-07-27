package _4_once_again

import (
	"database/sql"
	"reflect"
)

type User struct {
	ID    int    `sql:"id"`
	Name  string `sql:"name"`
	Email string `sql:"email"`
}

type UserDAO struct {
}

func (u *UserDAO) Save(db *sql.DB, user User) error {
	query := u.generateInsertSQL()

	_, err := db.Exec(query, user.ID, user.Name, user.Email)
	return err
}

func (u *UserDAO) SaveFixed(db *sql.DB, user User) error {
	query := u.generateInsertSQLFixed()

	_, err := db.Exec(query, user.ID, user.Name, user.Email)
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

// build select query from the struct tags
func (u *UserDAO) generateInsertSQLFixed() string {
	out := "INSERT INTO user ("

	typ := reflect.TypeOf(User{})

	values := ""

	for x := 0; x < typ.NumField(); x++ {
		if x > 0 {
			out += ","
			values += ","
		}
		out += typ.Field(x).Tag.Get("sql")
		values += "?"
	}

	out += ") VALUES (" + values + ")"

	return out
}
