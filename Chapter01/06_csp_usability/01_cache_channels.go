package _6_csp_usability

func NewCacheUsingChannels() *CacheUsingChannels {
	// create a cache
	out := &CacheUsingChannels{
		cache:      make(map[string]Person),
		operations: make(chan cacheOperation),
		shutDown:   make(chan struct{}),
	}

	// start worker to process cache operations
	go out.cacheWorker()

	return out
}

type CacheUsingChannels struct {
	cache map[string]Person

	operations chan cacheOperation
	shutDown   chan struct{}
}

func (c *CacheUsingChannels) Get(key string) <-chan Person {
	out := make(chan Person)

	if c.isShutdown() {
		close(out)
		return out
	}

	// schedule a get operation
	c.operations <- func() {
		out <- c.cache[key]
	}

	return out
}

func (c *CacheUsingChannels) Set(key string, value Person) {
	if c.isShutdown() {
		return
	}

	// schedule a set operation
	c.operations <- func() {
		c.cache[key] = value
	}
}

func (c *CacheUsingChannels) cacheWorker() {
	for cacheOperation := range c.operations {
		cacheOperation()
	}
}

func (c *CacheUsingChannels) isShutdown() bool {
	select {
	case <-c.shutDown:
		return true

	default:
		return false
	}
}

func (c *CacheUsingChannels) shutdown() {
	close(c.shutDown)
	close(c.operations)
}

type Person struct {
	Name string
}

type cacheOperation func()
