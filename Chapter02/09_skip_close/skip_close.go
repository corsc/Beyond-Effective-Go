package _9_skip_close

func Example(stopCh chan struct{}) {
	semaphore := make(chan struct{})

	for {
		select {
		case semaphore <- struct{}{}:
			go func() {
				doWork()

				<-semaphore
			}()

		case <-stopCh:
			// shut down
			return
		}
	}
}

func doWork() {
	// not implemented
}
