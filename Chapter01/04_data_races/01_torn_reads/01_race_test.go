package _1_torn_reads

import (
	"fmt"
	"sync"
	"testing"
)

func TestDataRace(t *testing.T) {
	message := make([]int, 20)

	wg := &sync.WaitGroup{}

	for x := 0; x < 100; x++ {
		wg.Add(2)

		go func(val int) {
			defer wg.Done()

			// generate replace message
			for i := 0; i < len(message); i++ {
				message[i] = val
			}
		}(x)

		go func() {
			defer wg.Done()
			fmt.Printf("msg: %#v\n", message)
		}()
	}

	wg.Wait()
}
