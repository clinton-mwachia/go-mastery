package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i) // Convert int to float64
	var u uint = uint(f)       // Convert float64 to uint

	fmt.Println(i, f, u)
}
