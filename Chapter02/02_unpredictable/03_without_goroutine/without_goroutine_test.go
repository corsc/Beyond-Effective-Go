package _3_without_goroutine

import (
	"testing"
)

func TestLoadAll(t *testing.T) {
	loader := NewLoader()
	loader.LoadAll()
	loader.LoadAll()
}
