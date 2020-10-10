package _1_mutex

import (
	"sync"
)

func NewAccount() *Account {
	return &Account{
		mutex: &sync.Mutex{},
	}
}

type Account struct {
	balance int
	mutex   *sync.Mutex
}

func (a *Account) Transfer(amount int, destination *Account) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// withdraw from this account
	a.balance -= amount

	// deposit into destination account
	destination.Deposit(amount)
}

func (a *Account) Deposit(amount int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.balance += amount
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	return a.balance
}
