package _1_function

func NewStatsDClient(config *Config) *StatsClient {
	return &StatsClient{
		// implementation removed
	}
}

type StatsClient struct {
	// implementation removed
}

type Config struct {
	host         string
	port         int
	sampleRate   int
	sendBuffer   int
	validateKeys bool
	tags         []string
}
