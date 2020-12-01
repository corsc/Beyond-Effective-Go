package _3_latches

type Consumer struct{}

func (c *Consumer) Consume(ordersCh chan *Order) {
	for order := range ordersCh {
		if !c.isValid(order) {
			continue
		}

		c.process(order)
	}
}

func (c *Consumer) isValid(order *Order) bool {
	// implementation removed
	return true
}

func (c *Consumer) process(order *Order) {
	// implementation remove
}

type Order struct {
	// fields removed
}
