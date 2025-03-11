package main

import "fmt"

func main() {
	// If-Else Statement
	age := 18
	fmt.Println("If-Else Statement:")
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// If with Short Statement
	fmt.Println("\nIf with Short Statement:")
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// Switch Statement
	fmt.Println("\nSwitch Statement:")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// Switch without Condition (Equivalent to Multiple If-Else)
	fmt.Println("\nSwitch without Condition:")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// For Loop
	fmt.Println("\nFor Loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// For Loop with range
	fmt.Println("\nFor Loop with range:")
	nums := []int{2, 3, 4}
	for i, num := range nums {
		fmt.Println("Index:", i, "Num: ", num)
	}

	// While-like Loop
	fmt.Println("\nWhile-like Loop:")
	j := 1
	for j <= 5 {
		fmt.Println("Value of j:", j)
		j++
	}

	// Infinite Loop with Break
	fmt.Println("\nInfinite Loop with Break:")
	k := 1
	for {
		if k > 3 {
			break
		}
		fmt.Println("k:", k)
		k++
	}

	// Loop with Continue
	fmt.Println("\nLoop with Continue:")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Skip iteration when i == 3
		}
		fmt.Println("i:", i)
	}

	// Defer Statement
	fmt.Println("\nDefer Statement:")
	defer fmt.Println("This will execute at the end of main function.")
	fmt.Println("Hello, World!")

	// Panic and Recover
	fmt.Println("\nPanic and Recover:")
	safeFunction()
}

// Function demonstrating panic and recover
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Executing safeFunction...")
	panic("Something went wrong!") // Causes a panic
}
