package _1_maps

import (
	"fmt"
	"strings"
)

func ToLower(in map[string]string) {
	for key := range in {
		in[key] = strings.ToLower(in[key])
	}
}

func printAll(in map[string]string) {
	for key, value := range in {
		fmt.Printf("%s: %s\n", key, value)
	}
}
