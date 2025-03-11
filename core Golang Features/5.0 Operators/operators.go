package main

import "fmt"

func main() {
	// Arithmetic Operators
	a, b := 10, 5
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Addition:", a+b)       // 10 + 5 = 15
	fmt.Println("Subtraction:", a-b)    // 10 - 5 = 5
	fmt.Println("Multiplication:", a*b) // 10 * 5 = 50
	fmt.Println("Division:", a/b)       // 10 / 5 = 2
	fmt.Println("Modulus:", a%b)        // 10 % 5 = 0

	// Relational (Comparison) Operators
	fmt.Println("\nRelational Operators:")
	fmt.Println("a == b:", a == b) // false
	fmt.Println("a != b:", a != b) // true
	fmt.Println("a > b:", a > b)   // true
	fmt.Println("a < b:", a < b)   // false
	fmt.Println("a >= b:", a >= b) // true
	fmt.Println("a <= b:", a <= b) // false

	// Logical Operators
	x, y := true, false
	fmt.Println("\nLogical Operators:")
	fmt.Println("x && y (AND):", x && y) // false
	fmt.Println("x || y (OR):", x || y)  // true
	fmt.Println("!x (NOT):", !x)         // false

	// Bitwise Operators
	fmt.Println("\nBitwise Operators:")
	fmt.Println("a & b (AND):", a&b)           // 10 & 5 = 0
	fmt.Println("a | b (OR):", a|b)            // 10 | 5 = 15
	fmt.Println("a ^ b (XOR):", a^b)           // 10 ^ 5 = 15
	fmt.Println("a << 1 (Left Shift):", a<<1)  // 10 << 1 = 20
	fmt.Println("a >> 1 (Right Shift):", a>>1) // 10 >> 1 = 5

	// Assignment Operators
	fmt.Println("\nAssignment Operators:")
	c := 10
	fmt.Println("Initial value of c:", c)
	c += 5 // c = c + 5
	fmt.Println("c += 5:", c)
	c -= 3 // c = c - 3
	fmt.Println("c -= 3:", c)
	c *= 2 // c = c * 2
	fmt.Println("c *= 2:", c)
	c /= 4 // c = c / 4
	fmt.Println("c /= 4:", c)
	c %= 3 // c = c % 3
	fmt.Println("c %= 3:", c)

	// Unary Operators
	fmt.Println("\nUnary Operators:")
	num := 10
	fmt.Println("num:", num)
	num++ // Increment
	fmt.Println("num++:", num)
	num-- // Decrement
	fmt.Println("num--:", num)
}
