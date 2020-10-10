package _1_copy_on_write

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyOnWrite(t *testing.T) {
	Store(&Config{Host: "0.0.0.0", Port: 1234})

	// multiple loads receive the same pointer
	firstLoad := Load()
	secondLoad := Load()

	// compare point address
	firstLoadAddress := fmt.Sprintf("%p", firstLoad)
	secondLoadAddress := fmt.Sprintf("%p", secondLoad)
	assert.Equal(t, firstLoadAddress, secondLoadAddress)

	// change the config
	Store(&Config{Host: "1.2.3.4", Port: 1234})

	// load should return a different address
	thirdLoad := Load()
	thirdLoadAddress := fmt.Sprintf("%p", thirdLoad)
	assert.NotEqual(t, firstLoadAddress, thirdLoadAddress)
}
