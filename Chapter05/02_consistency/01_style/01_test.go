package _1_style

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTestExample(t *testing.T) {
	input := "value"
	var expected error

	// vanilla
	resultErr := doSomething(input)
	if resultErr != expected {
		log.Fatalf("doSomething, result: %v; expected: %v", resultErr, expected)
	}

	// test assertion library
	resultErr = doSomething(input)
	require.NoError(t, expected, resultErr)
}

func doSomething(input string) error {
	return nil
}
