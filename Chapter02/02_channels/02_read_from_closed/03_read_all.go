package _2_read_from_closed

func ExampleReadAllNoBoolean() {
	events := make(chan Event)

	for {
		event := <-events

		processEvent(event)
	}
}

func ExampleReadAllWithBoolean() {
	events := make(chan Event)

	for {
		event, isClosed := <-events
		if isClosed {
			return
		}

		processEvent(event)
	}
}

func ExampleReadAllRange() {
	events := make(chan Event)

	for event := range events {
		processEvent(event)
	}
}

func processEvent(event Event) {
	// implementation omitted
}

type Event struct {
	// implementation omitted
}
