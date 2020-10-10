package main

import (
	"bytes"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	maxConcurrent = 5
	totalWork     = 10000000
)

func main() {
	rand.Seed(time.Now().UnixNano())

	wg := &sync.WaitGroup{}

	semaphore := make(chan struct{}, maxConcurrent)

	for x := 0; x < totalWork; x++ {
		wg.Add(1)

		semaphore <- struct{}{}
		go worker(wg, semaphore)
	}

	wg.Wait()
}

func worker(wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	defer func() { <-semaphore }()

	first := rand.Int63()
	second := rand.Int63()

	buffer := &bytes.Buffer{}

	_, _ = buffer.WriteString(strconv.FormatInt(first, 10))
	_, _ = buffer.WriteString("+")
	_, _ = buffer.WriteString(strconv.FormatInt(second, 10))
	_, _ = buffer.WriteString("=")
	_, _ = buffer.WriteString(strconv.FormatInt(first+second, 10))
}
