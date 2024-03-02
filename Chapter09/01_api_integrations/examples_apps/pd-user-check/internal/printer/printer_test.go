package printer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinterAPI_Print(t *testing.T) {
	scenarios := []struct {
		desc             string
		inUser           *User
		inShowOnlyErrors bool
		expectedLines    int
	}{
		{
			desc: "show all users",
			inUser: &User{
				ID:       "ABC123",
				Name:     "Bob",
				Email:    "bob@example.com",
				Teams:    56,
				EmailSet: true,
				PhoneSet: false,
				SMSSet:   true,
				PushSet:  false,
			},
			inShowOnlyErrors: false,
			expectedLines:    2,
		},
		{
			desc: "show only errors - no errors",
			inUser: &User{
				ID:       "ABC123",
				Name:     "Bob",
				Email:    "bob@example.com",
				Teams:    56,
				EmailSet: true,
				PhoneSet: true,
				SMSSet:   true,
				PushSet:  true,
			},
			inShowOnlyErrors: true,
			expectedLines:    1,
		},
		{
			desc: "show only errors - with errors",
			inUser: &User{
				ID:       "ABC123",
				Name:     "Bob",
				Email:    "bob@example.com",
				Teams:    56,
				EmailSet: false,
				PhoneSet: false,
				SMSSet:   false,
				PushSet:  false,
			},
			inShowOnlyErrors: true,
			expectedLines:    2,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			validator := &PrinterAPI{
				RequireEmail: true,
				RequireSMS:   true,
				RequirePhone: true,
				RequirePush:  true,
			}

			buffer := &bytes.Buffer{}
			resultErr := validator.Print(buffer, scenario.inUser, scenario.inShowOnlyErrors)

			assert.NoError(t, resultErr, scenario.desc)
			assert.Equal(t, scenario.expectedLines, len(strings.Split(buffer.String(), "\n")), scenario.desc)
		})
	}
}
