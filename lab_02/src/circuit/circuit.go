package circuit

import (
	"math"

	"gonum.org/v1/gonum/integrate/quad"
)

type Circuit struct {
	R   float64
	Le  float64
	Lk  float64
	Ck  float64
	Rk  float64
	Uc0 float64
	I0  float64
	Tw  float64
}

type ModulationRes struct {
	IRes   []float64
	RpRes  []float64
	UcRes  []float64
	T0Res  []float64
	IRpRes []float64
	TRes   []float64
}

func Default() Circuit {
	return Circuit{0.35, 12, 187e-6, 268e-6, 0.25, 1400, 1.5, 2000}
}

func (c Circuit) Modulate(from, to, h float64) ModulationRes {
	var (
		I   = c.I0
		Uc  = c.Uc0
		res ModulationRes
	)

	for _, t := range arange2(from, to, h) {
		T0 := c.getT0(I)
		Rp := c.getRp(I, T0, c.getM(I))
		I, Uc = c.getRungeKutta(t, I, Uc, h, Rp)

		if t <= h {
			continue
		}

		res.TRes = append(res.TRes, t)
		res.IRes = append(res.IRes, I)
		res.RpRes = append(res.RpRes, Rp)
		res.UcRes = append(res.UcRes, Uc)
		res.T0Res = append(res.T0Res, T0)
		res.IRpRes = append(res.IRpRes, I*Rp)
	}

	return res
}

func (c Circuit) getT0(I float64) float64 {
	return interpolate(I, curTbl[0], curTbl[1])
}

func (c Circuit) getM(I float64) float64 {
	return interpolate(I, curTbl[0], curTbl[2])
}

func (c Circuit) getRp(I, T0, m float64) float64 {
	f := func(x float64) float64 {
		return c.getSigma(c.getT(x, T0, m)) * x
	}
	val := quad.Fixed(f, 0, 1, 30, nil, 0)

	return c.Le / (2 * math.Pi * c.R * c.R * val)
}

func (c Circuit) getT(z, T0, m float64) float64 {
	return T0 + (c.Tw-T0)*math.Pow(z, m)
}

func (c Circuit) getF(y, z, Rp float64) float64 {
	return (z - (c.Rk+Rp)*y) / c.Lk
}

func (c Circuit) getPhi(y float64) float64 {
	return -y / c.Ck
}

func (c Circuit) getSigma(T float64) float64 {
	return interpolate(T, tmpTbl[0], tmpTbl[1])
}
