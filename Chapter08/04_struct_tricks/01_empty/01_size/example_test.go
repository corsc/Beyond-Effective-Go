package _1_size

import (
	"fmt"
	"unsafe"
)

type emptyStruct struct{}

func Example() {
	myEmptyStruct := emptyStruct{}

	// Output: 0 bytes
	fmt.Printf("%d bytes", unsafe.Sizeof(myEmptyStruct))
}
