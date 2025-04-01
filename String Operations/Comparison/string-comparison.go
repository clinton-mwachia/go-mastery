package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "GoLang"
	str2 := "golang"

	if strings.EqualFold(str1, str2) {
		fmt.Println("str1 and str2 are equal (case-insensitive).")
	} else {
		fmt.Println("str1 and str2 are not equal (case-insensitive).")
	}
}
