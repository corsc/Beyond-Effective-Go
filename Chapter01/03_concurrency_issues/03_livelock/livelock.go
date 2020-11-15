package livelock

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type PoliteChild struct {
	name    string
	sibling *PoliteChild
	full    bool
}

func (p *PoliteChild) Eat(ctx context.Context, wg *sync.WaitGroup, chopstick1, chopstick2 *chopstick) {
	defer wg.Done()

	for {
		// give up if context is done
		select {
		case <-ctx.Done():
			fmt.Printf("%s: gave up\n", p.name)
			return

		default:
			// proceed
		}

		p.pickup(chopstick1)

		// attempt to avoid deadlock by checking the other process
		// that might hold the mutex
		if p.sibling.IsHungry() {
			// attempt to yield and let the other process proceed
			p.putDown(chopstick1)

			<-time.After(100 * time.Millisecond)
			continue
		}

		p.pickup(chopstick2)

		p.eatUntilFull()

		chopstick1.PutDown()
		chopstick2.PutDown()
		return
	}
}

func (p *PoliteChild) eatUntilFull() {
	p.full = true
}

func (p *PoliteChild) IsHungry() bool {
	return !p.full
}

func (p *PoliteChild) pickup(chopstick *chopstick) {
	fmt.Printf("%s: pick up\n", p.name)

	chopstick.Pickup()
}

func (p *PoliteChild) putDown(chopstick *chopstick) {
	fmt.Printf("%s: put down\n", p.name)

	chopstick.PutDown()
}

func NewChopstick() *chopstick {
	return &chopstick{
		mutex: &sync.Mutex{},
	}
}

type chopstick struct {
	mutex *sync.Mutex
}

func (c *chopstick) Pickup() {
	c.mutex.Lock()
}

func (c *chopstick) PutDown() {
	c.mutex.Unlock()
}
