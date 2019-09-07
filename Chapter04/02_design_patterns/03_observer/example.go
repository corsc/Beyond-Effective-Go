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
	mutex *sync.RWMutex
}

// Subscribe/Add to list of Observers
func (o *Celebrity) Follow(responseCh chan Post) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	o.fans = append(o.fans, responseCh)
}

// Unsubscribe/Remove from the list of Observers
func (o *Celebrity) Unfollow(in chan Post) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	for index, observer := range o.fans {
		if observer == in {
			o.fans = append(o.fans[:index], o.fans[index+1:]...)
			return
		}
	}
}

// Notify all Observers
func (o *Celebrity) Post(post Post) {
	o.mutex.RLock()
	defer o.mutex.RUnlock()

	for _, fan := range o.fans {
		// optional write so that Observers cannot block this method
		select {
		case fan <- post:
		default:
		}
	}
}

type SuperFan struct {
	eventCh chan Post
}

func (s *SuperFan) Watch() {
	for post := range s.eventCh {
		fmt.Printf("Celebrity has posted %s", post.Content)
	}
}

func NewSuperFan() *SuperFan {
	return &SuperFan{
		// use buffered channel to avoid blocking and data loss
		eventCh: make(chan Post, 1),
	}
}
