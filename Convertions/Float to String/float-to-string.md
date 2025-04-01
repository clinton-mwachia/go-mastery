## Converting Float to String

To convert a floating-point number to a string, you can use either the `strconv.FormatFloat` function or the `fmt.Sprintf` function.

_Using `strconv.FormatFloat`:_

The `strconv.FormatFloat` function converts a `float64` to a string according to the specified format, precision, and bit size.

_Example:_

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    floatValue := 3.14159
    strValue := strconv.FormatFloat(floatValue, 'f', 2, 64)
    fmt.Println("Converted string value:", strValue)
}
```

In this example:

- The float value `3.14159` is converted to a string with two decimal places.
- The `'f'` format indicates decimal point notation.
- The precision `2` specifies two digits after the decimal point.
- The `bitSize` `64` corresponds to a `float64` value.

_Using `fmt.Sprintf`:_

Alternatively, `fmt.Sprintf` can format a float as a string with more flexibility.

_Example:_

```go
package main

import (
    "fmt"
)

func main() {
    floatValue := 3.14159
    strValue := fmt.Sprintf("%.2f", floatValue)
    fmt.Println("Converted string value:", strValue)
}
```

In this example:

- `fmt.Sprintf` formats the float to a string with two decimal places using the `%.2f` verb.

**Error Handling**

Proper error handling is crucial when parsing strings to floats, as invalid input can cause the conversion to fail. Always check the error returned by `strconv.ParseFloat` to ensure the conversion was successful.

_Example with Error Handling:_

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "invalid"
    floatValue, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("Conversion failed:", err)
        return
    }
    fmt.Printf("Converted float value: %f\n", floatValue)
}
```

In this case, attempting to parse the string `"invalid"` will result in an error, which is caught and handled gracefully.

By utilizing these functions and handling errors appropriately, you can effectively manage string and float conversions in your Go programs.
