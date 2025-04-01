package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := strings.SplitN("a,b,c,d", ",", 3)
	fmt.Printf("Parts: %q\n", parts)
}
