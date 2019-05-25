package _5_select_data_flow

func FanOutExample(inputCh chan int, outputChA, outputChB chan int) {
	for data := range inputCh {
		// write to whichever channel is empty
		select {
		case outputChA <- data:

		case outputChB <- data:

		}
	}
}
