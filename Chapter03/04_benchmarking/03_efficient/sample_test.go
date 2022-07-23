package _3_efficient

import (
	"fmt"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=15s ./Chapter03/04_benchmarking/03_efficient/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Beyond-Effective-Go/Chapter03/04_benchmarking/03_efficient
// BenchmarkToString/1_person-8         	100000000	       177 ns/op
// BenchmarkToString/10_people-8        	10000000	      2585 ns/op
// BenchmarkToString/100_people-8       	  500000	     44207 ns/op

// ensure the test has a side effect
var result string

func BenchmarkToString(b *testing.B) {
	scenarios := []struct {
		desc        string
		totalPeople int
	}{
		{
			desc:        "1 person",
			totalPeople: 1,
		},
		{
			desc:        "10 people",
			totalPeople: 10,
		},
		{
			desc:        "100 people",
			totalPeople: 100,
		},
	}

	for _, s := range scenarios {
		scenario := s
		b.Run(scenario.desc, func(b *testing.B) {
			people := make([]Person, scenario.totalPeople)
			for x := 0; x < scenario.totalPeople; x++ {
				people[x] = Person{
					ID:   x,
					Name: fmt.Sprintf("test %d", x),
				}
			}

			for i := 0; i < b.N; i++ {
				result = toString(people)
			}
		})
	}
}

func toString(people []Person) string {
	var out string
	for _, person := range people {
		out += fmt.Sprintf("ID: %d\nName: %s", person.ID, person.Name)
	}
	return out
}

type Person struct {
	ID   int
	Name string
}
