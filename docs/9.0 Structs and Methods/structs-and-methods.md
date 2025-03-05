## Structs and Methods in Go

### Introduction

In Go, a **struct** is a composite data type that groups together variables under a single name, allowing you to create complex data structures. A **method** is a function with a special receiver argument, enabling you to define behavior associated with a particular type, typically a struct. citeturn0search4

---

### Code Example

```go
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

func main() {
    c := Circle{Radius: 5}
    fmt.Printf("Circle: Radius = %.2f\n", c.Radius)
    fmt.Printf("Area = %.2f\n", c.Area())
    fmt.Printf("Perimeter = %.2f\n\n", c.Perimeter())

    r := Rectangle{Width: 4, Height: 7}
    fmt.Printf("Rectangle: Width = %.2f, Height = %.2f\n", r.Width, r.Height)
    fmt.Printf("Area = %.2f\n", r.Area())
    fmt.Printf("Perimeter = %.2f\n", r.Perimeter())
}
```

---

### Explanation

**1. Defining Structs**

A struct is defined using the `type` keyword followed by the struct name and the `struct` keyword. citeturn0search6

```go
type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width, Height float64
}

type Person struct {
	Name string
}
```

In these examples, `Circle`, `Person` and `Rectangle` are structs with fields that describe their properties.

**2. Defining Methods**

Methods are functions associated with a specific type. They are defined with a receiver argument, which appears between the `func` keyword and the method name. citeturn0search4

```go
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (p Person) Greet() string {
	return "Hello " + p.Name
}
```

Here, `Area` methods are associated with `Circle` and `Rectangle` types, respectively. The receiver `(c Circle)` and `(r Rectangle)` indicates that these methods belong to their respective structs. `Greet` method is associated with Person.

**3. Using Methods**

Methods are called using the dot notation on an instance of the struct.

```go
c := Circle{Radius: 5}
fmt.Println(c.Area())

r := Rectangle{Width: 4, Height: 7}
fmt.Println(r.Area())

p := Person{Name: "Clinton Moshe"}
fmt.Printf("Greetings: %s", p.Greet())
```

**4. Pointer Receivers vs. Value Receivers**

Methods can have either value or pointer receivers. A method with a value receiver operates on a copy of the value, while a method with a pointer receiver can modify the original value. citeturn0search0

```go
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}
```

Using a pointer receiver allows the `Scale` method to modify the `Radius` of the original `Circle` instance.

---

### Running the Code

1. Save the code in a file named `structs-and-methods.go`.
2. Open a terminal and navigate to the directory containing `structs-and-methods.go`.
3. Run the code using the command:

   ```sh
   go run structs-and-methods.go
   ```

4. The expected output will be:

   ```go
   Circle: Radius = 5.00
   Area = 78.54
   Perimeter = 31.42

   Rectangle: Width = 4.00, Height = 7.00
   Area = 28.00
   Perimeter = 22.

   Greetings: Hello Clinton Moshe
   ```
