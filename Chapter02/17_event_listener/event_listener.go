package _7_event_listener

type Clock struct {
	listeners []chan int64
}

func (c *Clock) AddListener(in chan int64) {
	c.listeners = append(c.listeners, in)
}

func (c *Clock) onTick(ts int64) {
	for _, listener := range c.listeners {
		// non-blocking write
		select {
		case listener <- ts:

		default:
		}
	}
}

func (c *Clock) RemoveListener(in chan int64) {
	foundIndex := -1

	for index, listener := range c.listeners {
		if listener == in {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		return
	}

	// remove and zero the requested item
	copy(c.listeners[foundIndex:], c.listeners[foundIndex+1:])
	c.listeners[len(c.listeners)-1] = nil
	c.listeners = c.listeners[:len(c.listeners)-1]
}
