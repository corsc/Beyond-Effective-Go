package _3_clean_up

func Example2() {
	stopCh := make(chan struct{})
	defer close(stopCh)

	go func() {
		for {
			doSomething()

			// check for stop signal
			select {
			case <-stopCh:
				// stop the goroutine
				return

			default:
				// do another loop
			}
		}
	}()
}
