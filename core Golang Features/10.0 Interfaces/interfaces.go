package main

import (
	"fmt"
	"math"
)

// Define the Shape interface with two method signatures
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a Circle struct
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Implement the Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Define a Rectangle struct
type Rectangle struct {
	Width, Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// A function that takes a Shape interface and displays its area and perimeter
func displayShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}
	fmt.Println("Circle:")
	displayShapeInfo(c)

	r := Rectangle{Width: 4, Height: 7}
	fmt.Println("\nRectangle:")
	displayShapeInfo(r)
}
