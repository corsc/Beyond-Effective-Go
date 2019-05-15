package livelock

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestChildrenEat(t *testing.T) {
	// ensure the test does not run forever
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// make some chopsticks
	chopstickA := NewChopstick()
	chopstickB := NewChopstick()

	// let the children eat
	wg := &sync.WaitGroup{}
	wg.Add(2)

	sally := &PoliteChild{
		name: "Sally",
	}

	bob := &PoliteChild{
		name: "Bob",
	}

	sally.sibling = bob
	bob.sibling = sally

	// the chopsticks are intentionally not in the same order to encourage the livelock
	go sally.Eat(ctx, wg, chopstickA, chopstickB)
	go bob.Eat(ctx, wg, chopstickB, chopstickA)

	wg.Wait()
}
