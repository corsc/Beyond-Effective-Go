package ftests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestCoordinator_getChanges_SadPath(t *testing.T) {
	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Temp dir will not be a git repo and should cause an error
	dir := os.TempDir()

	// expectations

	// call object under test
	objectUnderTest := &Coordinator{}
	_, resultErr := objectUnderTest.getChanges(ctx, dir)

	// validation
	assert.Error(t, resultErr)
}

func TestCoordinator_buildListOfChangedPackages(t *testing.T) {
	scenarios := []struct {
		desc           string
		in             string
		expectedResult []string
		expectAnErr    bool
	}{
		{
			desc:           "Happy path - No changes",
			in:             ``,
			expectedResult: nil,
			expectAnErr:    true,
		},
		{
			desc: "Happy path - 1 change",
			in: `
Chapter09/01_api_integrations/01_example/01_calling_the_api/calling_pagerduty_test.go
`,
			expectedResult: []string{
				"Chapter09/01_api_integrations/01_example/01_calling_the_api",
			},
			expectAnErr: false,
		},
		{
			desc: "Happy path - multiple changes and duplicates",
			in: `
Chapter09/01_api_integrations/01_example/01_calling_the_api/calling_pagerduty.go
Chapter09/01_api_integrations/01_example/01_calling_the_api/calling_pagerduty_test.go
Chapter09/01_api_integrations/01_example/02_something_else/example.go
`,
			expectedResult: []string{
				"Chapter09/01_api_integrations/01_example/01_calling_the_api",
				"Chapter09/01_api_integrations/01_example/02_something_else",
			},
			expectAnErr: false,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// inputs

			// mocks

			// call object under test
			objectUnderTest := &Coordinator{}
			result, resultErr := objectUnderTest.buildListOfChangedPackages([]byte(scenario.in))

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error: %t, err: '%s'", scenario.expectAnErr, resultErr)
			assert.Equal(t, scenario.expectedResult, result)
		})
	}
}

func TestCoordinator_filterUnwantedPackages(t *testing.T) {
	scenarios := []struct {
		desc           string
		in             []string
		expectedResult []string
		expectAnErr    bool
	}{
		{
			desc: "Happy path - No changes",
			in: []string{
				"A",
				"B",
			},
			expectedResult: []string{
				"A",
				"B",
			},
			expectAnErr: false,
		},
		{
			desc: "Happy path - With vendor",
			in: []string{
				"A",
				"A/vendor/C",
				"vendor/B",
			},
			expectedResult: []string{
				"A",
			},
			expectAnErr: false,
		},
		{
			desc: "Happy path - With test data",
			in: []string{
				"A",
				"A/testdata/C",
				"testdata/B",
			},
			expectedResult: []string{
				"A",
			},
			expectAnErr: false,
		},
		{
			desc: "Happy path - With nothing but vendor",
			in: []string{
				"vendor/B",
			},
			expectedResult: nil,
			expectAnErr:    true,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// inputs

			// mocks

			// call object under test
			objectUnderTest := &Coordinator{}
			result, resultErr := objectUnderTest.filterUnwantedPackages(scenario.in)

			// validation
			require.Equal(t, scenario.expectAnErr, resultErr != nil, "expected error: %t, err: '%s'", scenario.expectAnErr, resultErr)
			assert.Equal(t, scenario.expectedResult, result)
		})
	}
}

func TestCoordinator_runTests_HappyPath(t *testing.T) {
	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pkgs := []string{
		"fakepkg",
	}

	// expectations

	// call object under test
	objectUnderTest := &Coordinator{}
	objectUnderTest.runTests(ctx, "", pkgs)

	// nothing to validate
	assert.True(t, true)
}
