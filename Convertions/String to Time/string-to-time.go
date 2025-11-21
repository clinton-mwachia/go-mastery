package main

import (
	"fmt"
	"time"
)

func main() {
	parsedTime, err := time.Parse("02-01-2006 15:04:05", time.Now().Format("02-01-2006 15:04:05"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedTime)
}
