package main

/*
#include "hello.c"
*/
import "C"

func main() {
	C.cHello()
}
