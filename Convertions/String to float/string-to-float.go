package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "3.14159"
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
		return
	}
	fmt.Printf("Converted float value: %f\n", floatValue)
}
