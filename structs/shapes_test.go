package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape        Shape
		hasPerimeter float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, hasPerimeter: 40.0},
		{shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
		{shape: Triangle{SideA: 6, SideB: 6, SideC: 6}, hasPerimeter: 18.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
		}
	}

	// checkPerimeter := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Perimeter()
	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }

	// t.Run("rectangle perimeter", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkPerimeter(t, rectangle, 40.0)
	// })

	// t.Run("circle perimeter", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkPerimeter(t, circle, 62.83185307179586)
	// })
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }

	// t.Run("rectangle area", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkArea(t, rectangle, 100.0)
	// })

	// t.Run("circle area", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
