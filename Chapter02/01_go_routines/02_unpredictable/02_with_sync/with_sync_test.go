package _2_with_sync

import (
	"testing"
)

func TestLoadAll(t *testing.T) {
	loader := NewLoader()
	loader.LoadAll()
}
