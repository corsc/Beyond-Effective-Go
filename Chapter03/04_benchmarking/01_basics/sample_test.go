package _1_basics

import (
	"testing"
)

func BenchmarkBasics(b *testing.B) {
	// build input definition and perform initialization

	// Remove init time from measurement
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// call function/object under test
	}
}
