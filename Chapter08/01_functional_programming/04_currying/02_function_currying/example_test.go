package _2_function_currying

import (
	"fmt"
)

func ExampleMultiply_ShortForm() {
	result := Multiply(5)(3)

	// Output: 15
	fmt.Println(result)
}

func ExampleMultiply_LongForm() {
	multiply5 := Multiply(5)

	result := multiply5(3)

	// Output: 15
	fmt.Println(result)
}
