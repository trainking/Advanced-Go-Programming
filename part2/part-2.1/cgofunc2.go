package main

//#include "hello.h"
import "C"

/*
* 只要引入文件和c代码在一个文件夹下，直接用go build一下就可以生成
*/
func main() {
	C.SayHello(C.CString("Hello, World!\n"))
}