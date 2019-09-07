package version1

type HostConfig interface {
	Host() string
	Port() int
}

type ApplicationConfig struct {
	host HostConfig
}

func (m *ApplicationConfig) GetHostConfig() HostConfig {
	return m.host
}
