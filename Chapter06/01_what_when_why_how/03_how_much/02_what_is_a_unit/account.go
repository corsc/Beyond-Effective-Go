package _2_what_is_a_unit

import (
	"errors"
)

type Account struct{}

func (a *Account) Transfer(amount int, recipient string) error {
	return errors.New("not implemented")
}
