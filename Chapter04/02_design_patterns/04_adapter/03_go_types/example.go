package _3_go_types

import (
	"github.com/corsc/Beyond-Effective-Go/Chapter04/02_design_patterns/04_adapter/03_go_types/version1"
	"github.com/corsc/Beyond-Effective-Go/Chapter04/02_design_patterns/04_adapter/03_go_types/version2"
)

func Usage() {
	// Without the adapter this code throws
	//	# github.com/corsc/Beyond-Effective-Go/Chapter04/02_design_patterns/04_adapter/03_go_types
	//	Chapter04/02_design_patterns/04_adapter/03_go_types/example.go:11:22: cannot use configVersion1 (type *version1.AppConfig) as type version2.AppConfig in argument to version2.UsageConfig:
	//	*version1.AppConfig does not implement version2.AppConfig (wrong type for GetHostConfig method)
	//	have GetHostConfig() version1.HostConfig
	//	want GetHostConfig() version2.HostConfig
	//
	//configVersion1 := &version1.AppConfig{}
	//version2.UsageConfig(configVersion1)

	configVersion1 := &version1.ApplicationConfig{}
	adaptedConfig := &configAdapter{
		configVersion1: configVersion1,
	}

	version2.UsageConfig(adaptedConfig)
}

type configAdapter struct {
	configVersion1 *version1.ApplicationConfig
}

func (c *configAdapter) GetHostConfig() version2.HostConfig {
	return c.configVersion1.GetHostConfig()
}
