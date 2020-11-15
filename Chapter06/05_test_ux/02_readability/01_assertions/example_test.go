package _1_assertions

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVanillaGo(t *testing.T) {
	t.Skip("disabled")

	resultErr := errors.New("sample")
	var expectedErr error

	if resultErr != expectedErr {
		t.Fatalf("resultErr '%v' did not match expectedErr '%v'", resultErr, expectedErr)
	}
}

func TestAssertion(t *testing.T) {
	t.Skip("disabled")

	resultErr := errors.New("sample")
	var expectedErr error

	require.Equal(t, expectedErr, resultErr, "resultErr did not match expectedErr")
}
