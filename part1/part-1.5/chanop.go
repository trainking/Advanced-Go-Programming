package main

import (
	"fmt"
)

// func Println(s string) {
// 	fmt.Println(s)
// }

func main() {
	done := make(chan int)
	go func() {
		// Println("hello, world!")
		fmt.Println("hello, world!")
		done <- 1
	}()
	<-done
}