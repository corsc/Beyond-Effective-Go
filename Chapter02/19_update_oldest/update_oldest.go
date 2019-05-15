package _9_update_oldest

import (
	"sync"
	"time"
)

type SiteMonitor struct {
	statuses sync.Map

	sites chan string
}

func (s *SiteMonitor) updater(stopCh chan struct{}) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			// pop oldest
			url := <-s.sites

			// use it
			result := update(url)

			s.statuses.Store(url, result)

			// push back onto the channel
			s.sites <- url

		case <-stopCh:
			return
		}
	}
}

func update(url string) int {
	// not implemented
	return 0
}
