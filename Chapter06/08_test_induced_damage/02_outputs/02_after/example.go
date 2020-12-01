package _2_after

func Example() {
	DoSomething()
}

func DoSomething() {
	_ = doSomething()
}

func doSomething() error {
	// implementation removed
	return nil
}
