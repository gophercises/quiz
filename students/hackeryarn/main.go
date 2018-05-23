package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/gophercises/quiz/students/hackeryarn/problem"
	"github.com/gophercises/quiz/students/hackeryarn/quiz"
)

// ReadCSV parses the CSV file into a Problem struct
func ReadCSV(reader io.Reader) quiz.Quiz {
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

	return quiz.New(problems)
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Could not open file", err)
	}

	quiz := ReadCSV(file)
	quiz.Run(os.Stdout, os.Stdin)
}
