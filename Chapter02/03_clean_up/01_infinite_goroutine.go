package _3_clean_up

func Example01() {
	go func() {
		for {
			doSomething()
		}
	}()
}

func doSomething() {
	// not implemented
}
