package main

import (
	"fmt"
	"math"
)

// Define a struct for a Circle
type Circle struct {
	Radius float64
}

// Define a struct for a Rectangle
type Rectangle struct {
	Width, Height float64
}

// Define a struct for a Person
type Person struct {
	Name string
}

// Method to calculate the area of a Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Method to calculate the perimeter of a Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Method to calculate the area of a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method to calculate the perimeter of a Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// greet the person
func (p Person) Greet() string {
	return "Hello " + p.Name
}

func main() {
	c := Circle{Radius: 5}
	fmt.Printf("Circle: Radius = %.2f\n", c.Radius)
	fmt.Printf("Area = %.2f\n", c.Area())
	fmt.Printf("Perimeter = %.2f\n\n", c.Perimeter())

	r := Rectangle{Width: 4, Height: 7}
	fmt.Printf("Rectangle: Width = %.2f, Height = %.2f\n", r.Width, r.Height)
	fmt.Printf("Area = %.2f\n", r.Area())
	fmt.Printf("Perimeter = %.2f\n\n", r.Perimeter())

	p := Person{Name: "Clinton Moshe"}
	fmt.Printf("Greetings: %s", p.Greet())
}
