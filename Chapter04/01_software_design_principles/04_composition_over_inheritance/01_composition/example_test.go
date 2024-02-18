package _1_composition

import "fmt"

func Example() {
	duck := Duck{}
	parrot := Parrot{}

	duck.Fly()
	parrot.Fly()

	fmt.Println(duck.Talk())

	fmt.Println(parrot.Talk())

	// Output: Quack!
	// Squawk!
}
