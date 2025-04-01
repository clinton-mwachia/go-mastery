package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s, ok := i.(string)
	if ok {
		fmt.Println(s) // Successfully asserted i holds a string
	} else {
		fmt.Println("i does not hold a string")
	}
}
