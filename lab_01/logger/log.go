package logger

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func Log(matrix [][]interface{}) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New(
		"X",
		"Picard, 1st approx",
		"Picard, 2nd approx",
		"Picard, 3rd approx",
		"Picard, 4th approx",
		"Euler",
		"Implicit Euler")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, row := range matrix {
		tbl.AddRow(row...)
	}

	tbl.Print()
}
