package main


import (
	"sync"
	"sync/atomic"
)

type singleton struct {	}

var (
	instance *singleton
	initialized unit32
	mu sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 建立实例的过程互斥化
	mu.Loack()
	defer mu.Unlock()
	if instance == nil {
		def atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}