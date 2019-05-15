package _7_slow_consumers

func Example(signalCh chan struct{}) {
	select {
	case signalCh <- struct{}{}:
		// send signal

	default:
		// drop signal as one is already pending
	}
}
