package _8_fastest_responder

import (
	"net/http"
)

func InternetIsAccessible() bool {
	blogCh := make(chan bool, 1)
	go isAlive("https://www.coreyscott.dev/", blogCh)

	golangCh := make(chan bool, 1)
	go isAlive("https://golang.org/", golangCh)

	var result bool
	select {
	case result = <-blogCh:

	case result = <-golangCh:
	}

	return result
}

func isAlive(url string, responseCh chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		responseCh <- false
		return
	}

	responseCh <- resp.StatusCode == http.StatusOK
}
