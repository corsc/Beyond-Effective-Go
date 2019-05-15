package _1_anonymous_closures

import (
	"fmt"
	"sync"
)

func ExampleV2() {
	inputs := []int{1, 2, 3, 4, 5, 6}

	outputs := make([]int, len(inputs))
	outputMutex := &sync.Mutex{}

	wg := &sync.WaitGroup{}
	wg.Add(len(inputs))

	for index, value := range inputs { // line 17
		go func() {
			defer wg.Done()

			outputMutex.Lock()
			outputs[index] = value * 2 // line 22
			outputMutex.Unlock()
		}()
	}

	wg.Wait()

	for index, value := range outputs {
		fmt.Printf("%d -> %d\n", index, value)
	}
}
