package main

import (
	"fmt"
	"strconv"
)

func main() {
	float := 20.22
	stringValue := strconv.FormatFloat(float, 'f', -1, 64)

	fmt.Printf("Converted float value: %s\n", stringValue)
}
