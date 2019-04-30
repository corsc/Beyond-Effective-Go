package _8_indeterminate_select

import (
	"context"
	"time"
)

func AsyncResult(ctx context.Context) {
	requestsCh := make(chan request)
	updateTicker := time.NewTicker(30 * time.Second)

	// note: this is nil
	var updateResultCh chan data

	for {
		select {
		case <-updateTicker.C:
			updateResultCh = make(chan data, 1)
			// this is a potentially slow request, so we are intentionally not blocking the select
			go loadUpdate(updateResultCh)

		case updateResult := <-updateResultCh:
			// when updateResultCh is nil, this case will not be considered
			processUpdate(updateResult)

			// disable this case by setting the channel to nil
			updateResultCh = nil

		case request := <-requestsCh:
			processRequest(request)

		case <-ctx.Done():
			// shut down
			return
		}
	}
}

func loadUpdate(responseCh chan data) {
	// load data from potentially slow source and publish to the provided channel
}
