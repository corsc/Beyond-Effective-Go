package _1_unnecessary_construction

import (
	"crypto/md5"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s ./Chapter03/05_patterns/02_unnecessary_conversion/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/02_unnecessary_conversion
// BenchmarkExample-8   	100000000	       226 ns/op
// BenchmarkFixed-8     	100000000	       171 ns/op

func sign(msg string) string {
	hash := md5.New()

	result := hash.Sum([]byte(msg))
	return string(result)
}

func BenchmarkExample(b *testing.B) {
	input := "Hello World!"

	for i := 0; i < b.N; i++ {
		sign(input)
	}
}

func signFixed(msg []byte) []byte {
	hash := md5.New()
	return hash.Sum(msg)
}

func BenchmarkFixed(b *testing.B) {
	input := []byte("Hello World!")

	for i := 0; i < b.N; i++ {
		signFixed(input)
	}
}
