package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go gopher, go!"
	substr := "go"
	index := strings.LastIndex(str, substr)
	if index != -1 {
		fmt.Printf("The last occurrence of %q starts at index %d.\n", substr, index)
	} else {
		fmt.Printf("%q not found in %q.\n", substr, str)
	}
}
