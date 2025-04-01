## Numeric Convertions

In Go, type conversion is the process of converting a value from one data type to another. Due to Go's statically typed nature, such conversions must be explicit, ensuring type safety and clarity in code. Below is an overview of common type conversions in Go:

**1. Numeric Conversions:**

Go supports explicit conversion between different numeric types, including integers and floating-point numbers. This is essential when performing operations that require specific numeric types.

_Example:_

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i) // Convert int to float64
    var u uint = uint(f)       // Convert float64 to uint

    fmt.Println(i, f, u)
}
```

In this example, an `int` is converted to a `float64`, and then to a `uint`. Each conversion is explicitly specified to ensure type compatibility.
