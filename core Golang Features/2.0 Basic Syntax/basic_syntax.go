// Package declaration: Every Go program must belong to a package.
// The "main" package is the entry point for an executable program.
package main

// Import statement: Used to include standard and custom packages.
import "fmt"

// The main function: The execution starts from the main function in the main package.
func main() {
	// Printing to the console using the fmt package.
	fmt.Println("Welcome to Go!")

	// Declaring a variable using the 'var' keyword.
	var message string = "Go is awesome!"
	fmt.Println(message)

	// Declaring a variable using the shorthand ':=' keyword.
	message2 := "Go is awesome too!"
	fmt.Println(message2)
}
