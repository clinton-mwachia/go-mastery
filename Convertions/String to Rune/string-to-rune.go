package main

import "fmt"

func main() {
	s := "こんにちは"    // "Hello" in Japanese
	r := []rune(s)  // Convert string to []rune
	s2 := string(r) // Convert []rune back to string

	fmt.Println(s, r, s2)
}
