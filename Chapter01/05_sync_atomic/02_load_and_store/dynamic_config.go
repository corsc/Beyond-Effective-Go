package _2_load_and_store

import (
	"sync/atomic"
	"time"
)

// a collection of configuration values (stored in a database or key/value store)
type Config struct {
	FeatureFlagA bool
	FeatureFlagB bool
	FeatureFlagC bool
}

type ConfigManager struct {
	data   atomic.Value
	stopCh chan struct{}
}

// wrapper around the data to make it easier to use
func NewConfigManager() *ConfigManager {
	out := &ConfigManager{
		stopCh: make(chan struct{}),
	}

	// store the default (empty data)
	out.data.Store(&Config{})

	// periodically refresh the data from the data store
	go out.monitorConfig()

	return out
}

func (c *ConfigManager) monitorConfig() {
	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-ticker.C:
			cfg := loadConfigFromDatastore()
			c.data.Store(cfg)

		case <-c.stopCh:
			// shut down
			return
		}
	}
}

func (c *ConfigManager) ShutDown() {
	close(c.stopCh)
}

func (c *ConfigManager) Config() *Config {
	// return a copy of the config
	return c.data.Load().(*Config)
}

func loadConfigFromDatastore() *Config {
	// not implemented
	return &Config{}
}
