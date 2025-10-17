package main

import (
	"fmt"
	"math"
)

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle type
type Circle struct {
	Radius float64
}

// Implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Implement Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that accepts any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Println("Rectangle:")
	printShapeInfo(rect)

	fmt.Println("\nCircle:")
	printShapeInfo(circle)
}
