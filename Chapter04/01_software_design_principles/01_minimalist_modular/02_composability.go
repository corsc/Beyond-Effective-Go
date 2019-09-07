package _1_minimalist_modular

import (
	"fmt"
)

func recordUserLogin(user *User) {
	fmt.Printf("UserID %d logged in", user.ID)
}

func recordUserLoginByID(user UserID) {
	fmt.Printf("UserID %d logged in", user.UserID())
}

type UserID interface {
	UserID() int
}

type User struct {
	ID    int
	Name  string
	Email string
}
