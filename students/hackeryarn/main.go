package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/gophercises/quiz/students/hackeryarn/problem"
	"github.com/gophercises/quiz/students/hackeryarn/quiz"
)

const (
	// FileFlag is the flat to be used to set a file
	FileFlag = "file"
	// FileFlagValue is the value that is used when no FileFlag is provided
	FileFlagValue = "problems.csv"
	// FileFlagUsage is the help string for the FileFlag
	FileFlagUsage = "Questions file"
)

// Flagger configures the flags used
type Flagger interface {
	StringVar(p *string, name, value, usage string)
}

type quizFlagger struct{}

func (q *quizFlagger) StringVar(p *string, name, value, usage string) {
	flag.StringVar(p, name, value, usage)
}

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

var file string

// ConfigFlags sets all the flags used by the application
func ConfigFlags(f Flagger) {
	f.StringVar(&file, FileFlag, FileFlagValue, FileFlagUsage)
}

func init() {
	flagger := &quizFlagger{}
	ConfigFlags(flagger)

	flag.Parse()
}

func main() {
	file, err := os.Open(file)
	if err != nil {
		log.Fatalln("Could not open file", err)
	}

	quiz := ReadCSV(file)
	quiz.Run(os.Stdout, os.Stdin)
}
