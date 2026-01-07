package main

/*
#include "hello.c"
*/
import "C"
import "fmt"

func main() {
	result := C.add(2, 6)
	fmt.Println("Result from C:", result)
}
