package _1_style

import (
	"fmt"
)

func NamingExample() {
	// vanilla
	r := do()
	fmt.Printf("result: %#v\n", r)

	// longer names
	result := do()
	fmt.Printf("result: %#v\n", result)
}

func do() string {
	// implementation removed
	return ""
}
