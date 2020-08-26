package watch

import (
	"sync"
	"time"
	"fmt"
	"math/rand"
)

type KVStoreServiceInterface = interface {
	Set(request [2]string, reply *struct{}) error
	Get(request string, reply *string) error
	Watch(request int, reply *string) error
}

type KVStoreService struct {
	m map[string]string
	filter map[string]func(key string)
	mu  sync.Mutex
}

// 创建
func NewKVStoreService() *KVStoreService {
	return &KVStoreService {
		m: make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

// setter 
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0],kv[1]
	// 判断不同值进行过滤
	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}
	p.m[key] = value
	return nil
}

// getter
func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v,ok := p.m[key]; ok {
		*value = v
		return nil
	}
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10)
	p.mu.Lock()
	p.filter[id] = func (key string)  {
		ch <- key
	}
	p.mu.Unlock()
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
			return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
	return nil
}

