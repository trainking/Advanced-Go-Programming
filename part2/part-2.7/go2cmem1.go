package main

/*
#include<stdio.h>
void printString(const char* s, int n) {
	int i;
	for (i=0; i< n; i++) {
		putchar(s[i]);
	}
	putchar('\n');
}
*/
import "C"
import "reflect"
import "unsafe"

func printString(s string) {
	// 直接反射转换，就不需要free了
	// C 语言直接使用Go的空间
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	C.printString((*C.char)(unsafe.Pointer(p.Data)), C.int(len(s)))
}

func main() {
	s:="hello"
	printString(s)
}