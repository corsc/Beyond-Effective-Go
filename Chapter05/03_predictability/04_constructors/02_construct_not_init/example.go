package _2_construct_not_init

func NewStorage() *Storage {
	return &Storage{
		cache:  map[string]interface{}{},
		stopCh: make(chan struct{}),
	}
}

type Storage struct {
	cache  map[string]interface{}
	stopCh chan struct{}
}
