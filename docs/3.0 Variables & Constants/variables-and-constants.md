# Variables and Constants in Go

## Introduction

Variables and constants are fundamental components of any programming language. In Go, variables store values that can change, while constants hold values that remain fixed throughout the program.

---

## Code

```go
// Package declaration: Required for every Go program.
package main

// Importing the fmt package for formatted output.
import "fmt"

func main() {
    // Declaring a variable using the 'var' keyword
    var name string = "Gopher"
    fmt.Println("Name:", name)

    // Declaring a variable without specifying the type (type inference)
    var age = 5
    fmt.Println("Age:", age)

    // Short variable declaration using ':=' (only inside functions)
    city := "Nairobi"
    fmt.Println("City:", city)

    // Declaring multiple variables at once
    var firstName, lastName string = "John", "Doe"
    fmt.Println("Full Name:", firstName, lastName)

    // Declaring a constant (value cannot change)
    const pi float64 = 3.1415
    fmt.Println("Value of Pi:", pi)

    // Constants with inferred type
    const country = "Kenya"
    fmt.Println("Country:", country)

    // Trying to modify a constant (Uncommenting this will cause an error)
    // pi = 3.14  // Error: cannot assign to pi
}
```

## Explanation

### **1. Variables in Go**

- Variables store data that can be modified during execution.
- Declared using the `var` keyword or the `:=` shorthand (inside functions only).
- Example:
  ```go
  var x int = 10
  y := 20
  ```
- Variables can be initialized without specifying a type (Go infers it).

### **2. Constants in Go**

- Constants are immutable (cannot be changed after declaration).
- Declared using the `const` keyword.
- Example:
  ```go
  const pi = 3.1415
  ```
- Unlike variables, constants **must** be assigned a value at declaration.

### **3. Multiple Variable Declaration**

- Variables can be declared in a single statement:
  ```go
  var a, b, c int = 1, 2, 3
  ```
- This improves readability and organizes code better.

---

## Running the Code

1. Save the file as `variables_and_constants.go`.
2. Open a terminal and navigate to the file's location.
3. Run:

   ```sh
   go run variables_and_constants.go
   ```

4. Expected output:

   ```
   Name: Gopher
   Age: 5
   City: Nairobi
   Full Name: John Doe
   Value of Pi: 3.1415
   Country: Kenya
   ```

---
