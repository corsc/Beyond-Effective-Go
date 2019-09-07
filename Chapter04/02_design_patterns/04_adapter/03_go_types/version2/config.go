package version2

type HostConfig interface {
	Host() string
	Port() int
}

type AppConfig interface {
	GetHostConfig() HostConfig
}

func UsageConfig(hostConfig AppConfig) {
	// implementation removed
}
