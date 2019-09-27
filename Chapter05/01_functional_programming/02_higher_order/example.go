package _2_higher_order

import (
	"strings"
)

func forEach(in []string, operation func(string) string) []string {
	out := make([]string, len(in))

	for index, thisString := range in {
		// apply transformation using supplied operation
		result := operation(thisString)

		out[index] = result
	}

	return out
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
