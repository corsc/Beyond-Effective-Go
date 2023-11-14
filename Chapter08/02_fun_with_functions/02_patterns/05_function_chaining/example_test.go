package _5_function_chaining

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionChaining(t *testing.T) {
	// Inputs
	data := make(chan *User, 3)
	data <- fred
	data <- john
	data <- paul

	// Close the channel to signal the end of the data
	close(data)

	// Perform transforms and filtering
	results := TransformAndFilter(data, AgeOver20, ProperCaseName, EmailIsValid)

	assert.Equal(t, 1, len(results))
}

var (
	fred = &User{
		ID:    1,
		Name:  "Fred",
		Age:   10,
		Email: "fred@flintstones.com",
	}

	john = &User{
		ID:    2,
		Name:  "john",
		Age:   35,
		Email: "john@thebeatles.com",
	}

	paul = &User{
		ID:    3,
		Name:  "Paul",
		Age:   81,
		Email: "",
	}
)
