package _7_slow_consumers

func Example() {
	signalCh := make(chan struct{}, 1)

	select {
	case signalCh <- struct{}{}:
	// send signal

	default:
		// drop signal as one is already pending
	}

}
