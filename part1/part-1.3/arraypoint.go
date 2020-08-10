package main

import (
	"fmt"
	// "image"
	// "io"
	// "image/png"
	// "image/jpeg"
)

func main()  {
	// var a = [...]int{1,2,3}
	// var b = &a
	// fmt.Println(a[0], a[1])
	// fmt.Println(b[0], b[1])
	// fmt.Println("fot---b----")
	// for i,v := range b {
	// 	fmt.Println(i,v)
	// }
	// fmt.Println("fot---b----")
	// for i,v := range a {
	// 	fmt.Println(i,v)
	// }
	// fmt.Println("len:", len(a))
	// fmt.Println("cap:", cap(a))
	// fmt.Println("len:", len(b))
	// fmt.Println("cap:", cap(b))

	// var times [5][0]int
	// fmt.Println(times)
	// for range times {
	// 	fmt.Println("hello")
	// }

	// 字符串数组
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{10:"世界", 0:"你好", }
	fmt.Printf("s1: %T\n", s1)
	fmt.Printf("s2: %#v\n", s2)
	fmt.Printf("s3: %#v\n", s3)

	// 结构体数组
	// var line1 [2]image.Point
	// var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	// var line3 = [...]image.Point{{0,0}, {1,1}}
	// fmt.Println(line1)
	// fmt.Println(line2)
	// fmt.Println(line3)

	// 函数数组
	// var decoder1 [2]func(io.Reader) (image.Image, error)
	// var decoder2 = [...]func(io.Reader)(image.Image, error) {
	// 	png.Decode,
	// 	jpeg.Decode,
	// }
	// fmt.Println(decoder1)
	// fmt.Println(decoder2)

	// 接口数组
	// var unknown1 [2]interface{}
	// var unknown2 = [...]interface{} {123, "你好"}
	// fmt.Println(unknown1)
	// fmt.Println(unknown2)

	// 通道数组
	// var chanList = [2]chan int{}
	// fmt.Println(chanList)

	// 空数组
	// var d [0]int
	// var e = [0]int
	// var f = [...]int{}

	// c1 := make(chan [0]int)
	// go func() {
	// 	fmt.Println("c1")
	// 	c1 <- [0]int{}
	// }()
	// <-c1   // 调用通道， 空数组不占空间作为通道返回不占开销

	// c2 := make(chan struct{})
	// go func (g string)  {
	// 	fmt.Println(g)
	// 	c2 <- struct{}{}
	// }("你好啊")
	// <-c2
}