// Package declaration: Required for every Go program.
package main

// Importing the fmt package for formatted output.
import "fmt"

func main() {
	// Declaring a variable using the 'var' keyword
	var name string = "Gopher"
	fmt.Println("Name:", name)

	// Declaring a variable without specifying the type (type inference)
	var age = 5
	fmt.Println("Age:", age)

	// Short variable declaration using ':='
	city := "Nairobi"
	fmt.Println("City:", city)

	// Declaring multiple variables at once
	var firstName, lastName string = "John", "Doe"
	fmt.Println("Full Name:", firstName, lastName)

	// Declaring a constant (value cannot change)
	const pi float64 = 3.1415
	fmt.Println("Value of Pi:", pi)

	// Constants with inferred type
	const country = "Kenya"
	fmt.Println("Country:", country)

	// Trying to modify a constant (Uncommenting this will cause an error)
	// pi = 3.14 // Error: cannot assign to pi
}
