package circuit

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/interp"
)

func arange2(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N)
	for x := range rnge {
		rnge[x] = start + step*float64(x)
	}
	return rnge
}

func interpolate(x float64, xs, ys []float64) float64 {
	var as interp.AkimaSpline

	err := as.Fit(xs, ys)
	if err != nil {
		fmt.Println("Failed to initialize spline")
	}

	return as.Predict(x)
}

func RotateMtrx(matrix [][]float64) [][]interface{} {
	result := make([][]interface{}, len(matrix[0]))
	for i := 0; i < len(result); i++ {
		result[0] = make([]interface{}, 0)
	}

	for _, row := range matrix {
		for j, v := range row {
			result[j] = append(result[j], fmt.Sprintf("%.6g", v))
		}
	}

	return result
}
