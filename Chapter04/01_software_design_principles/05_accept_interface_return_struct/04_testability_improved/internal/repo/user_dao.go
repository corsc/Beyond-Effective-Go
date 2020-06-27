package repo

import (
	"database/sql"

	"github.com/corsc/Advanced-Go-Programming/Chapter04/01_software_design_principles/05_accept_interface_return_struct/04_testability_improved/internal/user"
)

type UserDAO struct {
	Database *sql.DB
}

func (u *UserDAO) Save(user *user.User) error {
	// implementation removed
	return nil
}
