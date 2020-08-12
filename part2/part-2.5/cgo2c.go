package main

//int sum(int a,int b) {return a +b;}
import "C"
import "fmt"

func main() {
	fmt.Println(C.sum(1, 1));
}

// go tool cgo cgo2c.go 生成一个_obj文件夹，里面有整个cgo调用生成的文件
