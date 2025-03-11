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
	err := validateAge(16)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
