package solvers

import "fmt"

func InitX(x0, h float64, n int) []float64 {
	x := make([]float64, 0)
	for i := 0; i <= n; i++ {
		x = append(x, x0)
		x0 += h
	}
	return x
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
