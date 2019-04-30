package _2_separate_logic

func UpdateDestinationV2() {
	sourceMutex.Lock()
	results := doCalculation(sourceMap)
	sourceMutex.Unlock()

	// release locks
	destinationMutex.Lock()
	destinationMap = results
	destinationMutex.Unlock()
}
