package main

import "fmt"

// http://golang.site/go/article/17-Go-%EB%A9%94%EC%84%9C%EB%93%9C

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{30, 2}
	area := rect.area()
	area2 := rect.area2()
	area3 := rect.area()
	fmt.Println(area)
	fmt.Println(area2)
	fmt.Println(area3)
}
