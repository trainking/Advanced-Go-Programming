package main

// #include <stdio.h>
// #include <stdlib.h>
//#include "hello.h"
import "C"

func main() {
	C.SayHello(C.CString("Hello, World!\n"))
}