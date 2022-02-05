package main

import (
	"fmt"
	logger "lab_01/logger"
	solvers "lab_01/solver"
	"math"
)

func initX(x0, h float64, n int) []float64 {
	x := make([]float64, 0)
	for i := 0; i <= n; i++ {
		x = append(x, x0)
		x0 += h
	}
	return x
}

func inputArgs() (float64, float64, float64, error) {
	fmt.Println("Please, enter range for calculation (from and to):")

	var from, to float64
	if n, err := fmt.Scanf("%f %f", &from, &to); n != 2 || err != nil {
		return 0, 0, 0, err
	}
	fmt.Println()

	fmt.Println("Please, enter step for calculation:")
	var h float64
	if n, err := fmt.Scanf("%f", &h); n != 1 || err != nil {
		return 0, 0, 0, err
	}
	fmt.Println()

	return from, to, h, nil
}

func main() {
	from, to, h, err := inputArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	ys := 0.

	n := int(math.Ceil(math.Abs(to-from) / h))

	solutionMatrix := make([][]float64, 0)
	solutionMatrix = append(solutionMatrix, initX(from, h, n))

	solutionMatrix = append(solutionMatrix, solvers.NewPicard().Solution(from, h, n)...)
	solutionMatrix = append(solutionMatrix, solvers.NewEuler().Solution(from, ys, h, n))
	solutionMatrix = append(solutionMatrix, solvers.NewEuler().ImplicitSolution(from, ys, h, n))

	for _, v := range solutionMatrix {
		fmt.Println(len(v))
	}

	logger.Log(mtrx(solutionMatrix))
}

func mtrx(matrix [][]float64) [][]interface{} {
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
