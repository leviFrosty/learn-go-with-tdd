package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
