package main

import (
	_ "embed"
	"fmt"
)

//go:embed config.yaml
var configData string

func main() {
	fmt.Println(configData)
}
