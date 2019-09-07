package _4_testability_improved

import (
	"github.com/PacktPublishing/Advanced-Go-Programming/Chapter04/01_software_design_principles/05_accept_interface_return_struct/04_testability_improved/internal/user"
)

func CreateUser(repository UserRepository, user *user.User) error {
	err := validateUser(user)
	if err != nil {
		return err
	}

	return repository.Save(user)
}

//go:generate mockery -name=UserRepository -case underscore -testonly -inpkg
type UserRepository interface {
	Save(user *user.User) error
}

func validateUser(user *user.User) error {
	// implementation removed
	return nil
}
