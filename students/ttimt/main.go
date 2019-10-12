package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const defaultFileName = "problems.csv"
const defaultMaxTimeInSeconds = 30

// Store flags from CLI
var fileName string
var maxTimeInSeconds int

func init() {
	flag.StringVar(&fileName, "csv", defaultFileName, "This is the filename of the quiz in csv format")
	flag.IntVar(&maxTimeInSeconds, "limit", defaultMaxTimeInSeconds, "This is the max time in seconds to answer all the questions")
}

func main() {
	// Parse CLI flags
	flag.Parse()

	var (
		nrOfCorrectAnswer int
		userAnswer        string

		// channel
		hasEnded = make(chan bool)
	)

	// Open csv file of name "fileName"
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(errors.New("error: file not found: " + fileName))
		return
	}
	defer csvFile.Close()

	// Read the csv file
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(errors.New("error: file not read: " + err.Error()))
		return
	}

	// Start the quiz and timer on detection of enter keystroke
	fmt.Println("Press enter to start the quiz!")
	_, _ = fmt.Scanln(&userAnswer)

	go func() {
		// For each question
		for _, record := range records {
			fmt.Print(record[0] + ": ")
			_, _ = fmt.Scanln(&userAnswer)

			if strings.TrimSpace(userAnswer) == record[1] {
				nrOfCorrectAnswer++
			}
		}

		close(hasEnded)
	}()

	// Time out or quiz end first
	select {
	case <-hasEnded:
		break
	case <-time.After(time.Duration(maxTimeInSeconds) * time.Second):
		fmt.Println("\nTime out!")
	}

	// Print out the result
	fmt.Println("Number of correct answer: ", nrOfCorrectAnswer)
}
