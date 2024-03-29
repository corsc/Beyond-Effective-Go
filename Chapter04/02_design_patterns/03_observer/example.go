package _3_observer

import (
	"fmt"
	"sync"
)

type Post struct {
	Content string
}

// This is the Subject
type Celebrity struct {
	fans  []chan Post
	mutex sync.RWMutex
}

// Subscribe/add to list of observers
func (c *Celebrity) Follow(responseCh chan Post) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.fans = append(c.fans, responseCh)
}

// Unsubscribe/Remove from the list of observers
func (c *Celebrity) Unfollow(responseCh chan Post) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for index, observer := range c.fans {
		if observer == responseCh {
			c.fans = append(c.fans[:index], c.fans[index+1:]...)

			// close the channel and stop the watch loop
			close(responseCh)

			return
		}
	}
}

// Notify all observers
func (c *Celebrity) Upload(post Post) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for _, fan := range c.fans {
		// optional write so that observers cannot block this method
		select {
		case fan <- post:
			// successfully notified observer

		default:
			// skip observer that isn't ready
		}
	}
}

type SuperFan struct {
	eventCh chan Post
}

func (s *SuperFan) Watch() {
	for post := range s.eventCh {
		fmt.Printf("Celebrity has posted: %s", post.Content)
	}
}

func NewSuperFan() *SuperFan {
	return &SuperFan{
		// use buffered channel to avoid blocking and data loss
		eventCh: make(chan Post, 1),
	}
}
