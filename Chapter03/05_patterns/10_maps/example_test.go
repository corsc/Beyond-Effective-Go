package _0_maps

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// $ go test -run=Bench. -bench=. -benchtime=10s ./Chapter03/05_patterns/10_maps/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/10_maps
// BenchmarkMapString-8   	100000000	       198 ns/op
// BenchmarkMapInt-8      	100000000	       137 ns/op

func BenchmarkMapString(b *testing.B) {
	data := map[string]string{}
	for x := 100000; x < 1000000; x++ {
		data[strconv.Itoa(x)] = "X"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := data[strconv.Itoa(500000)]
		require.Equal(b, "X", result)
	}
}

func BenchmarkMapInt(b *testing.B) {
	data := map[int]string{}
	for x := 100000; x < 1000000; x++ {
		data[x] = "X"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := data[500000]
		require.Equal(b, "X", result)
	}
}
