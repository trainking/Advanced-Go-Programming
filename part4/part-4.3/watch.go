package main

import (
	"sync"
	"fmt"
)

type KVStoreService struct {
	m   map[string]string
	filter map[string]func(key string)
	mu sync.Mutex
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m: make(map[string]string),
		filter: make(map[string]func(key string))
	}
}

func main() {
	
}