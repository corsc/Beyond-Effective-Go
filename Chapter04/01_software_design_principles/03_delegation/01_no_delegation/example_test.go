package main

import "fmt"

func Example_noDelegation() {
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

	if userA.Name == userB.Name &&
		userA.Age == userB.Age &&
		userA.Email == userB.Email {
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
