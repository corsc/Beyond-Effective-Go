package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	sharedData = map[int64]string{
		0: "",
		1: "",
		2: "",
		3: "",
		4: "",
		5: "",
		6: "",
		7: "",
		8: "",
		9: "",
	}
	sharedDataMutex = &sync.RWMutex{}
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go dataUpdater(ctx, wg)
	go dataUser(ctx, wg)

	wg.Wait()
}

func dataUser(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// simulate a high number of data reads/uses
	readUsage := time.NewTicker(1 * time.Microsecond)

	runs := 0

	for {
		select {
		case <-readUsage.C:
			// lock the data
			runs++

			sharedDataMutex.RLock()
			useTheData(sharedData)
			sharedDataMutex.RUnlock()

		case <-ctx.Done():
			// shutdown and report
			readUsage.Stop()

			fmt.Printf("runs: %d\n", runs)
			return
		}
	}
}

func useTheData(in map[int64]string) {
	for key, value := range in {
		fmt.Printf("%d -> %s\n", key, value)
	}
}

func dataUpdater(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	// simulate an infrequent update
	update := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-update.C:
			sharedDataMutex.Lock()

			updateTheData(sharedData)

			sharedDataMutex.Unlock()

		case <-ctx.Done():
			// shutdown
			update.Stop()
			return
		}
	}
}

func updateTheData(in map[int64]string) {
	now := time.Now()
	in[now.Unix()%10] = now.Format(time.RFC3339)
}
