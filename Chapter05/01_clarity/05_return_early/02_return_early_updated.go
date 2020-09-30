package _5_return_early

import (
	"errors"
)

func validateItemsUpdated(items []item) error {
	if len(items) > maxItems {
		return errors.New("too many items")
	}

	return nil
}
