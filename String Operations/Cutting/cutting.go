package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "example.txt"
	suffix := ".txt"
	if before, found := strings.CutSuffix(str, suffix); found {
		fmt.Printf("String without suffix: %s\n", before)
	} else {
		fmt.Println("Suffix not found")
	}
}
