package main

func main() {
	userA := &User{
		Name:  "Bob",
		Age:   16,
		Email: "bob@home.com",
	}
	userB := &User{
		Name:  "Jane",
		Age:   23,
		Email: "jane@example.com",
	}

	if userA.Name == userB.Name &&
		userA.Age == userB.Age &&
		userA.Email == userB.Email {
		println("Users A and B are the same!")
	} else {
		println("Users are not the same")
	}
}

type User struct {
	Name  string
	Age   int
	Email string
}
