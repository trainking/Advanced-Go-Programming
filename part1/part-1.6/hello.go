package main

import (
	"fmt"
	"sync"
)

func main() {
	// var mu sync.Mutex
	// mu.Lock()
	// go func() {
	// 	fmt.Println("你好，世界！")
	// 	mu.Unlock()
	// }()
	// mu.Lock()

	// 利用通道同步
	// done := make(chan int)
	// go func() {
	// 	fmt.Println("你好，世界!")
	// 	done<-1
	// }()
	// r := <-done
	// fmt.Println("接收到r:", r)

	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(v int) {
			fmt.Println("你好世界:", v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}