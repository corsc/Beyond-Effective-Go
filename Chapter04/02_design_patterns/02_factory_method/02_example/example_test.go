package _2_example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	scenarios := []struct {
		desc            string
		inFormat        string
		expectedHeading string
		expectedBold    string
		expectedList    string
	}{
		{
			desc:            "markdown",
			inFormat:        "md",
			expectedHeading: "# Header",
			expectedBold:    "**Bold Text**",
		},
		{
			desc:            "HTML",
			inFormat:        "html",
			expectedHeading: "<h1>Header</h1>",
			expectedBold:    "<strong>Bold Text</strong>",
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			documentFormat := NewDocumentFormat(scenario.inFormat)

			resultHeader := documentFormat.Header("Header")
			resultBold := documentFormat.Bold("Bold Text")

			assert.Equal(t, scenario.expectedHeading, resultHeader, scenario.desc)
			assert.Equal(t, scenario.expectedBold, resultBold, scenario.desc)
		})
	}

}
