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

// type switch

// var t interface{}
// t = functionOfSomeType()
// switch t := t.(type) {
// default:
//     fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
// case bool:
//     fmt.Printf("boolean %t\n", t)             // t has type bool
// case int:
//     fmt.Printf("integer %d\n", t)             // t has type int
// case *bool:
//     fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
// case *int:
//     fmt.Printf("pointer to integer %d\n", *t) // t has type *int
// }
