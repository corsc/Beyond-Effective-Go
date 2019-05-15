package _8_read_from_closed

import (
	"fmt"
)

func Example() {
	dataCh := make(chan string)

	data := <-dataCh

	// "use" the data
	fmt.Printf("result: %#v\n", data)
}
