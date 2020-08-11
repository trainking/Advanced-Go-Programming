package main
//#include<stdio.h>
import "C"

func main() {
	// println("hello cgo")
	C.puts(C.CString("Hello World"))    // 使用C.puts必须在注释中引入 stdio.h
}