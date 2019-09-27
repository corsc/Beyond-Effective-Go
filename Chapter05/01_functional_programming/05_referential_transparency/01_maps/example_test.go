package _1_maps

func ExampleToLower() {
	in := map[string]string{"A": "B"}
	ToLower(in)
	printAll(in)

	// Output: A: b
}
