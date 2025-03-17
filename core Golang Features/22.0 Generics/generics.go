package main

import "fmt"

// Number is a constraint that permits only int and float64.
type Number interface {
	~int | ~float64
}

// Sum calculates the sum of a slice of numbers.
func Sum[T Number](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	intSlice := []int{1, 2, 3}
	fmt.Println(Sum(intSlice)) // Output: 6

	floatSlice := []float64{1.1, 2.2, 3.3}
	fmt.Println(Sum(floatSlice)) // Output: 6.6
}
