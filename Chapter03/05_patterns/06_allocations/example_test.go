package _6_allocations

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// $ go test -run=Bench. -bench=. -benchmem -benchtime=10s ./Chapter03/05_patterns/06_allocations/
// goos: darwin
// goarch: amd64
// pkg: github.com/corsc/Advanced-Go-Programming/Chapter03/05_patterns/06_allocations
// BenchmarkExample-8   	20000000	      1068 ns/op	     272 B/op	       7 allocs/op
// BenchmarkFixed-8     	20000000	      1016 ns/op	     240 B/op	       6 allocs/op

func Marshal(in []byte) (User, error) {
	out := User{}

	err := json.Unmarshal(in, &out)

	return out, err
}

func BenchmarkExample(b *testing.B) {
	// inputs
	testUser := User{ID: 666, Name: "Bob"}
	data, err := json.Marshal(testUser)
	require.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result, err := Marshal(data)

		require.NoError(b, err)
		require.Equal(b, testUser, result)
	}
}

func MarshalFixed(in []byte, user *User) error {
	return json.Unmarshal(in, user)
}

func BenchmarkFixed(b *testing.B) {
	// inputs
	testUser := User{ID: 666, Name: "Bob"}
	data, err := json.Marshal(testUser)
	require.NoError(b, err)

	result := &User{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := MarshalFixed(data, result)

		require.NoError(b, err)
		require.EqualValues(b, testUser, *result)
	}
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
