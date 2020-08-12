package main

/*
struct A {
	int i;
	float f;
	int type;
	int _type;
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
	fmt.Println(a._type)  // type 是golang的关键字，前加_规避, 如果存在_type成员，则type不可访问（被屏蔽）
}