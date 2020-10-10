package _1_anonymous_closures

import (
	"fmt"
	"sync"
)

func Example() {
	inputs := []int{1, 2, 3, 4, 5, 6}

	outputs := make([]int, len(inputs))

	wg := &sync.WaitGroup{}
	wg.Add(len(inputs))

	for index, value := range inputs {
		go func() {
			defer wg.Done()

			outputs[index] = value * 2
		}()
	}

	wg.Wait()

	for index, value := range outputs {
		fmt.Printf("%d -> %d\n", index, value)
	}
}
