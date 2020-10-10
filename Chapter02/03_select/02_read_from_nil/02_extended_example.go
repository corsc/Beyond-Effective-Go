package _2_read_from_nil

import (
	"context"
	"time"
)

func extendedExample(ctx context.Context) {
	updateTicker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			// shut down
			return

		case <-updateTicker.C:
			result := performUpdate()
			useResult(result)
		}
	}
}

func performUpdate() string {
	// not implemented
	return ""
}

func useResult(in string) {
	// not implemented
}
