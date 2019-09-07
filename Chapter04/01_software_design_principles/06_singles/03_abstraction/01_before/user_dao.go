package _1_before

import (
	"database/sql"
)

type UserDAO struct {
	db *sql.DB
}

// save the supplied user to the database
func (u *UserDAO) Save(user *User) (int, error) {
	// implementation removed
	return 0, nil
}
