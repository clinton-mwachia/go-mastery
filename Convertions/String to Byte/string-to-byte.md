## 2. String and Byte Slice Conversions

Strings in Go are immutable sequences of bytes. Converting between strings and byte slices (`[]byte`) is a common operation, especially when dealing with binary data or performing low-level manipulations.

_Example:_

```go
package main

import "fmt"

func main() {
    s := "hello"
    b := []byte(s) // Convert string to []byte
    s2 := string(b) // Convert []byte back to string

    fmt.Println(s, b, s2)
}
```

Here, a string is converted to a byte slice and back to a string. This is useful when you need to modify the contents of a string since strings are immutable, but byte slices are mutable.
