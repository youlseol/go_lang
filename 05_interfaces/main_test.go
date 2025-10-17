package main

import (
	"math"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 3}
	expected := 15.0
	result := rect.Area()
	if result != expected {
		t.Errorf("Rectangle.Area() = %f; want %f", result, expected)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	rect := Rectangle{Width: 5, Height: 3}
	expected := 16.0
	result := rect.Perimeter()
	if result != expected {
		t.Errorf("Rectangle.Perimeter() = %f; want %f", result, expected)
	}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 4}
	expected := math.Pi * 16
	result := circle.Area()
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("Circle.Area() = %f; want %f", result, expected)
	}
}

func TestCirclePerimeter(t *testing.T) {
	circle := Circle{Radius: 4}
	expected := 2 * math.Pi * 4
	result := circle.Perimeter()
	if math.Abs(result-expected) > 0.0001 {
		t.Errorf("Circle.Perimeter() = %f; want %f", result, expected)
	}
}

func TestShapeInterface(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
	}

	for _, shape := range shapes {
		// Just verify that the interface methods can be called
		area := shape.Area()
		perimeter := shape.Perimeter()
		if area <= 0 || perimeter <= 0 {
			t.Errorf("Shape methods returned invalid values: Area=%f, Perimeter=%f", area, perimeter)
		}
	}
}
