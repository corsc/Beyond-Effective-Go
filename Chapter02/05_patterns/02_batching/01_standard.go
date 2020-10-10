package _2_batching

func ReadFromChannel(dataCh chan Item) {
	for item := range dataCh {
		go doWork(item)
	}
}

func doWork(item Item) {
	// not implemented
}
