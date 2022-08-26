package shapes

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

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
	Side1  float64
	Side2  float64
	Side3  float64
}

// Perimeter returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area returns the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of a triangle.
func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

// Area returns the area of a triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
