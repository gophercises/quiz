package main

import (
	"encoding/csv"
	"io"
	"log"
)

// Problem represents a single question answer pair
type Problem struct {
	question string
	answer   string
}

// ReadCSV parses the CSV file into a Problem struct
func ReadCSV(reader io.Reader) []Problem {
	csvReader := csv.NewReader(reader)

	problems := []Problem{}
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error reading CSV:", err)
		}

		problem := Problem{question: record[0], answer: record[1]}
		problems = append(problems, problem)
	}

	return problems
}
