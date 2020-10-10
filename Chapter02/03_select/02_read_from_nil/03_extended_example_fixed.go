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
			// create channel to receive result
			updateCh = make(chan string, 1)

			// asynchronously perform update
			go performUpdateAsync(updateCh)

		case result := <-updateCh:
			// receive result when available
			useResult(result)

			// turn off this select case
			updateCh = nil
		}
	}
}

func performUpdateAsync(result chan string) {
	// not implemented
}
