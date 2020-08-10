package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// var total struct {
// 	sync.Mutex
// 	value int
// }

// func worker(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i <= 100; i ++ {
// 		total.Lock()
// 		total.value += i
// 		total.Unlock()
// 	}
// }

var total uint64
func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64    // 因为go 语言不会隐式地转型，所以要采用申明同样的类型来实现
	for i = 0; i <= 100; i ++  {
		atomic.AddUint64(&total, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total)
}