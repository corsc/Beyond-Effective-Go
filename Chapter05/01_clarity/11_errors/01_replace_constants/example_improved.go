package _1_replace_constants

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrBadInput      = errors.New("bad input")
	ErrInvalidFormat = errors.New("invalid format")
	ErrParseFailed   = errors.New("failed to parse version")
)

// Extract version from ID in the format xx-yyyy
func ExtractVersionFromIDImproved(id string) (int, error) {
	if id == "" {
		return 0, ErrBadInput
	}

	chunks := strings.Split(id, "-")
	if len(chunks) < 1 {
		return 0, ErrInvalidFormat
	}

	version, err := strconv.Atoi(chunks[1])
	if err != nil {
		return 0, ErrParseFailed
	}

	return version, nil
}
