package _9_avoid_foreign_idioms

import (
	"encoding/json"
)

type ObjectPool interface {
	Get() *myObject
	Put(object *myObject)
}

func NewEncoderDI(pool ObjectPool) *EncoderDI {
	return &EncoderDI{
		pool: pool,
	}
}

type EncoderDI struct {
	pool ObjectPool
}

func (e *EncoderDI) Encode() ([]byte, error) {
	object := e.pool.Get()
	result, resultErr := json.Marshal(object)
	e.pool.Put(object)

	return result, resultErr
}
