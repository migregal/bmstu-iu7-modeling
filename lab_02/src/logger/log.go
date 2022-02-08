package logger

import (
	"encoding/csv"
	"log"
	"os"
)

func LogToCsv(headers []interface{}, matrix [][]interface{}, filename string) {
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
	for _, row := range matrix {
		err := writer.Write(toString(row))
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
