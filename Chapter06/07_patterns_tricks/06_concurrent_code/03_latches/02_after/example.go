package _3_latches

func NewConsumer() *Consumer {
	return &Consumer{
		orderDropped: func(order *Order) {},
	}
}

type Consumer struct {
	orderDropped func(*Order)
}

func (c *Consumer) Consume(ordersCh chan *Order) {
	for order := range ordersCh {
		if !c.isValid(order) {
			c.orderDropped(order)

			continue
		}

		c.process(order)
	}
}

func (c *Consumer) isValid(order *Order) bool {
	// implementation removed
	return order != nil
}

func (c *Consumer) process(order *Order) {
	// implementation remove
}

type Order struct {
	// fields removed
}
