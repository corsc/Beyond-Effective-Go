package _5_pool

func usePooledAccountFixed() int {
	account := pool.Get().(*Account)
	defer pool.Put(account)

	// reset account to known state
	account.total = 0

	// use account object
	account.total += 5

	return account.total
}
