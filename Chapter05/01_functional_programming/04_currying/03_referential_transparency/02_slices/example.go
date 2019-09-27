package _2_slices

import (
	"fmt"
	"strings"
)

func ToLower(in []string) {
	for index := range in {
		in[index] = strings.ToLower(in[index])
	}
}

func printAll(in []string) {
	for index, value := range in {
		fmt.Printf("%d: %s\n", index, value)
	}
}
