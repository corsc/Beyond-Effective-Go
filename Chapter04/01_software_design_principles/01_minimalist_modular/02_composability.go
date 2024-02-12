package _1_minimalist_modular

import (
	"fmt"
)

func recordUserLogin(user *User) {
	fmt.Printf("UserID %d logged in", user.ID)
}

func recordUserLoginByID(user IDer) {
	fmt.Printf("UserID %d logged in", user.GetID())
}

type IDer interface {
	GetID() int
}

type User struct {
	ID    int
	Name  string
	Email string
}

func (u *User) GetID() int {
	return u.ID
}
