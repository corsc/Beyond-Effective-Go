package _2_separate_logic

import (
	"sync"
)

var (
	sourceMap   = map[string]string{}
	sourceMutex = &sync.Mutex{}

	destinationMap   = map[string]string{}
	destinationMutex = &sync.Mutex{}
)

func UpdateDestinationV1() {
	// ensure we are the only ones accessing the data
	sourceMutex.Lock()
	destinationMutex.Lock()

	// perform calculation and update destination
	destinationMap = doCalculation(sourceMap)

	// release locks
	destinationMutex.Unlock()
	sourceMutex.Unlock()
}

func doCalculation(in map[string]string) map[string]string {
	// intentionally empty
	return in
}
