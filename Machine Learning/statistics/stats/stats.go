package stats

import (
	"errors"
	"math"
	"sort"
)

// Min returns the minimum value in nums.
func Min(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	min := nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// Max returns the maximum value in nums.
func Max(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	max := nums[0]
	for _, v := range nums[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// Sum returns the total sum of nums.
func Sum(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	total := 0.0
	for _, v := range nums {
		total += v
	}
	return total, nil
}

// Mean returns the arithmetic mean (average).
func Mean(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	sum, _ := Sum(nums)
	return sum / float64(len(nums)), nil
}

// Median returns the median of nums.
func Median(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}

	// Avoid modifying original slice
	n := len(nums)
	sorted := make([]float64, n)
	copy(sorted, nums)
	sort.Float64s(sorted)

	// odd
	if n%2 == 1 {
		return sorted[n/2], nil
	}

	// even
	m1 := sorted[(n/2)-1]
	m2 := sorted[n/2]
	return (m1 + m2) / 2.0, nil
}

// Variance returns sample variance.
func Variance(nums []float64) (float64, error) {
	if len(nums) < 2 {
		return 0, errors.New("need at least 2 values")
	}

	mean, _ := Mean(nums)
	sum := 0.0

	for _, v := range nums {
		diff := v - mean
		sum += diff * diff
	}

	return sum / float64(len(nums)-1), nil
}

// StdDev returns the sample standard deviation.
func StdDev(nums []float64) (float64, error) {
	variance, err := Variance(nums)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(variance), nil
}
