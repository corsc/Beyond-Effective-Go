package _8_read_from_closed

import (
	"fmt"
)

func ExampleExtended() {
	dataCh := make(chan string)

	data, isClosed := <-dataCh

	// "use" the data
	fmt.Printf("result: %#v isClosed: %t\n", data, isClosed)
}
