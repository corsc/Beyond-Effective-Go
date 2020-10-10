package _3_clean_up

func Example1() {
	go func() {
		for {
			doSomething()
		}
	}()
}

func doSomething() {
	// not implemented
}
