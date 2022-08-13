package _1_select_control_flow

func SelectExample(dataCh chan Data, stopCh chan struct{}) {
	defer close(dataCh)

	for {
		select {
		case data := <-dataCh:
			processData(data)

		case <-stopCh:
			// stop processing the data
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
