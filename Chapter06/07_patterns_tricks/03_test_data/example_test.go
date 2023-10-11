package _3_test_data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	config, err := os.ReadFile("testdata/test-config.json")
	require.NoError(t, err)

	expected := `{"address": "0.0.0.0:8080"}`
	assert.Equal(t, expected, string(config))
}

func TestGenerateJSON(t *testing.T) {
	destination := "testdata/result.json"
	testFixture := "testdata/expected.json"

	// call function under test
	resultErr := generateReceiptFile(destination)
	require.NoError(t, resultErr)

	// clean up the created file
	defer os.Remove(destination)

	// compare the generated file with the expected file
	resultContents, err := os.ReadFile(destination)
	require.NoError(t, err)

	expectedContents, err := os.ReadFile(testFixture)
	require.NoError(t, err)

	assert.Equal(t, string(expectedContents), string(resultContents))
}

func TestGenerateJSONWithGenerator(t *testing.T) {
	destination := "testdata/result.json"
	testFixture := "testdata/expected.json"

	if os.Getenv("UPDATE_FIXTURES") == "true" {
		generateReceiptFile(testFixture)
		return
	}

	// call function under test
	resultErr := generateReceiptFile(destination)
	require.NoError(t, resultErr)

	// clean up the created file
	defer os.Remove(destination)

	// compare the generated file with the expected file
	resultContents, err := os.ReadFile(destination)
	require.NoError(t, err)

	expectedContents, err := os.ReadFile(testFixture)
	require.NoError(t, err)

	assert.Equal(t, string(expectedContents), string(resultContents))
}

func generateReceiptFile(filename string) error {
	receipt := &Receipt{
		Name:  "Sophia",
		Total: 12.34,
	}

	contents, err := json.Marshal(receipt)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, contents, 0644)
}

type Receipt struct {
	Name  string
	Total float64
}
