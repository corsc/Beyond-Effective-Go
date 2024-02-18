package _5_implicit_interfaces

import (
	"errors"
)

type UserDAO interface {
	Load(ID int) (*User, error)
	Save(user *User) error
}

func NewUserDAOInterface() UserDAO {
	return &UserDAOImpl{}
}

func NewUserDAOStruct() *UserDAOImpl {
	return &UserDAOImpl{}
}

type UserDAOImpl struct{}

func (u *UserDAOImpl) Load(ID int) (*User, error) {
	// implementation removed
	return nil, errors.New("not implemented")
}

func (u *UserDAOImpl) Save(user *User) error {
	// implementation removed
	return errors.New("not implemented")
}

type User struct {
	ID   int
	Name string
}
