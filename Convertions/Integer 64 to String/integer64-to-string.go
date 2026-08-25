package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64 = 123456789

	str := strconv.FormatInt(num, 10)

	fmt.Println(str)
}
