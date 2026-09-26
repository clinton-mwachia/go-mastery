package main

/*
#include "math.h"
*/
import "C"
import "fmt"

func main() {
	addResult := C.add(2, 6)
	fmt.Println("Add Result from C:", addResult)

	subResult := C.sub(2, 6)
	fmt.Println("Sub Result from C:", subResult)
}
