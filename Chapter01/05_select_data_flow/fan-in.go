package _5_select_data_flow

func FanInExample() {
	stopCh := make(chan struct{})

	inputChA := make(chan int)
	inputChB := make(chan int)
	outputCh := make(chan int)

	for {
		var data int
		var isOpen bool

		// read from whichever channel has data
		select {
		case data = <-inputChA:
			if !isOpen {
				inputChA = nil

				if inputChB == nil {
					// both channels are closed
					return
				}
			}

		case data = <-inputChB:
			if !isOpen {
				// disable this case
				inputChB = nil

				if inputChA == nil {
					// both channels are closed
					return
				}
			}

		case <-stopCh:
			// give up
			return
		}

		// write to output channel
		outputCh <- data
	}
}
