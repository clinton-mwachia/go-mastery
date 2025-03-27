package main

import (
	"fmt"
	"strings"
)

func main() {
	original := "Hello, Gophers!"
	cloned := strings.Clone(original)

	fmt.Println("Original String:", original)
	fmt.Println("Cloned String:", cloned)
}
