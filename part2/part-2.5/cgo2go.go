package main

// int sum(int a, int b);
import "C"
// export sum
func sum(a, b C.int) C.int {
	return a + b
}

// go build -buildmode=c-archive -o cgo2go.a cgo2go.go
// 上面命令将go变成C语言库
// go tool cgo cgo2go.go
// 上面命令查看生成过程文件
func main() {

}