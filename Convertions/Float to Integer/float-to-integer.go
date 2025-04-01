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
