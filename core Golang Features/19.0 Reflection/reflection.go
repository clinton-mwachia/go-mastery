package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 30}
	v := reflect.ValueOf(p)

	for i := range v.NumField() {
		fmt.Printf("Field %d: %v\n", i, v.Field(i))
	}
}
