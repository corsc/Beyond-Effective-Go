package _1_greetings

import (
	"fmt"
)

func Example_buildGreeting() {
	greetSophia := buildGreeting("Sophia")

	fmt.Println(greetSophia())
	fmt.Println(greetSophia())

	// Output: Hello Sophia. Nice to meet you!
	// Hello Sophia. Nice to meet you!
}
