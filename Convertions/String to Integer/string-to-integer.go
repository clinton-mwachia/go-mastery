package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123"
	i, err := strconv.Atoi(s) // Convert string to int
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}
	fmt.Println(i)
}
