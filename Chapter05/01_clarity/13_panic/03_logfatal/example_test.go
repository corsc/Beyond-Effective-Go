package _3_logfatal

import (
	"log"
	"testing"
	"time"
)

func TestGenerateHash(t *testing.T) {
	result, err := GenerateHash(time.Now().UnixNano())
	if err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

	if result == "" {
		log.Fatal("result should not be empty")
	}
}

func GenerateHash(nano int64) (string, error) {
	// implementation removed
	return "fu", nil
}
