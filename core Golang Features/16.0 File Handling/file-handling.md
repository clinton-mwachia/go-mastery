## File Handling

In Go, file handling is facilitated by the `os`, `io`, and `bufio` packages, enabling operations such as creating, reading, writing, and deleting files. Below is an overview of these operations, accompanied by code examples and documentation.

**1. Creating a File**

To create a new file, use the `os.Create` function. This function creates the file if it doesn't exist or truncates it if it does, returning an `os.File` object and an error.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Create a new file or truncate if it exists
    file, err := os.Create("example.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close() // Ensure the file is closed when the function exits
}
```

**2. Writing to a File**

After creating or opening a file, you can write to it using methods like `Write` or `WriteString`. Here's how to write a string to a file:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Open the file in write-only mode. If it doesn't exist, create it with permissions 0644.
    file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Write a string to the file
    _, err = file.WriteString("Hello, Gophers!\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
    fmt.Println("Data written successfully.")
}
```

**3. Reading from a File**

To read the entire content of a file into memory, you can use the `os.ReadFile` function:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Read the entire file content
    data, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("File content:")
    fmt.Println(string(data))
}
```

**4. Reading a File Line by Line**

For larger files, reading line by line is more efficient. This can be achieved using `bufio.Scanner`:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Open the file for reading
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a new scanner for the file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // Print each line
        fmt.Println(scanner.Text())
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
```

**5. Deleting a File**

To remove a file, use the `os.Remove` function:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Delete the file
    err := os.Remove("example.txt")
    if err != nil {
        fmt.Println("Error deleting file:", err)
        return
    }
    fmt.Println("File deleted successfully.")
}
```

**Documentation and Resources**

- The `os` package provides a platform-independent interface to operating system functionality, including file operations.
- The `bufio` package implements buffered I/O, offering efficient reading and writing.

For more detailed information, refer to the official Go documentation:

- [os package](https://pkg.go.dev/os)
- [bufio package](https://pkg.go.dev/bufio)

By leveraging these functions and packages, you can effectively manage file operations in your Go applications.
