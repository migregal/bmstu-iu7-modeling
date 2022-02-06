package solvers

import "math"

type Picard struct {
}

func NewPicard() Picard {
	return Picard{}
}

func (p Picard) Solution(x0, h float64, n int) [][]float64 {
	r := make([][]float64, 4)
	for i := 0; i < 4; i++ {
		r[i] = make([]float64, 0)
	}

	for i := 0; i <= n; i++ {
		r[0] = append(r[0], p.approx1(x0))
		r[1] = append(r[1], p.approx2(x0))
		r[2] = append(r[2], p.approx3(x0))
		r[3] = append(r[3], p.approx4(x0))

		x0 += h
	}

	return r
}

func (p Picard) approx1(x float64) float64 {
	return x * x * x / 3
}

func (p Picard) approx2(x float64) float64 {
	return p.approx1(x) + math.Pow(x, 7)/63
}

func (p Picard) approx3(x float64) float64 {
	return p.approx2(x) + 2*math.Pow(x, 11)/2079 + math.Pow(x, 15)/59535
}

func (p Picard) approx4(x float64) float64 {
	return p.approx2(x) + 2*math.Pow(x, 11)/2079 + 13*math.Pow(x, 15)/218295 +
		82*math.Pow(x, 19)/37328445 + 662*math.Pow(x, 23)/10438212015 +
		4*math.Pow(x, 27)/3341878155 + math.Pow(x, 31)/109876902975
}
