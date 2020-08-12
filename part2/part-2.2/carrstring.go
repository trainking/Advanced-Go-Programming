package main

/*
static char arr[10];
static char *s = "Hello";
*/
import "C"
import (
	"fmt"
)

func main() {
	// 通过reflect.SliceHeader转换
	var arr0 []byte
	var arr0Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arr0))
	arr0Hdr.Data = uintptr(unsafe.Pointer(&C.arr[0]))
	arr0Hdr.Len = 10
	arr0Hdr.Cap = 10
	// 通过切片语法转换
	arr1 := (*[31]byte)(unsafe.Pointer(&C.arr[0]))[:10:10]
}