package _2_higher_order

import (
	"strings"
)

func forEach(in []string, operation func(string) string) []string {
	out := make([]string, len(in))

	for index, thisString := range in {
		// apply transformation using the supplied operation
		result := operation(thisString)

		out[index] = result
	}

	return out
}

func toUpper(in string) string {
	return strings.ToUpper(in)
}

func addPeriod(in string) string {
	return in + "."
}
