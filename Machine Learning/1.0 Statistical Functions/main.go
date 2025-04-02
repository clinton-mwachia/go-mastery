package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"stats/utils"
	"strconv"
)

func main() {
	// Open the file containing the data
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a slice to store the data
	var data []float64

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert each line to a float64 and append to data slice
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Printf("Skipping invalid data: %v\n", line)
			continue
		}
		data = append(data, value)
	}

	// Check for errors during reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Float64s(data)

	fmt.Println("Min:", utils.Min(data))
	fmt.Println("Max:", utils.Max(data))
	fmt.Println("Mean:", utils.Mean(data))
	fmt.Println("Median:", utils.Median(data))
	fmt.Println("Variance:", utils.Variance(data))
	fmt.Println("Standard Deviation:", math.Sqrt(utils.Variance(data)))
}
