package _2_defer

import (
	"sync"
)

var (
	mutex   = &sync.Mutex{}
	counter int
)

func LockWithDefer() {
	mutex.Lock()
	defer mutex.Unlock()

	doSomething(counter)
	doSomethingElse()
}

func LockNoDefer(mutex *sync.Mutex) {
	mutex.Lock()
	doSomething(counter)
	mutex.Unlock()

	doSomethingElse()
}

func UseCopy(mutex *sync.Mutex) {
	mutex.Lock()
	copyOfCounter := counter
	mutex.Unlock()

	doSomething(copyOfCounter)
	doSomethingElse()
}

func doSomething(val int) {
	// not implemented
}

func doSomethingElse() {
	// not implemented
}
