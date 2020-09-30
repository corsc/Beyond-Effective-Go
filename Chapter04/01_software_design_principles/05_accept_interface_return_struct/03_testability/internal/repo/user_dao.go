package repo

import (
	"database/sql"

	"github.com/corsc/Beyond-Effective-Go/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/user"
)

type UserDAO struct {
	Database *sql.DB
}

func (u *UserDAO) Save(user *user.User) error {
	// implementation removed
	return nil
}
