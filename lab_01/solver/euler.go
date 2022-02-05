package solvers

type Euler struct {
}

func NewEuler() Euler {
	return Euler{}
}

func (e Euler) Solution(x0, y0, h float64, n int) []float64 {
	r := make([]float64, 0)

	for i := 0; i <= n; i++ {
		r = append(r, y0)
		y0 += h * equation(x0, y0)
		x0 += h
	}

	return r
}
