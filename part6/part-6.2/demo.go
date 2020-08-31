package main

import (
	"sync"
	"fmt"
)

// 全局变量
var counter int
// 锁
var l sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter ++
			l.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}