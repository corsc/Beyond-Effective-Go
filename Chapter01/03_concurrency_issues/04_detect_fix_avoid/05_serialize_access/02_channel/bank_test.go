package _1_mutex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank(t *testing.T) {
	// create some accounts
	accountMay := NewAccount()
	accountMay.Deposit(100)

	accountSophia := NewAccount()

	// transfer from May to Sophia
	accountMay.Transfer(25, accountSophia)

	// we need to use in-delta because we are using floats
	assert.Equal(t, 75, accountMay.Balance(), 0.1)
	assert.InDelta(t, 25, accountSophia.Balance(), 0.1)
}

func TestBank_noDataRace(t *testing.T) {
	// create some accounts
	account := NewAccount()

	// transfer to yourself
	account.Transfer(25, account)
}
