package _7_concrete

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

// $ go test -run=Bench. -bench=. -benchmem -benchtime=10s ./Chapter03/05_patterns/07_allocations2/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Advanced-Go-Programming/Chapter03/05_patterns/07_allocations2
// BenchmarkSign-8        	300000000	        71.5 ns/op	      53 B/op	       1 allocs/op
// BenchmarkSignFixed-8   	500000000	        65.9 ns/op	      48 B/op	       0 allocs/op

func Sign(writer io.Writer) error {
	_, err := writer.Write([]byte("Signed by Me!"))
	return err
}

func BenchmarkSign(b *testing.B) {
	// inputs
	in := bytes.NewBufferString("test")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := Sign(in)
		require.NoError(b, err)
	}
}

func SignFixed(writer *bytes.Buffer) error {
	_, err := writer.Write([]byte("Signed by Me!"))
	return err
}

func BenchmarkSignFixed(b *testing.B) {
	// inputs
	in := bytes.NewBufferString("test")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := SignFixed(in)
		require.NoError(b, err)
	}
}
