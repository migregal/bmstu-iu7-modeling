package main

import (
	"flag"
	logger "lab_01/logger"
	solvers "lab_01/solver"
	"math"
)

type Flags struct {
	from *float64
	to   *float64
	h    *float64
	y0   *float64
	each *int64

	genOutput      *bool
	outputFilename *string
}

func inputArgs() Flags {
	flags := Flags{
		from: flag.Float64("from", 0, "Start of serach interval"),
		to:   flag.Float64("to", 2, "End of search interval"),
		h:    flag.Float64("h", 1e-4, "Step for search interval"),
		y0:   flag.Float64("y0", 0, "Base value for algorythms"),
		each: flag.Int64("each", 1, "Step for result to display"),

		genOutput:      flag.Bool("csv", false, "Output the result of the csv file"),
		outputFilename: flag.String("output", "result.csv", "The output filename"),
	}

	flag.Parse()

	return flags
}

func process(f Flags) [][]float64 {
	n := int(math.Ceil(math.Abs(*f.to-*f.from) / *f.h))

	result := make([][]float64, 0)
	result = append(result, solvers.InitX(*f.from, *f.h, n))

	result = append(result,
		solvers.NewPicard().Solution(*f.from, *f.h, n)...)

	if *f.from >= 0 {
		result = append(result,
			solvers.NewEuler().Solution(*f.from, *f.y0, *f.h, n))
		result = append(result,
			solvers.NewEuler().ImplSolution(*f.from, *f.y0, *f.h, n))
		result = append(result,
			solvers.NewRK().Solution(*f.from, *f.y0, 0.5, *f.h, n))
		return result
	}

	r := reverseSlice(solvers.NewEuler().Solution(0, *f.y0, -*f.h, n/2))
	r = append(r, solvers.NewEuler().Solution(0+*f.h, *f.y0, *f.h, n/2-1)...)
	result = append(result, r)

	r = reverseSlice(solvers.NewEuler().ImplSolution(0, *f.y0, -*f.h, n/2))
	r = append(r, solvers.NewEuler().ImplSolution(0+*f.h, *f.y0, *f.h, n/2-1)...)
	result = append(result, r)

	r = reverseSlice(solvers.NewRK().Solution(0, *f.y0, 0.5, -*f.h, n/2))
	r = append(r, solvers.NewRK().Solution(0, *f.y0, 0.5, *f.h, n/2-1)...)
	result = append(result, r)

	return result
}

func main() {
	f := inputArgs()

	result := process(f)

	logger.Log(solvers.RotateMtrx(result), *f.each)
	if *f.genOutput {
		logger.LogToCsv(solvers.RotateMtrx(result),  *f.each, *f.outputFilename)
	}
}
