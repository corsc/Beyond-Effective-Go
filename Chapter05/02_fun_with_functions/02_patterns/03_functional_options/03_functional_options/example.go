package _3_functional_options

const (
	defaultPort = 8125
)

func NewStatsDClient(host string, options ...Option) *StatsClient {
	client := &StatsClient{
		host: host,
	}

	// apply options
	for _, option := range options {
		option(client)
	}

	return client
}

type StatsClient struct {
	host         string
	port         int
	rate         int
	sendBuffer   int
	validateKeys bool
	tags         []string
}

type Option func(*StatsClient)

func Port(port int) Option {
	return func(client *StatsClient) {
		// validate and set port
		if port <= 0 {
			port = defaultPort
		}
		client.port = port
	}
}

func SampleRate(rate int) Option {
	return func(client *StatsClient) {
		// implementation removed
	}
}

func SendBuffer(size int) Option {
	return func(client *StatsClient) {
		// implementation removed
	}
}

func ValidateKeys(validate bool) Option {
	return func(client *StatsClient) {
		// implementation removed
	}
}

func Tags(tags []string) Option {
	return func(client *StatsClient) {
		// implementation removed
	}
}
