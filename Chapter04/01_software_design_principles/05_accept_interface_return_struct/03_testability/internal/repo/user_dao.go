package repo

import (
	"database/sql"
	"fmt"

	"github.com/corsc/Beyond-Effective-Go/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/user"
)

type UserDAO struct {
	Database *sql.DB
}

func (u *UserDAO) Save(user *user.User) error {
	query := `INSERT INTO user (id,name,email) VALUES (?,?,?)`
	params := []interface{}{user.ID, user.Name, user.Email}

	_, err := u.Database.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("failed to save with err: %s", err)
	}

	return nil
}
