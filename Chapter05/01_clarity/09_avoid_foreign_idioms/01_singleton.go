package _9_avoid_foreign_idioms

import (
	"encoding/json"
)

var pool = &ObjectPoolSingleton{}

type Encoder struct{}

func (e *Encoder) Encode() ([]byte, error) {
	object := pool.Get()
	result, resultErr := json.Marshal(object)
	pool.Put(object)

	return result, resultErr
}

type myObject struct{}

type ObjectPoolSingleton struct{}

func (o ObjectPoolSingleton) Get() *myObject {
	// implementation removed
	return nil
}

func (o ObjectPoolSingleton) Put(_ *myObject) {
	// implementation removed
}
