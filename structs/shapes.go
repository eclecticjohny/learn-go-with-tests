package main

import (
	"fmt"
	"math"
)

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
	SideA  float64
	SideB  float64
	SideC  float64
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
	return t.SideA + t.SideB + t.SideC
}

// Area returns the area of a triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle Perimeter: ", r.Perimeter())
	fmt.Println("Rectangle Area: ", r.Area())

	c := Circle{Radius: 5}
	fmt.Println("Circle Perimeter: ", c.Perimeter())
	fmt.Println("Circle Area: ", c.Area())

	t := Triangle{Base: 3, Height: 4, SideA: 3, SideB: 4, SideC: 5}
	fmt.Println("Triangle Perimeter: ", t.Perimeter())
	fmt.Println("Triangle Area: ", t.Area())
}
