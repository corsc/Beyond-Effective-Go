package _1_leak_dependency

import (
	"github.com/corsc/Beyond-Effective-Go/Chapter05/03_predictability/03_encapsulation/01_leak_dependency/internal/storage"
)

func Authenticate(username, password string) (*storage.User, error) {
	err := performBusinessLogic(username, password)
	if err != nil {
		return nil, err
	}

	return storage.LoadByUsernamePassword(username, password)
}

func performBusinessLogic(username string, password string) error {
	// implementation removed
	return nil
}
