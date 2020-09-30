package _3_whitespace

import (
	"fmt"
)

func speakWhitespace(animal string) {
	switch animal {
	case "dog":
		fmt.Print("Woof!")
		fmt.Print("Woof!")

	case "cat":
		fmt.Print("Meow")

	case "mouse":
		fmt.Print("Squeak")
		fmt.Print("Squeak")

	default:
		fmt.Print("???")
	}
}
