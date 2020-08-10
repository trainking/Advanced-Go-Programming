package main

import (
	"fmt"
)

// 指定为插入
func slice_insert(a []int, i int, x int) (r []int) {
	r = append(a, 0)
	copy(r[i+1:], r[i:])
	r[i] = x
	return r
}

// 指定位置批量插入
func slice_insert_any(a []int, i int, x []int)(r []int) {
	r = append(a, x...)
	copy(r[i + len(x):], a[i:])
	copy(r[i:], x)
	return r
}

// 原地删除空字符
func TrimSpace(s []byte) []byte {
	b := s[:0]    // 这样获取的是s的地址位置的一个空切片
	for _,x := range s {
		if x != ' '{
			b = append(b, x)
		}
	}
	return b
}

func main()  {
	// var a = []int{1,2,3,4}
	// a = append(a, 1)
	// a = append(a, 1,2,3)
	// a = append(a, []int{1,2,3}...)
	// a = append([]int{0}, a...)
	// a = append([]int{-3,-2,-1}, a...)
	// a = slice_insert(a,2,9)
	// a = slice_insert_any(a, 1,[]int{0,0,0})
	// fmt.Println(a)
	g := []byte{' ','a', ' ', 'c', ' ', 'g' }
	g = TrimSpace(g)
	fmt.Println(g)
}

