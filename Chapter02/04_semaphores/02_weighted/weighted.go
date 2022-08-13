package _2_limit_pending

import (
	"context"

	"golang.org/x/sync/semaphore"
)

const (
	lightWorkCost = 1
	heavyWorkCost = 4

	maxConcurrency = 10
)

func Example(stopCh chan struct{}, lightWorkCh chan Data, heavyWorkCh chan Data) {
	// create weight semaphore will max concurrency of 10
	weightedSemaphore := semaphore.NewWeighted(maxConcurrency)

	for {
		select {
		case data := <-lightWorkCh:
			go func() {
				// acquire a semaphore
				err := weightedSemaphore.Acquire(context.Background(), lightWorkCost)
				if err != nil {
					return
				}

				doLightWork(data)

				// release semaphore
				weightedSemaphore.Release(lightWorkCost)
			}()

		case data := <-heavyWorkCh:
			go func() {
				// acquire a semaphore
				err := weightedSemaphore.Acquire(context.Background(), heavyWorkCost)
				if err != nil {
					return
				}

				doHeavyWork(data)

				// release semaphore
				weightedSemaphore.Release(heavyWorkCost)
			}()

		case <-stopCh:
			// wait for everything to be done
			_ = weightedSemaphore.Acquire(context.Background(), maxConcurrency)
			return
		}
	}
}

func doLightWork(data Data) {
	// not implemented
}

func doHeavyWork(data Data) {
	// not implemented
}

type Data struct {
	// some fields
}
