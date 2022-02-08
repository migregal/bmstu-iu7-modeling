package main

import (
	"flag"
	"lab_02/circuit"
	"lab_02/logger"
)

type Flags struct {
	from *float64
	to   *float64
	h    *float64

	outputFilename *string
}

func inputArgs() Flags {
	flags := Flags{
		from: flag.Float64("from", 0, "Start of serach interval"),
		to:   flag.Float64("to", 8e-4, "End of search interval"),
		h:    flag.Float64("h", 1e-6, "Step for search interval"),

		outputFilename: flag.String("output", "result.csv", "The output filename"),
	}

	flag.Parse()

	return flags
}

var (
	headers = []interface{}{"t", "I", "Rp", "Uc", "T0", "IRp"}
)

func main() {
	f := inputArgs()

	modulation := circuit.Default().Modulate(
		*f.from, *f.to, *f.h)

	logger.LogToCsv(
		headers,
		circuit.RotateMtrx(
			[][]float64{
				modulation.TRes,
				modulation.IRes,
				modulation.RpRes,
				modulation.UcRes,
				modulation.T0Res,
				modulation.IRpRes,
			},
		),
		*f.outputFilename,
	)
}
