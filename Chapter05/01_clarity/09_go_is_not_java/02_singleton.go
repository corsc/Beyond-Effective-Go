package _9_go_is_not_java

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
