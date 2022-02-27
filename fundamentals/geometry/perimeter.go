package geometry

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	w, l float64
}

type Circle struct {
	r float64
}

type Triangle struct {
	base, height float64
}

func Perimeter(r Rectangle) float64 {
	return (r.w + r.l) * 2
}

func (r Rectangle) Area() float64 {
	return r.w * r.l
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.55
}
