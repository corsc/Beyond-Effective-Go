package _2_read_from_nil

import (
	"context"
	"time"
)

func extendedExampleFixed(ctx context.Context) {
	var updateCh chan string

	updateTicker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			// shut down
			return

		case <-updateTicker.C:
			// asynchronously perform update
			updateCh = performUpdateAsync()

		case result := <-updateCh:
			// receive result when available
			useResult(result)

			// turn off this select case
			updateCh = nil
		}
	}
}

func performUpdateAsync() chan string {
	result := make(chan string, 1)

	go func() {
		defer close(result)

		// not implemented
	}()

	return result
}
