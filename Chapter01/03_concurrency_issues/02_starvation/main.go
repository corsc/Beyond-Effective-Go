package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// collection of sites and their currency availability
var (
	sites = map[string]int{
		"https://www.coreyscott.dev": http.StatusOK,
		"https://golang.org/":        http.StatusOK,
	}

	// mutex to protect concurrent access to sites
	sitesMutex = &sync.Mutex{}
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go selfishChecker(ctx, wg, "https://www.coreyscott.dev")

	wg.Add(1)
	go politeChecker(ctx, wg, "https://golang.org")

	// wait until all goroutines have shutdown
	wg.Wait()
}

func selfishChecker(ctx context.Context, wg *sync.WaitGroup, url string) {
	defer wg.Done()

	// track how many updates we perform
	totalAttempts := 0

	for {
		totalAttempts++

		select {
		case <-ctx.Done():
			// time to shut down
			fmt.Printf("selish: total updates %d\n", totalAttempts)
			return

		default:
			// continue
		}

		// lock access to sites to prevent concurrent updates
		sitesMutex.Lock()

		resp, err := http.Head(url)
		if err != nil {
			// unlock access to sites
			sitesMutex.Unlock()
			continue
		}

		// update collection of statuses
		sites[url] = resp.StatusCode
		fmt.Printf("%s: %d\n", url, resp.StatusCode)

		// unlock access to sites when done
		sitesMutex.Unlock()
	}
}

func politeChecker(ctx context.Context, wg *sync.WaitGroup, url string) {
	defer wg.Done()

	// track how many updates we perform
	totalAttempts := 0

	for {
		totalAttempts++

		select {
		case <-ctx.Done():
			// time to shut down
			fmt.Printf("polite: total updates %d\n", totalAttempts)
			return

		default:
			// continue
		}

		resp, err := http.Head(url)
		if err != nil {
			continue
		}

		// lock access to sites to prevent concurrent updates
		sitesMutex.Lock()

		// update collection of statuses
		sites[url] = resp.StatusCode
		fmt.Printf("%s: %d\n", url, resp.StatusCode)

		// unlock access to sites when done
		sitesMutex.Unlock()
	}
}
