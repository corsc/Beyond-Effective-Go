package _0_channel_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	t.Skip("this test is broken")

	john := &Person{Name: "John"}

	// send an pointer via the channel
	peopleCh := make(chan *Person, 1)
	peopleCh <- john

	// modify the value (after the send but before the read)
	john.Name = "Paul"

	// What we receive is not what we sent
	result := <-peopleCh
	assert.Equal(t, "John", result.Name)
}

type Person struct {
	Name string
}
