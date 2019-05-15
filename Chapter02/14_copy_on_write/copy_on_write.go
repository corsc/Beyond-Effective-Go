package _4_copy_on_write

import (
	"sync/atomic"
	"time"
	"unsafe"
)

var config unsafe.Pointer

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func Load() *Config {
	return (*Config)(atomic.LoadPointer(&config))
}

func Store(cfg *Config) {
	atomic.StorePointer(&config, unsafe.Pointer(cfg))
}
