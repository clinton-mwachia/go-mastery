package main

import "fmt"

func main() {
	s := "hello"
	b := []byte(s)  // Convert string to []byte
	s2 := string(b) // Convert []byte back to string

	fmt.Println(s, b, s2)
}
