package utils

// calculate the minimum value
func Min(x []float64) float64 {
	return x[0]
}

// calculate the maximum value
func Max(x []float64) float64 {
	return x[len(x)-1]
}

// calculate the mean value
func Mean(x []float64) float64 {
	sum := float64(0)
	for _, v := range x {
		sum = sum + v
	}
	return sum / float64(len(x))
}

// calculate the Median value
func Median(x []float64) float64 {
	length := len(x)
	if length%2 == 1 {
		// Odd
		return x[(length-1)/2]
	} else {
		// Even
		return (x[length/2] + x[(length/2)-1]) / 2
	}
}

// calculate the variance
func Variance(x []float64) float64 {
	mean := Mean(x)
	sum := float64(0)
	for _, v := range x {
		sum = sum + (v-mean)*(v-mean)
	}
	return sum / float64(len(x))
}
