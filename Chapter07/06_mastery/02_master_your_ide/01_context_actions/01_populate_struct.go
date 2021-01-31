package _1_context_actions

import "fmt"

func Example() {
	var user *User

	user = &User{
		Name:  "",
		Age:   0,
		Email: "",
	}

	fmt.Println(user)
}

type User struct {
	Name  string
	Age   int
	Email string
}
