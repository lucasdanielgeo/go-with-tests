package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Widht  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Widht * r.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Widht)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
