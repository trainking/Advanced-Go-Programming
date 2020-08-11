package main

/*
#cgo LDFLAGS: -L ./ -l
#include "hello.h"
*/
import "C"

func main() {
	C.SayHello(C.CString("Hello, World!\n"))
}