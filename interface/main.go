package main

// interface define
type Shape interface {
	area() float64
	perimeter() float64
}

// interface implementation
type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width * r.height) }
func main() {

}
