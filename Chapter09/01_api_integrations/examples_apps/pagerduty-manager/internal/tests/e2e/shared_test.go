package e2e

import "os"

type testConfig struct {
	baseURL string
}

func (t *testConfig) Filename() string {
	return "./test_data/e2e.json"
}

func (t *testConfig) AuthToken() string {
	return os.Getenv("PD_TOKEN")
}

func (t *testConfig) Debug() bool {
	return true
}

func (t *testConfig) BaseURL() string {
	return t.baseURL
}
