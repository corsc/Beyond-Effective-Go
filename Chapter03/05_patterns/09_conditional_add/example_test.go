package _9_conditional_add

import (
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=10s ./Chapter03/05_patterns/09_conditional_add/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/05_patterns/09_conditional_add
// BenchmarkExample-8   	    1000	  15271128 ns/op
// BenchmarkFixed-8     	   20000	    885699 ns/op

func BenchmarkExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var collection []int

		for x := 0; x < 10000; x++ {
			for _, thisItem := range collection {
				if thisItem == x {
					return
				}
			}
			collection = append(collection, x)
		}
	}
}

func BenchmarkFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var collection []int
		var hash = map[int]struct{}{}

		for x := 0; x < 10000; x++ {
			_, found := hash[x]
			if found {
				return
			}
			collection = append(collection, x)
			hash[x] = struct{}{}
		}
	}
}
