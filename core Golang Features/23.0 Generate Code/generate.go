// package main

package main

import "fmt"

//go:generate stringer -type=Day

// Day represents days of the week.
type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)    // Output: 0
	fmt.Println(Wednesday) // Output: 3
}
