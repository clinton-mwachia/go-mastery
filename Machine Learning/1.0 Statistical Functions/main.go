package main

import (
	"fmt"
	"math"
	"sort"
	"stats/utils"
)

func main() {
	data := []float64{2.0, 1.0, 3.0}
	sort.Float64s(data)
	fmt.Println("Min:", utils.Min(data))
	fmt.Println("Max:", utils.Max(data))
	fmt.Println("Mean:", utils.Mean(data))
	fmt.Println("Median:", utils.Median(data))
	fmt.Println("Variance:", utils.Variance(data))
	fmt.Println("Standard Deviation:", math.Sqrt(utils.Variance(data)))
}
