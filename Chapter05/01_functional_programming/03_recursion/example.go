package _3_recursion

import (
	"strings"
)

func forEach(in []string, operation func(string) string) []string {
	return forE(in, 0, operation)
}

func forE(in []string, index int, operation func(string) string) []string {
	if index >= len(in) {
		return nil
	}

	// apply transformation using supplied operation
	result := operation(in[index])

	return append([]string{result}, forE(in, index+1, operation)...)
}

func toUpper(in string) string {
	return strings.ToUpper(in)
}

func reverse(in string) string {
	out := ""
	for x := len(in) - 1; x >= 0; x-- {
		out += in[x : x+1]
	}
	return out
}
