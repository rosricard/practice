package interfaces

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rect struct {
	Width  float64
	Height float64
}

func (r *Rect) Area() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func getArea(s Shape) float64 {
	return s.Area()
}

// TODO: implement a type switch
