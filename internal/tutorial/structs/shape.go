package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Circle) Perimeter() float64 {
	return 2 * r.Radius * math.Pi
}

func (r Circle) Area() float64 {
	return r.Radius * r.Radius * math.Pi
}

func (r Triangle) Area() float64 {
	return r.Base * r.Height / 2
}
