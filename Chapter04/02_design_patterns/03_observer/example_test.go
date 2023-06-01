package _3_observer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestObserverPattern(t *testing.T) {
	// Create a celebrity and fan
	kentBeck := &Celebrity{}
	corey := NewSuperFan()

	// start following the celebrity
	kentBeck.Follow(corey.eventCh)
	assert.Equal(t, 1, len(kentBeck.fans))

	// wait for new content
	go func() {
		corey.Watch()
	}()

	kentBeck.Post(Post{Content: "Hello World!"})

	// wait then unfollow
	<-time.After(1 * time.Second)

	kentBeck.Unfollow(corey.eventCh)
	assert.Equal(t, 0, len(kentBeck.fans))
}
