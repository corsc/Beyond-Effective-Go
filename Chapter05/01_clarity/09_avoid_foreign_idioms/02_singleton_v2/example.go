package _2_singleton_v2

import (
	"encoding/json"
)

type ObjectPool interface {
	Get() *myObject
	Put(object *myObject)
}

func NewEncoder(pool ObjectPool) *Encoder {
	return &Encoder{
		pool: pool,
	}
}

type Encoder struct {
	pool ObjectPool
}

func (e *Encoder) Encode() ([]byte, error) {
	object := e.pool.Get()
	result, resultErr := json.Marshal(object)
	e.pool.Put(object)

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
