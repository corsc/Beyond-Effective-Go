package _1_timeout

import (
	"errors"
	"time"
)

func readWithTimeout(closure func() string, timeout time.Duration) (string, error) {
	resultCh := make(chan string, 1)
	go func() {
		resultCh <- closure()
	}()

	select {
	case result := <-resultCh:
		// happy path, data read result from channel
		return result, nil

	case <-time.After(timeout):
		return "", errors.New("read timed out")
	}
}
