package _5_return_early

import (
	"errors"
)

const maxItems = 100

func validateItems(items []item) error {
	if len(items) <= maxItems {
		return nil
	} else {
		return errors.New("too many items")
	}
}

type item interface {
	// implementation removed
}
