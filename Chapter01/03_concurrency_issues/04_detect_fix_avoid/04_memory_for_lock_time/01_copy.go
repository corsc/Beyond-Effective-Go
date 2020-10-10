package main

import (
	"context"
	"sync"
	"time"
)

var (
	cache = []string{}

	cacheMutex = &sync.Mutex{}
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	go updater(ctx)
	go reader(ctx)
}

func reader(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// shutdown
			return

		default:
			// proceed normally
		}

		// copy the cache
		cacheMutex.Lock()
		cacheCopy := make([]string, len(cache))
		copy(cacheCopy, cache)
		cacheMutex.Unlock()

		useTheData(cacheCopy)
	}
}

func updater(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			updates := loadUpdates()

			cacheMutex.Lock()
			cache = updates
			cacheMutex.Unlock()

		case <-ctx.Done():
			// shutdown
			return
		}
	}
}

func useTheData(cacheCopy []string) {
	// not implemented
}

func loadUpdates() []string {
	// not implemented
	return []string{}
}
