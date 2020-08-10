package main

import (
	"fmt"
	"io"
	"bytes"
	"os"
	"strings"
)

type UpperWriter struct {
	m string
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int,err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

type UpperString string
func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main()  {
	// 通过& 的方式定义结构实例
	// fmt.Fprintln(&UpperWriter{"ggg", os.Stdout}, "hello, world")	
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
}