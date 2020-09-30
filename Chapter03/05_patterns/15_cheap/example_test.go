package _5_cheap

import (
	"regexp"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s -benchmem ./Chapter03/05_patterns/15_cheap/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Beyond-Effective-Go/Chapter03/05_patterns/15_cheap
// BenchmarkExampleRegexOnlyBadInput-8    	200000000	        75.4 ns/op	       0 B/op	       0 allocs/op
// BenchmarkExampleRegexOnlyGoodInput-8   	30000000	       431 ns/op	       0 B/op	       0 allocs/op
// BenchmarkExampleFastFailBadInput-8     	10000000000	         2.21 ns/op	       0 B/op	       0 allocs/op
// BenchmarkExampleFastFailGoodInput-8    	30000000	       427 ns/op	       0 B/op	       0 allocs/op

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var result bool

func BenchmarkExampleRegexOnlyBadInput(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		result = isEmail(input)
	}
}

func BenchmarkExampleRegexOnlyGoodInput(b *testing.B) {
	input := "me@home.com"

	for i := 0; i < b.N; i++ {
		result = isEmail(input)
	}
}

func isEmail(in string) bool {
	return emailRegex.MatchString(in)
}

func BenchmarkExampleFastFailBadInput(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		result = isEmailFastFail(input)
	}
}

func BenchmarkExampleFastFailGoodInput(b *testing.B) {
	input := "me@home.com"

	for i := 0; i < b.N; i++ {
		result = isEmailFastFail(input)
	}
}

func isEmailFastFail(in string) bool {
	if in == "" {
		return false
	}

	return emailRegex.MatchString(in)
}
