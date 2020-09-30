package _1_normal

import (
	"fmt"
)

func ExampleMultiply() {
	result := Multiply(3, 5)

	// Output: 15
	fmt.Println(result)
}

func ExampleMultiply_notFP() {
	inputs := []int{1, 2, 3, 4, 5}

	var result int
	for index, thisValue := range inputs {
		if index == 0 {
			result = thisValue
			continue
		}

		result = Multiply(result, thisValue)
	}

	// Output: 120
	fmt.Println(result)
}
