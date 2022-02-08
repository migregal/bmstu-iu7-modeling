package circuit

type RCoeffs64 struct {
	Kn float64
	Pn float64
}

const order int = 4

func (c Circuit) getRungeKutta(x, y, z, h, Rp float64) (float64, float64) {
	cfsArr := make([]RCoeffs64, order)

	for i := 0; i < order; i++ {
		v := i
		if i == 0 {
			v = order - 1
		}
		_, yAdd, zAdd := getCurAdd(h, cfsArr[v], i, order)
		cfsArr[i] = RCoeffs64{
			h * c.getF(y+yAdd, z+zAdd, Rp),
			h * c.getPhi(y+yAdd),
		}
	}

	return getNextMembs(y, z, cfsArr)
}

func getNextMembs(y, z float64, cfsArr []RCoeffs64) (float64, float64) {
	var (
		kSum float64 = 0
		pSum float64 = 0
		div  float64 = float64(2*(len(cfsArr)-2) + 2)
	)

	for i := 0; i < len(cfsArr); i++ {
		if i > 0 && i < len(cfsArr)-1 {
			kSum += 2 * cfsArr[i].Kn
			pSum += 2 * cfsArr[i].Pn
		} else {
			kSum += cfsArr[i].Kn
			pSum += cfsArr[i].Pn
		}
	}

	return y + kSum/div, z + pSum/div
}

func getCurAdd(h float64, cfs RCoeffs64, i, ord int) (float64, float64, float64) {
	if i == 0 {
		return 0, 0, 0
	}
	if i == ord-1 {
		return h, cfs.Kn, cfs.Pn
	}
	return h / 2, cfs.Kn / 2, cfs.Pn / 2
}
