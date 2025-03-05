## Interfaces in Go

### Introduction

In Go, an **interface** is a type that specifies a set of method signatures. Any type that implements these methods is said to satisfy the interface, enabling polymorphism and allowing different types to be treated uniformly based on shared behavior. citeturn0search3

---

### Code Example

```go
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
```

---

### Explanation

**1. Defining an Interface**

An interface is defined using the `type` keyword, followed by the interface name and the `interface` keyword. Inside the braces, we list the method signatures that any type must implement to satisfy the interface. citeturn0search3

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

In this example, the `Shape` interface requires two methods: `Area()` and `Perimeter()`, both returning a `float64`.

**2. Implementing the Interface**

Any type that provides implementations for all methods specified in an interface is said to implement that interface. Notably, Go uses **implicit implementation**, meaning there's no need to explicitly declare that a type implements an interface. citeturn0search3

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

Here, the `Circle` struct implements both `Area()` and `Perimeter()`, thereby satisfying the `Shape` interface.

**3. Utilizing the Interface**

Interfaces allow functions to accept parameters of any type that satisfies the interface, enabling polymorphic behavior.

```go
func displayShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
```

The `displayShapeInfo` function can accept any type that implements the `Shape` interface, allowing for flexible and reusable code.

**4. Practical Example**

In the `main` function, we create instances of `Circle` and `Rectangle`, both of which implement the `Shape` interface. We then pass these instances to the `displayShapeInfo` function to display their area and perimeter.

```go
func main() {
    c := Circle{Radius: 5}
    fmt.Println("Circle:")
    displayShapeInfo(c)

    r := Rectangle{Width: 4, Height: 7}
    fmt.Println("\nRectangle:")
    displayShapeInfo(r)
}
```

This demonstrates how interfaces enable polymorphism, allowing different types to be treated uniformly based on shared behavior.

---

### Running the Code

1. Save the code in a file named `interfaces.go`.
2. Open a terminal and navigate to the directory containing `interfaces.go`.
3. Run the code using the command:

   ```sh
   go run interfaces.go
   ```

4. The expected output will be:

   ```
   Circle:
   Area: 78.54
   Perimeter: 31.42

   Rectangle:
   Area: 28.00
   Perimeter: 22.00
   ```
