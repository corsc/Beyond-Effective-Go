package _1_mutex

func NewAccount() *Account {
	account := &Account{
		actionsCh: make(chan *action),
	}

	go account.processActions()

	return account
}

type Account struct {
	actionsCh chan *action
}

func (a *Account) Transfer(amount int, destination *Account) {
	// withdraw from this account
	withdrawAction := newAction(-1 * amount)
	a.actionsCh <- withdrawAction

	// deposit into destination account
	destination.Deposit(amount)
}

func (a *Account) Deposit(amount int) {
	action := newAction(amount)

	a.actionsCh <- action
}

func (a *Account) Balance() int {
	action := newAction(0)

	a.actionsCh <- action

	return <-action.resultCh
}

func (a *Account) processActions() {
	var balance int

	for thisAction := range a.actionsCh {
		balance += thisAction.amount
		thisAction.resultCh <- balance
	}
}

func newAction(amount int) *action {
	return &action{
		amount: amount,
		// use buffered channel so we don't block the processor
		resultCh: make(chan int, 1),
	}
}

type action struct {
	amount   int
	resultCh chan int
}
