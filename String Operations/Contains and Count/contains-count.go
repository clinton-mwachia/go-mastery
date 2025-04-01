package main

import (
	"fmt"
	"strings"
)

func main() {
	mainStr := "banana"
	subStr := "ana"
	char := "a"

	fmt.Printf("The substring \"%s\" appears %d times in \"%s\".\n", subStr, strings.Count(mainStr, subStr), mainStr)
	fmt.Printf("The character \"%s\" appears %d times in \"%s\".\n", char, strings.Count(mainStr, char), mainStr)
}
