package _3_output

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinter_outputHeader(t *testing.T) {
	expected := `Name                                     Email                                                        Email    Phone    SMS      Push    
-----------------------------------------------------------------------------------------------------------------------------------------------------
`
	result := &bytes.Buffer{}

	objectUnderTest := &Printer{
		logger: result,
	}
	objectUnderTest.outputHeader()

	assert.Equal(t, expected, result.String())
}

func TestPrinter_outputLine(t *testing.T) {
	input := &User{
		ID:         "A",
		Name:       "B",
		Email:      "C",
		EmailIsSet: true,
		PhoneIsSet: false,
		SMSIsSet:   true,
		PushIsSet:  false,
	}

	expected := `B                                        C                                                            true     false    true     false   
`
	result := &bytes.Buffer{}

	objectUnderTest := &Printer{
		logger: result,
	}
	objectUnderTest.outputLine(input)

	assert.Equal(t, expected, result.String())
}
