package main

import "fmt"

func Example_delegation() {
	userA := &User{
		Name:  "Bob",
		Age:   16,
		Email: "bob@example.com",
	}
	userB := &User{
		Name:  "Jane",
		Age:   23,
		Email: "jane@example.com",
	}

	if userA.Equals(userB) {
		fmt.Println("Users A and B are the same!")
	} else {
		fmt.Println("Users are not the same")
	}

	// Output: Users are not the same
}

type User struct {
	Name  string
	Age   int
	Email string
}

func (u *User) Equals(in *User) bool {
	return u.Name == in.Name &&
		u.Age == in.Age &&
		u.Email == in.Email
}
