package _5_function_chaining

import (
	"net/mail"
	"strings"
)

type User struct {
	ID    int64
	Name  string
	Age   int64
	Email string
}

type ChainedFunction func(in <-chan *User) <-chan *User

func TransformAndFilter(in <-chan *User, chainedFunctions ...ChainedFunction) []*User {
	for _, chainedFunc := range chainedFunctions {
		in = chainedFunc(in)
	}

	var results []*User

	for user := range in {
		results = append(results, user)
	}

	return results
}

func AgeOver20(in <-chan *User) <-chan *User {
	out := make(chan *User)

	go func() {
		defer close(out)

		for user := range in {
			if user.Age > 20 {
				out <- user
			}
		}
	}()

	return out
}

func ProperCaseName(in <-chan *User) <-chan *User {
	out := make(chan *User)

	go func() {
		defer close(out)

		for user := range in {
			user.Name = strings.ToTitle(user.Name)

			out <- user
		}
	}()

	return out
}

func EmailIsValid(in <-chan *User) <-chan *User {
	out := make(chan *User)

	go func() {
		defer close(out)

		for user := range in {
			_, err := mail.ParseAddress(user.Email)
			if err != nil {
				continue
			}

			out <- user
		}
	}()

	return out
}
