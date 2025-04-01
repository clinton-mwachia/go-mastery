package main

import (
	"fmt"
	"strings"
)

func main() {
	items := []string{"apple", "banana", "cherry"}
	list := strings.Join(items, ", ")
	fmt.Println(list) // Output: apple, banana, cherry
}
