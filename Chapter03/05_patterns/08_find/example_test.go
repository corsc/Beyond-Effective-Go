package _8_find

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// $ go test -run=Bench. -bench=. -benchmem -benchtime=10s ./Chapter03/05_patterns/08_find/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Beyond-Effective-Go/Chapter03/05_patterns/08_find
// BenchmarkExample-8        	  500000	     34997 ns/op	       0 B/op	       0 allocs/op
// BenchmarkExampleFixed-8   	500000000	        41.2 ns/op	       0 B/op	       0 allocs/op

func BenchmarkExample(b *testing.B) {
	total := 10000
	people := make([]*Person, total)
	for x := 0; x < total; x++ {
		people[x] = &Person{
			ID:   x,
			Name: "Test " + strconv.Itoa(x),
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := findPerson(people, "Test 9999")
		require.NotNil(b, result)
	}
}

func findPerson(people []*Person, name string) *Person {
	for _, person := range people {
		if person.Name == name {
			return person
		}
	}

	return nil
}

func BenchmarkExampleFixed(b *testing.B) {
	total := 10000
	people := make([]*Person, total)
	peopleMap := make(map[string]int, total)
	for x := 0; x < total; x++ {
		name := "Test " + strconv.Itoa(x)

		people[x] = &Person{
			ID:   x,
			Name: name,
		}

		peopleMap[name] = x
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result := findPersonMap(people, peopleMap, "Test 9999")
		require.NotNil(b, result)
	}
}

func findPersonMap(people []*Person, peopleMap map[string]int, name string) *Person {
	index, found := peopleMap[name]
	if !found {
		return nil
	}

	return people[index]
}

type Person struct {
	ID   int
	Name string
}
