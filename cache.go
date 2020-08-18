package main

import (
	"time"
	"sync"
	"unsafe"
	"fmt"
)

type Item struct {
	val interface{}   // 数据项
	expires int64     // 超时
}

// 判断数据项是否过期
func (item Item) isExpire() bool {
	// 无 设置
	if item.expires == 0 {
		return false
	}

	// fmt.Println("DEBUG now:", time.Now().UnixNano())
	// fmt.Println("DEBUG expires:", item.expires)
	return time.Now().UnixNano() > item.expires
}

type Cache struct {
	items map[string]Item
	mu sync.RWMutex
	interval   time.Duration  // 清理数据周期
	stopGC   chan bool
	max  int64               // 最大内存上限
}

func (c *Cache) deleteExprie() {
	now := time.Now().UnixNano()
	c.mu.Lock()
	defer c.mu.Unlock()

	for k,v :=range c.items {
		// 为0不删除
		if v.expires > 0 && now > v.expires {
			delete(c.items, k)
		}
	}
}

func (c *Cache) SetMaxMemory(size string) bool {
	var _max int64
	switch size {
		case "1KB":
			_max = 1024
		case "100KB":
			_max = 100 * 1024
		case "1MB":
			_max = 1024 * 1024
		case "2MB":
			_max = 2 * 1024 * 1024
		case "1GB":
			_max = 1024 * 1024 * 1024
	}
	c.max = _max
	return true
}

func (c *Cache) Set(key string, val interface{}, expire time.Duration) {
	var _item Item
	_item.val = val
	if expire > 0 {
		_item.expires = time.Now().Add(expire * time.Second).UnixNano()
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	// 超过最大内存
	if int64(unsafe.Sizeof(*c)) + int64(unsafe.Sizeof(_item)) > c.max {
		return
	}
	c.items[key] = _item
}

func (c *Cache) Get(key string) (interface{}, bool) {
	_item, _found := c.items[key]
	if !_found {
		return nil, false
	}
	if _item.isExpire() {
		return nil, false
	}
	return _item.val, true
}

func (c *Cache) Del(key string) bool {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
	return true
}

func (c *Cache) Flush() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = map[string]Item{}
	return true
}

func (c *Cache) Keys() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return int64(len(c.items))
}

func (c *Cache) gc() {
	_t := time.NewTicker(c.interval)
	for {
		select {
		case <-_t.C:
			c.deleteExprie()
		case <- c.stopGC:
			_t.Stop()
			return
		}
	}
}

// 创建缓存
func CreateCache() *Cache {
	c := &Cache {
		items: map[string]Item{},
		interval: 5 * time.Second,   // 默认5执行一次清理
		stopGC: make(chan bool),
	}
	go c.gc()
	return c
}

// 删除缓存
func RemoveCache(c *Cache) {
	c.stopGC <- true
	*c = Cache{}
}

func main() {
	var v *Cache = CreateCache()
	// 设置缓存大小
	v.SetMaxMemory("1KB")
	// 设置项
	v.Set("k1", "v1", 0)
	// 获取项
	// time.Sleep(10 * time.Second)
	fmt.Println(v.Get("k1"))
	// fmt.Println(v.Keys())
	// fmt.Println(v.Flush())
	// fmt.Println(v.Keys())
	// fmt.Println(v.Del("k1"))
	// fmt.Println(v.Get("k1"))
	RemoveCache(v)
}