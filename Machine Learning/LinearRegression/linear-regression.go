package LinearRegression

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat"
)

func LinearRegression() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}

	slope, intercept := stat.LinearRegression(x, y, nil, false)

	fmt.Println("Slope:", slope)
	fmt.Println("Intercept:", intercept)

	// Predict y values
	yPred := make([]float64, len(x))
	for i, xi := range x {
		yPred[i] = slope*xi + intercept
	}

	// Calculate R² (coefficient of determination)
	r2 := stat.RSquaredFrom(y, yPred, nil)
	fmt.Printf("R²: %.4f\n", r2)

	// Calculate RMSE (Root Mean Squared Error)
	sumSqErr := 0.0
	for i := range y {
		err := y[i] - yPred[i]
		sumSqErr += err * err
	}
	rmse := math.Sqrt(sumSqErr / float64(len(y)))
	fmt.Printf("RMSE: %.4f\n", rmse)

	// Predict y for a new x value
	xNew := 6.0
	yNew := slope*xNew + intercept
	fmt.Printf("Predicted y for x=%.2f: %.2f\n", xNew, yNew)

}
