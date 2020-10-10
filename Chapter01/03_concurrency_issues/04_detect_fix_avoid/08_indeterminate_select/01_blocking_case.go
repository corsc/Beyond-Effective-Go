package _8_indeterminate_select

import (
	"context"
	"time"
)

func BlockingCaseExample(ctx context.Context) {
	requestsCh := make(chan request)
	updatesCh := make(chan data)
	updateTicker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-updateTicker.C:
			// this will block if there are no updates available
			newData := <-updatesCh
			processUpdate(newData)

		case request := <-requestsCh:
			processRequest(request)

		case <-ctx.Done():
			// shut down
			return
		}
	}
}

// a single request
type request struct {
	// some fields
}

type data struct {
	// some fields
}

func processUpdate(_ data) {
	// intentional blank
}

func processRequest(_ request) {
	// intentionally blank
}
