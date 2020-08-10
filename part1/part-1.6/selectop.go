package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	// defer func() {
	// 	fmt.Println("Done!")
		
	// }()
	for {
		select{
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	// ch := make(chan int)
	// go func() {
	// 	for {
	// 		select {
	// 		case ch <- 0:
	// 		case ch <- 1:
	// 		// default:
	// 		// 	fmt.Println("default")
	// 		}
	// 	}
	// }()

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	// cannel := make(chan bool)
	ctx, cannel := context.WithTimeout(context.Background(), 10 * time.Second)
	var wg sync.WaitGroup
	for i:=0; i <= 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}
	time.Sleep(time.Second)
	cannel()
	wg.Wait()
}