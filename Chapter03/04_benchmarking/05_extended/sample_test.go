package _5_extended

import (
	"fmt"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=15s ./Chapter03/04_benchmarking/05_extended/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/04_benchmarking/05_extended
// BenchmarkToString/v1-8         	     100	 151002675 ns/op
// BenchmarkToString/v2-8         	   30000	    606889 ns/op

// $ go test -run=Bench. -bench=. -benchmem -benchtime=15s ./Chapter03/04_benchmarking/05_extended/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/04_benchmarking/05_extended
// BenchmarkToString/v1-8         	     200	 150015657 ns/op	1216181375 B/op	   42848 allocs/op
// BenchmarkToString/v2-8         	   30000	    598083 ns/op	  284765 B/op	    9901 allocs/op

// ensure the test has a side effect
var result string

func BenchmarkToString(b *testing.B) {
	scenarios := []struct {
		desc         string
		toStringFunc func([]Person) string
	}{
		{
			desc:         "v1",
			toStringFunc: toStringV1,
		},
		{
			desc:         "v2",
			toStringFunc: toString,
		},
	}

	totalPeople := 10000
	people := make([]Person, totalPeople)
	for x := 0; x < totalPeople; x++ {
		people[x] = Person{ID: x, Name: fmt.Sprintf("test %d", x)}
	}

	for _, s := range scenarios {
		scenario := s
		b.Run(scenario.desc, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result = scenario.toStringFunc(people)
			}
		})
	}
}

func toStringV1(people []Person) string {
	var out string
	for _, person := range people {
		out += fmt.Sprintf("ID: %d\nName: %s", person.ID, person.Name)
	}
	return out
}
