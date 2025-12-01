package main

import (
	"fmt"
	"go-mastery/statistics/stats"
)

func main() {
	values := []float64{12, 5, 18, 7, 9, 13}

	min, _ := stats.Min(values)
	max, _ := stats.Max(values)
	mean, _ := stats.Mean(values)
	median, _ := stats.Median(values)
	variance, _ := stats.Variance(values)
	stddev, _ := stats.StdDev(values)

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	fmt.Println("Mean:", mean)
	fmt.Println("Median:", median)
	fmt.Println("Variance:", variance)
	fmt.Println("StdDev:", stddev)
}
