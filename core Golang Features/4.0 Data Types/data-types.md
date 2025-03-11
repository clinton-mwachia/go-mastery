# Data Types in Go

## Introduction

Go is a statically typed language, meaning every variable must have a defined type. Go provides several built-in data types, categorized into:

- **Basic types** (Numbers, Strings, Booleans)
- **Composite types** (Arrays, Slices, Maps, Structs, Pointers)
- **Reference types** (Interfaces, Channels, Functions)

---

## Code

```go
package main

import "fmt"

func main() {
    // Integer types
    var a int = 42
    var b int8 = 127
    var c int16 = 32767
    var d int32 = 2147483647
    var e int64 = 9223372036854775807

    // Unsigned integers
    var u uint = 42
    var u8 uint8 = 255
    var u16 uint16 = 65535
    var u32 uint32 = 4294967295
    var u64 uint64 = 18446744073709551615

    // Floating-point types
    var f32 float32 = 3.1415
    var f64 float64 = 2.718281828

    // Boolean type
    var isGoAwesome bool = true

    // String type
    var message string = "Hello, Golang!"

    // Rune (represents a Unicode character)
    var r rune = 'G'

    // Byte (alias for uint8, represents ASCII characters)
    var bChar byte = 'A'

    // Printing the values
    fmt.Println("Integer values:", a, b, c, d, e)
    fmt.Println("Unsigned integers:", u, u8, u16, u32, u64)
    fmt.Println("Floating point values:", f32, f64)
    fmt.Println("Boolean:", isGoAwesome)
    fmt.Println("String:", message)
    fmt.Println("Rune:", r)
    fmt.Println("Byte:", bChar)
}
```

## Explanation

### **1. Integer Types**

- `int` (default size depends on system architecture: 32-bit or 64-bit)
- `int8`, `int16`, `int32`, `int64` (signed integers of different sizes)
- `uint`, `uint8`, `uint16`, `uint32`, `uint64` (unsigned integers)

### **2. Floating-Point Types**

- `float32` (single precision, ~6-7 decimal places)
- `float64` (double precision, ~15-16 decimal places)

### **3. Boolean Type**

- `bool` holds `true` or `false`.

### **4. String Type**

- Strings are immutable sequences of characters.

### **5. Rune Type**

- `rune` represents a Unicode character (alias for `int32`).

### **6. Byte Type**

- `byte` is an alias for `uint8`, often used for ASCII characters.

---

## Running the Code

1. Save the file as `data_types.go`.
2. Open a terminal and navigate to the file's location.
3. Run:

   ```sh
   go run data_types.go
   ```

4. Expected output:

   ```
   Integer values: 42 127 32767 2147483647 9223372036854775807
   Unsigned integers: 42 255 65535 4294967295 18446744073709551615
   Floating point values: 3.1415 2.718281828
   Boolean: true
   String: Hello, Golang!
   Rune: 71
   Byte: 65
   ```

```

```
