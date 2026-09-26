package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 18 {
		return errors.New("age must be 18 or older")
	}
	return nil
}

func main() {
	age := 16
	err := validateAge(age)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age: ", age)
	}
}
