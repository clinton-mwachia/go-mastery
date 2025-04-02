package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

// y = ax + b

type xy struct {
	x []float64
	y []float64
}

func main() {
	// Open the file containing the data
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	size := len(records)

	// Create a slice to store the data
	data := xy{
		x: make([]float64, size),
		y: make([]float64, size),
	}

	for i, v := range records {
		if len(v) != 2 {
			fmt.Println("Expected two elements")
			continue
		}
		if s, err := strconv.ParseFloat(v[0], 64); err == nil {
			data.y[i] = s
		}
		if s, err := strconv.ParseFloat(v[1], 64); err == nil {
			data.x[i] = s
		}
	}
	b, a := stat.LinearRegression(data.x, data.y, nil, false)
	fmt.Printf("%.4v x + %.4v\n", a, b)
	fmt.Printf("a = %.4v b = %.4v\n", a, b)
}
