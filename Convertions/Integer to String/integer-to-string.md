## 5. Integer to String Conversion

Converting an integer to a string can be accomplished using the `strconv` package's `Itoa` function.

_Example:_

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i := 123
    s := strconv.Itoa(i) // Convert int to string
    fmt.Println(s)
}
```

Here, `strconv.Itoa` converts an integer to its decimal string representation.
