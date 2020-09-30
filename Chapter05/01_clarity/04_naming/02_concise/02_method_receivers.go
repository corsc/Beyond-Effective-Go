package _2_concise

type Order struct {
	items []Item
}

func (o *Order) AddItem(item Item) {
	o.items = append(o.items, item)
}

type Item struct {
	// implementation removed
}
