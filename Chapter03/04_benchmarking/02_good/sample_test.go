package _2_good

import (
	"fmt"
	"testing"
)

// $ go test -run=Bench. -bench=. -benchtime=15s ./Chapter03/04_benchmarking/02_good/
// goos: darwin
// goarch: amd64
// pkg: github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/04_benchmarking/02_good
// BenchmarkPerson1-8      	100000000	       192 ns/op
// BenchmarkPerson10-8     	10000000	      2449 ns/op
// BenchmarkPerson1000-8   	   10000	   1957673 ns/op

// ensure the test has a side effect
var result string

func benchToString(b *testing.B, people []Person) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = toString(people)
	}
}

func BenchmarkPerson1(b *testing.B) {
	in := []Person{
		{ID: 1, Name: "John"},
	}

	benchToString(b, in)
}

func BenchmarkPerson10(b *testing.B) {
	in := make([]Person, 10)
	for x := 0; x < 10; x++ {
		in[x] = Person{ID: x, Name: fmt.Sprintf("test %d", x)}
	}

	benchToString(b, in)
}

func BenchmarkPerson1000(b *testing.B) {
	in := make([]Person, 1000)
	for x := 0; x < 1000; x++ {
		in[x] = Person{ID: x, Name: fmt.Sprintf("test %d", x)}
	}

	benchToString(b, in)
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
