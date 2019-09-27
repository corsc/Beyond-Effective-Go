package _2_state

import (
	"fmt"
)

func Example_buildTotaler() {
	totaler := buildTotaler()

	_ = totaler(1)
	_ = totaler(2)
	_ = totaler(3)
	_ = totaler(4)
	result := totaler(5)

	// Output: 15
	fmt.Printf("%d", result)
}
