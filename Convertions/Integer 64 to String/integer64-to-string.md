## 5. Integer64 to String Conversion

Converting an integer64 to a string can be accomplished using the `strconv` package's `FormatInt` function.

_Example:_

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
	var num int64 = 123456789

	str := strconv.FormatInt(num, 10)

	fmt.Println(str)
}
```

Here, `strconv.FormatInt` converts an integer 64 to its string.
