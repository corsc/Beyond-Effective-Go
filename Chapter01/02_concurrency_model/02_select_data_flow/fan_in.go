package _2_select_data_flow

func FanInExample(stopCh chan struct{}, inputChA, inputChB chan int, outputCh chan int) {
	for {
		var data int
		var isOpen bool

		// read from whichever channel has data
		select {
		case data, isOpen = <-inputChA:
			if !isOpen {
				inputChA = nil

				if inputChB == nil {
					// both channels are closed
					return
				}
			}

		case data, isOpen = <-inputChB:
			if !isOpen {
				// disable this case
				inputChB = nil

				if inputChA == nil {
					// both channels are closed
					return
				}
			}

		case <-stopCh:
			// shut down
			return
		}

		// write to output channel
		outputCh <- data
	}
}
