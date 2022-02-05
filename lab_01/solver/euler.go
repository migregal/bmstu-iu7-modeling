package solvers

import "math"

type Euler struct {
}

func NewEuler() Euler {
	return Euler{}
}

func (e Euler) Solution(x0, y0, h float64, n int) []float64 {
	r := make([]float64, 0)

	r = append(r, y0)

	for i := 0; i < n; i++ {
		y0 += h * equation(x0, y0)
		x0 += h

		r = append(r, y0)
	}

	return r
}

func (e Euler) ImplicitSolution(x0, y0, h float64, n int) []float64 {
	r := make([]float64, 0)

	r = append(r, y0)
	for i := 0; i < n; i++ {
		D := 1 - 4*h*(y0+h*(x0+h)*(x0+h))

		y0 = (1 - math.Sqrt(D)) / 2 / h
		x0 += h

		r = append(r, y0)
	}

	return r
}
