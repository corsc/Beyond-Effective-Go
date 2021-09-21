package _2_select_data_flow

func MultiplexingExample(stopCh chan struct{},
	inputChA, inputChB chan int,
	outputChC, outputChD chan int) {

	for {
		var data int
		var isOpen bool

		// read from whichever channel has data
		select {
		case data, isOpen = <-inputChA:
			if !isOpen {
				// channel is done, nil the channel to disable to case
				inputChA = nil
			}

		case data, isOpen = <-inputChB:
			if !isOpen {
				// channel is done, nil the channel to disable to case
				inputChB = nil
			}

		case <-stopCh:
			// give up
			return
		}

		if !isOpen {
			if inputChA == nil && inputChB == nil {
				// both channels are closed
				return
			}

			continue
		}

		// write to whichever channel is empty
		select {
		case outputChC <- data:

		case outputChD <- data:

		case <-stopCh:
			// give up
			return
		}
	}
}
