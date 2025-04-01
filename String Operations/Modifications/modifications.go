package main

import (
	"fmt"
	"strings"
)

func main() {
	original := "oink oink oink"
	modified := strings.ReplaceAll(original, "oink", "moo")
	fmt.Println(modified) // Output: "moo moo moo"
}
