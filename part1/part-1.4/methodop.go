package main


import (
	"fmt"
	"image/color"
	"sync"
)

// 组合方式
type Point struct {X, Y float64}
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// 嵌入匿名成员方法
type Cache struct {
	m map[string]string   // 定义map类型
	sync.Mutex
}

func (p *Cache) Lookup(key string) string {
	p.Lock()
	defer p.Unlock()
	return p.m[key]
}

func (p *Cache) fix(key string, val string) {
	p.Lock()
	defer p.Unlock()
	p.m[key] = val
}

func main() {

	// 达到类似继承的效果
	// var cp ColoredPoint
	// cp.X=1
	// fmt.Println(cp.Point.X)
	// cp.Point.Y = 2
	// fmt.Println(cp.Y)
	var cache Cache
	cache.m = map[string]string {"ggggg": "12#$"}
	s := cache.Lookup("ggggg")
	fmt.Println(s)
	cache.fix("ggggg", "23423")
	fmt.Println(cache.m)
}