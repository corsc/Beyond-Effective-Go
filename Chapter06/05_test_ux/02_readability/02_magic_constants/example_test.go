package _2_magic_constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleConfusing(t *testing.T) {
	result := NewPerson("Shah", 35, true)

	assert.Equal(t, "Shah", result.Name)
}

func TestExampleClean(t *testing.T) {
	name := "Shah"
	age := 35
	isRightHanded := true

	result := NewPerson(name, age, isRightHanded)

	assert.Equal(t, name, result.Name)
}

func NewPerson(name string, age int, rightHanded bool) *Person {
	return &Person{
		Name:        name,
		Age:         age,
		RightHanded: rightHanded,
	}
}

type Person struct {
	Name        string
	Age         int
	RightHanded bool
}
