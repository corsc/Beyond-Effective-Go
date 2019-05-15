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

	for x := 0; x < maxConcurrent; x++ {
		wg.Add(1)
		go worker(wg, totalWork/maxConcurrent)
	}

	wg.Wait()
}

func worker(wg *sync.WaitGroup, total int) {
	defer wg.Done()

	buffer := &bytes.Buffer{}

	for x := 0; x < total; x++ {
		first := rand.Int63()
		second := rand.Int63()

		_, _ = buffer.WriteString(strconv.FormatInt(first, 10))
		_, _ = buffer.WriteString("+")
		_, _ = buffer.WriteString(strconv.FormatInt(second, 10))
		_, _ = buffer.WriteString("=")
		_, _ = buffer.WriteString(strconv.FormatInt(first+second, 10))

		buffer.Reset()
	}
}
