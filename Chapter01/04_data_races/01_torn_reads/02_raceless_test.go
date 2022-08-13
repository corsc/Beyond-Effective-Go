package _1_torn_reads

import (
	"fmt"
	"sync"
	"testing"
)

func TestDataRaceLess(t *testing.T) {
	message := make([]int, 20)
	messageMutex := &sync.RWMutex{}

	wg := &sync.WaitGroup{}

	for x := 0; x < 100; x++ {
		wg.Add(2)

		go func(val int) {
			defer wg.Done()

			// generate and replace the messages
			messageMutex.Lock()
			for i := 0; i < len(message); i++ {
				message[i] = val
			}

			messageMutex.Unlock()
		}(x)

		go func() {
			defer wg.Done()

			messageMutex.RLock()
			fmt.Printf("msg: %#v\n", message)
			messageMutex.RUnlock()
		}()
	}

	wg.Wait()
}
