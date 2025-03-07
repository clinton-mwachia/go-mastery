# go-mastery

Welcome to the **go-mastery**, an open-source guide covering the core features of the Go programming language. This repository serves as a structured learning resource for beginners and experienced developers.

## **🚀 Features**

- Covers Go's **fundamental concepts** and **advanced topics**
- Includes **well-organized examples** for each topic
- Open-source and **community-driven**

---

## Table of Contents

### Core Golang Features

1. [Introduction](#introduction)
2. [Basic Syntax](#basic-syntax)
3. [Variables and Constants](#variables-and-constants)
4. [Data Types](#data-types)
5. [Operators](#operators)
6. [Control Flow](#control-flow)
7. [Functions](#functions)
8. [Pointers](#pointers)
9. [Structs and Methods](#structs-and-methods)
10. [Interfaces](#interfaces)
11. [Arrays and Slices](#arrays-and-slices)
12. [Maps](#maps)
13. [Concurrency (Goroutines & Channels)](#concurrency-goroutines--channels)
14. [Error Handling](#error-handling)
15. [File Handling](#file-handling)
16. [Packages and Modules](#packages-and-modules)
17. [Testing in Go](#testing-in-go)
18. [Reflection](#reflection)
19. [Web Development (HTTP)](#web-development-http)
20. [Database Interaction](#database-interaction)
21. [Advanced Topics](#advanced-topics)

### Golang Data Structures

22. [Lists (Using Slices)](#lists-using-slices)
23. [Stacks](#stacks)
24. [Queues](#queues)
25. [Hash Tables (Maps)](#hash-tables-maps)
26. [Graphs](#graphs)
27. [Trees](#trees)
28. [Heaps](#heaps)
29. [Disjoint Set (Union-Find)](#disjoint-set-union-find)

### Golang Algorithms

30. [Searching Algorithms](#searching-algorithms)
31. [Sorting Algorithms](#sorting-algorithms)
32. [Graph Algorithms](#graph-algorithms)
33. [Dynamic Programming](#dynamic-programming)
34. [String Algorithms](#string-algorithms)
35. [Cryptography](#cryptography)

---

## Project Structure

```
go-mastery/
├── README.md
├── FUNDING.yml
|── LICENSE
|── .gitignore
├── docs/
```

## Introduction

Go is a statically typed, compiled high-level general purpose programming language.

### Install Go

- [Official Installation Guide](https://go.dev/dl/)
- Verify installation:
  ```sh
  go version
  ```

### Your First Go Program

Create a file `hello-world.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:

```sh
go run hello-world.go
```

## Basic Syntax

- **Package Declaration**: Every Go program starts with a `package` declaration.
- **Imports**: Use `import` to include standard or third-party packages.
- **Main Function**: The entry point for execution.

```go
package main

import "fmt"

func main() {
    fmt.Println("Go Syntax Example")
}
```

---

## Variables and Constants

- Declare variables using `var` or `:=` (short declaration).
- Constants are declared with `const`.

```go
package main

import "fmt"

func main() {
    var name string = "Alice"
    age := 25
    const pi = 3.14159

    fmt.Println(name, age, pi)
}
```

### Data Types

| Type      | Example                |
| --------- | ---------------------- |
| `int`     | `var x int = 42`       |
| `float64` | `var y float64 = 3.14` |
| `string`  | `var s string = "Go"`  |
| `bool`    | `var b bool = true`    |

---

## Operators

```go
package main

import "fmt"

func main() {
    // Arithmetic Operators
    a, b := 10, 5
    fmt.Println("Arithmetic Operators:")
    fmt.Println("Addition:", a+b)  // 10 + 5 = 15
    fmt.Println("Subtraction:", a-b)  // 10 - 5 = 5
    fmt.Println("Multiplication:", a*b)  // 10 * 5 = 50
    fmt.Println("Division:", a/b)  // 10 / 5 = 2
    fmt.Println("Modulus:", a%b)  // 10 % 5 = 0

    // Relational (Comparison) Operators
    fmt.Println("\nRelational Operators:")
    fmt.Println("a == b:", a == b)  // false
    fmt.Println("a != b:", a != b)  // true
    fmt.Println("a > b:", a > b)  // true
    fmt.Println("a < b:", a < b)  // false
    fmt.Println("a >= b:", a >= b)  // true
    fmt.Println("a <= b:", a <= b)  // false

    // Logical Operators
    x, y := true, false
    fmt.Println("\nLogical Operators:")
    fmt.Println("x && y (AND):", x && y)  // false
    fmt.Println("x || y (OR):", x || y)  // true
    fmt.Println("!x (NOT):", !x)  // false

    // Bitwise Operators
    fmt.Println("\nBitwise Operators:")
    fmt.Println("a & b (AND):", a & b)  // 10 & 5 = 0
    fmt.Println("a | b (OR):", a | b)  // 10 | 5 = 15
    fmt.Println("a ^ b (XOR):", a ^ b)  // 10 ^ 5 = 15
    fmt.Println("a << 1 (Left Shift):", a << 1)  // 10 << 1 = 20
    fmt.Println("a >> 1 (Right Shift):", a >> 1)  // 10 >> 1 = 5

    // Assignment Operators
    fmt.Println("\nAssignment Operators:")
    c := 10
    fmt.Println("Initial value of c:", c)
    c += 5  // c = c + 5
    fmt.Println("c += 5:", c)
    c -= 3  // c = c - 3
    fmt.Println("c -= 3:", c)
    c *= 2  // c = c * 2
    fmt.Println("c *= 2:", c)
    c /= 4  // c = c / 4
    fmt.Println("c /= 4:", c)
    c %= 3  // c = c % 3
    fmt.Println("c %= 3:", c)

    // Unary Operators
    fmt.Println("\nUnary Operators:")
    num := 10
    fmt.Println("num:", num)
    num++  // Increment
    fmt.Println("num++:", num)
    num--  // Decrement
    fmt.Println("num--:", num)
}
```

## Functions

- Functions are defined using `func` and can return multiple values.

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    fmt.Println(add(3, 5))
    fmt.Println(swap("hello", "world"))
}
```

---

## Pointers

Pointers in Go allow us to work with memory addresses, enabling efficient function calls, avoiding unnecessary copying of large data structures, and modifying variables directly.

```go
  // Declaring a pointer
    var ptr *int
    num := 10
    ptr = &num // Storing the address of num

    fmt.Println("Value of num:", num)
    fmt.Println("Pointer Address:", ptr)
    fmt.Println("Value at Pointer Address:", *ptr)
```

## Structs and Methods

### Structs

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Bob", Age: 30}
    fmt.Println(p.Name, p.Age)
}
```

### Methods

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hello " + p.Name
}

func main() {
	person := Person{Name: "James"}
	greet := person.Greet()
	fmt.Println(greet)
}
```

### Interfaces

```go
package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

func main() {
    var a Animal = Dog{}
    fmt.Println(a.Speak())
}
```

---

## Arrays and Slices

```go
// Declaring an array of 5 integers
	var arr [5]int
	fmt.Println("Default array:", arr) // [0 0 0 0 0]

	// Initializing an array with specific values
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", arr2) // [1 2 3 4 5]

	// Array with inferred length
	arr3 := [...]int{10, 20, 30}
	fmt.Println("Array with inferred length:", arr3) // [10 20 30]

    // Creating a slice using a slice literal
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice) // [1 2 3 4 5]

	// Creating a slice from an array
	arr_slice := [5]int{10, 20, 30, 40, 50}
	slice2 := arr_slice[1:4]
	fmt.Println("Slice from array:", slice2) // [20 30 40]

```

## Maps

```go
package main

import "fmt"

func main() {
    // Initialize a map
    m := make(map[string]int)

    // Add key-value pairs
    m["apple"] = 2
    m["banana"] = 5

    // Update a value
    m["apple"] = 4

    fmt.Println(m) // Output: map[apple:4 banana:5]
}
```

## **📖 6. Concurrency**

### **🔹 Goroutines**

Run functions concurrently using `go` keyword:

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from Goroutine")
}

func main() {
    go sayHello()
    time.Sleep(time.Second)
}
```

### **🔹 Channels (Communication)**

```go
package main

import "fmt"

func main() {
    messages := make(chan string)
    go func() { messages <- "Hello, Go!" }()
    fmt.Println(<-messages)
}
```

---

## **📖 7. Error Handling**

### **🔹 Using `error` Interface**

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

---

## **📖 8. Packages & Modules**

- Create a module:

  ```sh
  go mod init example.com/mymodule
  ```

- Importing your module:

  ```go
  package main

  import (
      "fmt"
      "example.com/mymodule/mypackage"
  )

  func main() {
      fmt.Println(mypackage.SayHello())
  }
  ```

---

## **📖 9. Testing**

- Write unit tests in `_test.go` files.

```go
package main

import "testing"

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

- Run tests:
  ```sh
  go test
  ```

---

## **📖 10. Advanced Topics**

- Reflection
- Embedding
- WebAssembly
- Using `cgo` to interact with C code

---

## **❤️ 12. Support the Project**

If you find this reference useful, consider supporting via **GitHub Sponsors**:

[![Sponsor](https://img.shields.io/badge/Sponsor-GitHub%20Sponsors-pink.svg)](https://github.com/sponsors/clinton-mwachia)

---

## **🚀 Happy Coding in Go!**
