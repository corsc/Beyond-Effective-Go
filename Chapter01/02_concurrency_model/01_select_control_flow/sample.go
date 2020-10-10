package _1_select_control_flow

func Example() {
	dataCh := make(chan Data)
	stopCh := make(chan struct{})

	for {
		select {
		case data := <-dataCh:
			processData(data)

		case <-stopCh:
			// stop processing data
			return
		}
	}
}

func processData(data Data) {
	// intentionally blank
}

type Data struct {
	// some fields
}
