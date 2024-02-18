package _5_implicit_interfaces

import (
	"fmt"
)

func Example() {
	var userDaoA UserDAO = NewUserDAOInterface()
	var userDaoB UserDAO = NewUserDAOStruct()

	// Use values to avoid compiler complaints
	fmt.Printf("A: %T\n", userDaoA)
	fmt.Printf("B: %T\n", userDaoB)

	// Output: A: *_5_implicit_interfaces.UserDAOImpl
	// B: *_5_implicit_interfaces.UserDAOImpl
}
