package _1_mutexes

import (
	"sync"
	"testing"
)

func TestChildrenEat(t *testing.T) {
	t.Skip("this is skipped because it contains a data race")

	// make some chopsticks
	chopstickA := NewChopstick()
	chopstickB := NewChopstick()

	// let the children eat
	wg := &sync.WaitGroup{}
	wg.Add(2)

	sally := &GreedyChild{
		name: "Sally",
	}

	bob := &GreedyChild{
		name: "Bob",
	}

	// the chopsticks are intentionally not in the same order to encourage the deadlock
	go sally.Eat(wg, chopstickA, chopstickB)
	go bob.Eat(wg, chopstickB, chopstickA)

	wg.Wait()
}
