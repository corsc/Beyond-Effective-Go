package main

import (
	"fmt"
	"sync"
	"time"
)

var inOnline = true
var isOnlineMutex = &sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}

	// define a condition variable that monitors the server state
	serverStatus := sync.NewCond(isOnlineMutex)

	// start some goroutines that use the server when it is online
	wg.Add(1)
	go connection("A", serverStatus, wg)

	wg.Add(1)
	go connection("B", serverStatus, wg)

	// wait
	<-time.After(1 * time.Second)

	// take the server offline
	serverStatus.L.Lock()
	inOnline = false
	serverStatus.L.Unlock()

	// wait
	<-time.After(1 * time.Second)

	// signal the server is online
	serverStatus.L.Lock()
	inOnline = true
	serverStatus.L.Unlock()

	serverStatus.Broadcast()

	// run until complete
	wg.Wait()
}

func connection(name string, serverStatus *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()

	// read message from caller
	for x := 0; x < 30; x++ {
		// check server is alive
		serverStatus.L.Lock()
		for !inOnline {
			fmt.Printf("%s: blocked\n", name)
			serverStatus.Wait()
		}
		serverStatus.L.Unlock()

		sendToServer(name, x)
	}
}

func sendToServer(name string, index int) {
	fmt.Printf("%s: %d\n", name, index)

	// small delay to make the output nicer
	<-time.After(100 * time.Millisecond)
}
