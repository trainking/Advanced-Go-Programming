package main

/*
#include<errno.h>
static int add(int a,int b) {
	return a + b;
}
static int div(int a, int b) {
	if (b == 0) {
		errno = EINVAL;
		return 0;
	}
	return a / b;
}
static void noreturn() {}
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.add(1,2))
	v, err0:= C.div(6,3)
	fmt.Println(v, err0)
	v1,err1:=C.div(1,0)
	fmt.Println(v1,err1)
	v3, _ := C.noreturn()
	fmt.Println(v3)
}