## 3. String and Rune Slice Conversions

A rune in Go represents a Unicode code point and is an alias for `int32`. Converting a string to a slice of runes allows for manipulation of individual Unicode characters, which is particularly useful when dealing with multi-byte characters.

_Example:_

```go
package main

import "fmt"

func main() {
    s := "こんにちは" // "Hello" in Japanese
    r := []rune(s) // Convert string to []rune
    s2 := string(r) // Convert []rune back to string

    fmt.Println(s, r, s2)
}
```

This example demonstrates converting a string containing multi-byte characters to a rune slice and back, facilitating operations on individual characters.
