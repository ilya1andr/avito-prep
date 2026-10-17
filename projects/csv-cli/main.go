package main

import (
	"encoding/csv"
	"fmt"
	"io" // <-- добавляем импорт io
	"os"
)

func parseCSV(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)
	return reader.ReadAll()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: csv-cli <file.csv>")
		os.Exit(1)
	}

	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	rows, err := parseCSV(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse csv: %v\n", err)
		os.Exit(1)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}