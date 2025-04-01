## 1. Converting Float to Integer

When converting a `float64` to an `int`, Go truncates the decimal portion, effectively rounding towards zero. This means that both positive and negative numbers are rounded towards zero, which may not be the desired behavior if standard rounding is required.

_Example:_

```go
package main

import (
    "fmt"
)

func main() {
    f := 3.7
    i := int(f) // Truncates to 3
    fmt.Println(i)

    f = -3.7
    i = int(f) // Truncates to -3
    fmt.Println(i)
}
```

_Output:_

```
3
-3
```

_Rounding to Nearest Integer:_

To round a float to the nearest integer before converting to an `int`, use the `math.Round` function introduced in Go 1.10. This function rounds ties away from zero.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    f := 3.7
    i := int(math.Round(f)) // Rounds to 4
    fmt.Println(i)

    f = -3.7
    i = int(math.Round(f)) // Rounds to -4
    fmt.Println(i)
}
```

_Output:_

```
4
-4
```

_Note:_ When converting a floating-point number to an integer, if the float's value exceeds the range of the target integer type, the result is implementation-dependent and may lead to unexpected behavior. Always ensure that the float value is within the bounds of the target integer type to avoid overflow issues.
