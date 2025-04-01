package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	original := "hello, world!"
	titleCaser := cases.Title(language.English)
	title := titleCaser.String(original)
	fmt.Println("Original:", original)
	fmt.Println("Title Case:", title)
}
