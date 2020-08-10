package main

import (
	"fmt"
)

func main()  {
	// var s = "hello, world"
	// var hello = s[:5]
	// var world = s[7:]
	// var s1 = "hello, world"[:5]
	// var s2 = "hello, world"[7:]
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(hello)
	// fmt.Println(len(hello))
	// fmt.Println(world)
	// fmt.Println(len(world))
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(s2)
	// fmt.Println(len(s2))
	var s = "你好世界"
	fmt.Println(len(s))
	fmt.Printf("%#v\n", []byte(s))
	for i,c := range s {
		fmt.Println(i,string(c))

	}
	// fmt.Printf("%#v\n", []rune("世界"))
	// fmt.Printf("%#v\n", string([]rune{'世', '界'}))
}