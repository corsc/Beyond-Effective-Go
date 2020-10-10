package _1_deadlock

import (
	"errors"
	"sync"
)

type Dictionary struct {
	data  []string
	mutex sync.Mutex
}

func (d *Dictionary) Get(index int) (string, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// this line will deadlock as Count() also requires a lock
	if d.Count() >= index {
		return "", errors.New("out of bounds")
	}

	return d.data[index], nil
}

func (d *Dictionary) Count() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return len(d.data)
}
