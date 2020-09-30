package _7_bool_arguments

import (
	"fmt"
)

func AsInputExample() {
	user := &User{}

	result := user.CheckGender(true)

	fmt.Printf("Result: %#v\n", result)
}

type User struct {
	gender bool
}

func (u *User) CheckGender(isMale bool) bool {
	return u.gender == isMale
}
