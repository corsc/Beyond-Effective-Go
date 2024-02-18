package _1_meaningful

func Example_variableName() {
	// Bad
	_ = get()

	// Better
	_ = userDecision()

	// Best
	_ = getUserDecision()
}

func get() bool {
	return false
}

func userDecision() bool {
	return false
}

func getUserDecision() bool {
	return false
}
