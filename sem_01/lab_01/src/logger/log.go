package logger

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var (
	headers = []interface{}{"X", "1", "2", "3", "4", "e1", "e2", "rk"}
)

func Log(matrix [][]interface{}, each int64) {
	headerFmt := color.New(color.FgGreen, color.Underline)
	columnFmt := color.New(color.FgYellow)

	tbl := table.New(headers...)
	tbl.WithHeaderFormatter(headerFmt.SprintfFunc())
	tbl.WithFirstColumnFormatter(columnFmt.SprintfFunc())

	for i, row := range matrix {
		if i%int(each) != 0 {
			continue
		}

		tbl.AddRow(row...)
	}

	tbl.Print()
}

func LogToCsv(matrix [][]interface{}, each int64, filename string) {
	toString := func(slice []interface{}) []string {
		strArray := make([]string, len(slice))
		for i, arg := range slice {
			strArray[i] = arg.(string)
		}
		return strArray
	}
	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	matrix = append([][]interface{}{headers}, matrix...)
	for i, row := range matrix {
		if i%int(each) != 0 {
			continue
		}

		err := writer.Write(toString(row))
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
