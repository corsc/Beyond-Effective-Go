package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	updateCh := make(chan *update)
	go dataUpdater(ctx, wg, updateCh)
	go dataUser(ctx, wg, updateCh)

	wg.Wait()
}

func dataUser(ctx context.Context, wg *sync.WaitGroup, updateCh <-chan *update) {
	defer wg.Done()

	data := map[int64]string{
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

	// simulate a high number of data reads/uses
	readUsage := time.NewTicker(1 * time.Microsecond)

	runs := 0
	for {
		select {
		case <-readUsage.C:
			// lock the data
			runs++
			useTheData(data)

		case thisUpdate := <-updateCh:
			data[thisUpdate.key] = thisUpdate.value

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

func dataUpdater(ctx context.Context, wg *sync.WaitGroup, updateCh chan<- *update) {
	defer wg.Done()

	// simulate an infrequent update
	ticker := time.NewTicker(100 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			updateCh <- &update{
				key:   now.Unix() % 10,
				value: now.Format(time.RFC3339),
			}

		case <-ctx.Done():
			return
		}
	}
}

type update struct {
	key   int64
	value string
}
