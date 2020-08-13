package main

/*
#include<stdio.h>
#include<stdlib.h>
int printString(const char* s) {
	printf("%s", s);
	return 1;
}
*/
import "C"
import "unsafe"
import "fmt"

func printString(s string) int {
	cs := C.CString(s)
	// 函数结束时释放cs变量空间
	defer C.free(unsafe.Pointer(cs))
	a := C.printString(cs)
	return int(a)
}

func main() {
	s := "Hello"
	g := printString(s)
	fmt.Println(g)
}