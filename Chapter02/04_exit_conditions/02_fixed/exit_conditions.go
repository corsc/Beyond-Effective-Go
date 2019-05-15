package _1_broken

import (
	"context"
	"database/sql"
	"sync"
)

var db *sql.DB

var cache = map[string]Person{}
var cacheMutex = &sync.Mutex{}

func RefreshPeopleCache(ctx context.Context, names []string) error {
	errorCh := make(chan error, 1)

	wg := &sync.WaitGroup{}

	// limit concurrency to 3
	semaphoreCh := make(chan struct{}, 3)

	for _, name := range names {
		wg.Add(1)
		go loadRecord(ctx, wg, semaphoreCh, errorCh, name)
	}

	wg.Wait()

	close(errorCh)

	return <-errorCh
}

func loadRecord(ctx context.Context, wg *sync.WaitGroup, semaphoreCh chan struct{}, errorCh chan error, name string) {
	defer func() {
		// release semaphore
		<-semaphoreCh

		wg.Done()
	}()

	// acquire semaphore
	semaphoreCh <- struct{}{}

	result := db.QueryRowContext(ctx, "SELECT user, host FROM user WHERE user = ?", name)

	person := &Person{}
	err := result.Scan(person.Name, person.Host)
	if err != nil {
		select {
		case errorCh <- err:
			// record the error

		default:
			// leak/drop the error
		}
		return
	}

	cacheMutex.Lock()
	cache[name] = *person
	cacheMutex.Unlock()
}

type Person struct {
	Name string
	Host string
}
