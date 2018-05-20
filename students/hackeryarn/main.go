package main

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/hackeryarn/quiz/students/hackeryarn/problem"
)

// ReadCSV parses the CSV file into a Problem struct
func ReadCSV(reader io.Reader) []problem.Problem {
	csvReader := csv.NewReader(reader)

	problems := []problem.Problem{}
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error reading CSV:", err)
		}

		problems = append(problems, problem.New(record))
	}

	return problems
}
