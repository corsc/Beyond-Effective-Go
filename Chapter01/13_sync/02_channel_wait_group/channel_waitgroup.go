package _2_channel_wait_group

func Example() bool {
	doneCh := make(chan struct{}, 10)

	// start some goroutines
	for x := 0; x < 10; x++ {
		go func() {
			defer func() {
				// signal done
				doneCh <- struct{}{}
			}()

			// do something
		}()
	}

	// wait for all goroutines to complete
	for x := 0; x < 10; x++ {
		// read done signal
		<-doneCh
	}

	// all goroutines have finished
	return true
}
