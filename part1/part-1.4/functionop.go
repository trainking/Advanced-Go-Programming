package main

import (
	"fmt"
)

// 普通具名函数
func Add(a, b int) int {  // a,b 推到出都是int型
	return a + b
}

// 多参数 多返回值
func Swap(a, b int) (int, int) {
	return b,a
}

// 可变数量参数
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

func Inc() (v int) {
	defer func() {v++}()
	return 42
}

func main()  {
	// 匿名函数
	// var Add = func(a, b int) int {
	// 	return a + b
	// }

	// var a = []interface{}{123, "abc"}
	// Print(a...)
	// Print(a)
	// Print(Inc())
	Print(f(2))
	Print(g())
}

func f(x int) *int {
	return &x
}

func g() int {
	x := new(int)
	return *x
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}