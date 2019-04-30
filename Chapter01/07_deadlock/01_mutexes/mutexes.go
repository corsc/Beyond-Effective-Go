package _1_mutexes

import (
	"runtime"
	"sync"
)

type GreedyChild struct {
	name string
}

func (g *GreedyChild) Eat(wg *sync.WaitGroup, chopsticks ...*chopstick) {
	defer wg.Done()

	for _, c := range chopsticks {
		c.Pickup()

		// interrupt scheduler (forcing the appearance of deadlock)
		runtime.Gosched()
	}

	g.eatUntilFull()

	for _, c := range chopsticks {
		c.PutDown()
	}
}

func (g *GreedyChild) eatUntilFull() {
	// intentionally blank
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
