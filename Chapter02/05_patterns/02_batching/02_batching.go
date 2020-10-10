package _2_batching

const batchSize = 20

type BatchedSender struct {
	buffer   [batchSize]Item
	position int
}

func (b *BatchedSender) Send(dataCh chan<- [batchSize]Item, item Item) {
	b.buffer[b.position] = item

	b.position++

	if b.position >= batchSize {
		dataCh <- b.buffer
		b.position = 0
	}
}

type Item struct {
	ID int
}
