package example_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/corsc/Beyond-Effective-Go/Chapter05/02_consistency/08_constructor_params/02_config_injection"
	"github.com/stretchr/testify/require"
)

// this only exists so that example can be written as a function
func TestExample(t *testing.T) {
	// these are not included in the example
	var storage example.Storage
	var risk example.RiskManager

	require.NoError(t, configInjectionExample(storage, risk))
}

func configInjectionExample(storage example.Storage, risk example.RiskManager) error {
	// load the config
	rawJSON, err := os.ReadFile("testdata/config.json")
	if err != nil {
		return fmt.Errorf("failed to load config with err: %w", err)
	}

	config := &AppConfig{}
	err = json.Unmarshal(rawJSON, config)
	if err != nil {
		return fmt.Errorf("failed to load config with err: %w", err)
	}

	// Call out constructor
	example.NewUserManager(config, storage, risk)

	return nil
}

type AppConfig struct {
	MinPwdLen int `json:"minPwdLen"`
	MaxPwdLen int `json:"maxPwdLen"`
}

func (a *AppConfig) GetDBPool() *sql.DB {
	// intentionally omitted from example
	return nil
}

func (a *AppConfig) GetMinPwdLen() int {
	return a.MinPwdLen
}

func (a *AppConfig) GetMaxPwdLen() int {
	return a.MaxPwdLen
}
