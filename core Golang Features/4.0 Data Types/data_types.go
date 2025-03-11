package main

import "fmt"

func main() {
	// Integer types
	var a int = 42
	var b int8 = 127
	var c int16 = 32767
	var d int32 = 2147483647
	var e int64 = 9223372036854775807

	// Unsigned integers
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	// Floating-point types
	var f32 float32 = 3.1415
	var f64 float64 = 2.718281828

	// Boolean type
	var isGoAwesome bool = true

	// String type
	var message string = "Hello, Golang!"

	// Rune (represents a Unicode character)
	var r rune = 'G'

	// Byte (alias for uint8, represents ASCII characters)
	var bChar byte = 'A'

	// Printing the values
	fmt.Println("Integer values:", a, b, c, d, e)
	fmt.Println("Unsigned integers:", u, u8, u16, u32, u64)
	fmt.Println("Floating point values:", f32, f64)
	fmt.Println("Boolean:", isGoAwesome)
	fmt.Println("String:", message)
	fmt.Println("Rune:", r)
	fmt.Println("Byte:", bChar)
}
