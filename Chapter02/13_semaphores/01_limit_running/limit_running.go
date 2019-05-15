package _1_limit_running

import (
	"sync"
)

func Example(stopCh chan struct{}, workCh chan Data) {
	wg := &sync.WaitGroup{}

	// create semaphore will max concurrency of 10
	semaphore := make(chan struct{}, 10)

	for {
		select {
		case data := <-workCh:
			wg.Add(1)
			go func() {
				defer wg.Done()

				// acquire semaphore
				semaphore <- struct{}{}

				doWork(data)

				// release semaphore
				<-semaphore
			}()

		case <-stopCh:
			// wait for everything to be done
			wg.Wait()

			return
		}
	}
}

func doWork(data Data) {
	// not implemented
}

type Data struct {
	// some fields
}
