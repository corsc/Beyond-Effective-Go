package _8_struct_init

import (
	"fmt"
)

func examples() {
	jane := &Person{
		"Jane",
		"jane@example.com",
		22,
		65,
	}

	mary := &Person{
		Name:   "Mary",
		Email:  "mary@example.com",
		Age:    33,
		Weight: 56,
	}

	// added to make the compiler happy
	fmt.Printf("FU: %#v\n", jane)
	fmt.Printf("FU: %#v\n", mary)
}

type Person struct {
	Name   string
	Email  string
	Age    int
	Weight int
}
