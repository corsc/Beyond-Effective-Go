package _1_anonymous_closures

import (
	"fmt"
	"sync"
)

func ExampleV3() {
	inputs := []int{1, 2, 3, 4, 5, 6}

	outputs := make([]int, len(inputs))
	outputMutex := &sync.Mutex{}

	wg := &sync.WaitGroup{}
	wg.Add(len(inputs))

	for index, value := range inputs {
		go double(wg, outputs, outputMutex, index, value)
	}

	wg.Wait()

	for index, value := range outputs {
		fmt.Printf("%d -> %d\n", index, value)
	}
}

func double(wg *sync.WaitGroup, outputs []int, outputMutex *sync.Mutex, index int, value int) {
	defer wg.Done()

	outputMutex.Lock()
	outputs[index] = value * 2
	outputMutex.Unlock()
}
