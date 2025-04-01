## 4. String to Integer Conversion

Converting a string to an integer requires parsing the string and handling potential errors if the string does not represent a valid integer. The `strconv` package provides functions for this purpose.

_Example:_

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "123"
    i, err := strconv.Atoi(s) // Convert string to int
    if err != nil {
        fmt.Println("Conversion error:", err)
        return
    }
    fmt.Println(i)
}
```

In this code, `strconv.Atoi` is used to parse a string into an integer, with error handling to manage invalid input gracefully.
