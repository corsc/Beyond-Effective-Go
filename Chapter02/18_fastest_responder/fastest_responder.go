package _8_fastest_responder

import (
	"net/http"
)

func InternetIsAccessible() bool {
	googleCh := make(chan bool, 1)
	go isAlive("https://www.google.com/", googleCh)

	packtCh := make(chan bool, 1)
	go isAlive("https://www.packtpub.com/", packtCh)

	var result bool
	select {
	case result = <-googleCh:

	case result = <-packtCh:
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
