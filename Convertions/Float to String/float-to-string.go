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
