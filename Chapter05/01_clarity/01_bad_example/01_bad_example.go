package _1_bad_example

import (
	"sync"
)

func do(o O) error {
	if e := v(o); e == nil {
		if r, e := c(o); e == nil {
			ec := make(chan error, 2)
			wg := &sync.WaitGroup{}
			wg.Add(2)
			go s(o, r, ec)
			go i(o, ec)
			wg.Wait()
			for e := range ec {
				if e != nil {
					return e
				}
			}
			return nil
		} else {
			return e
		}
	} else {
		return e
	}
}

func v(o O) error {
	return nil
}

func c(o O) (int, error) {
	return 0, nil
}

func s(o O, n int, ec chan<- error) {
	ec <- nil
}

func i(o O, ec chan<- error) {
	ec <- nil
}

type O struct{}
