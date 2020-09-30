package _1_replace_constants

import (
	"strconv"
	"strings"
)

// Extract version from ID in the format xx-yyyy
func ExtractVersionFromID(id string) int {
	if id == "" {
		return -1 // bad input
	}

	chunks := strings.Split(id, "-")
	if len(chunks) < 1 {
		return -2 // invalid format
	}

	version, err := strconv.Atoi(chunks[1])
	if err != nil {
		return -3 // failed to parse version number
	}

	return version
}
