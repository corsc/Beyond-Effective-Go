package _3_testability

import (
	"github.com/corsc/Beyond-Effective-Go/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/repo"
	"github.com/corsc/Beyond-Effective-Go/Chapter04/01_software_design_principles/05_accept_interface_return_struct/03_testability/internal/user"
)

func CreateUser(repository *repo.UserDAO, user *user.User) error {
	err := validateUser(user)
	if err != nil {
		return err
	}

	return repository.Save(user)
}

func validateUser(user *user.User) error {
	// implementation removed
	return nil
}
