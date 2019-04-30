package main

import (
	"time"
)

func main() {
	table := make(chan *ball)

	go newPlayer("paul: ping", table)
	go newPlayer("sally: pong", table)

	// throw a ball at paul
	table <- &ball{}
	table <- &ball{}

	<-time.After(3 * time.Second)
}

func newPlayer(action string, table chan *ball) {
	for thisBall := range table {
		// print action
		println(action)

		// send the ball back
		table <- thisBall
	}
}

// data type to make the example easier to read
type ball struct {
	hits int
}
