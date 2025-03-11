### **📄 hello-world**

````md
# Hello, World! in Go

## Introduction

"Hello, World!" is the simplest program that demonstrates the basic syntax of the Go programming language. It prints "Hello, World!" to the console.

## Code

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
````

## Explanation

- `package main`: Defines the main package (the entry point for execution).
- `import "fmt"`: Imports the `fmt` package, which provides I/O functions.
- `func main() {}`: The `main` function is the program's starting point.
- `fmt.Println("Hello, World!")`: Prints "Hello, World!" to the console.

## How to Run the Program

1. Open the terminal and navigate to _docs/1.0 Introduction_.

```
cd "docs/1.0 Introduction"
```

2. Run the following command:

   ```sh
   go run hello-world.go
   ```

3. You should see the output:

   ```
   Hello, World!
   ```
