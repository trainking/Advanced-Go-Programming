package main

/*
struct A {
	int size: 10;  // 位字段无法访问
	float arr[];  // 零长的数组也无法访问
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	// 错误，这两个字段都不能直接访问，只能通过C语言的辅助函数调用
	fmt.Println(a.size)
	fmt.Println(a.arr)
}