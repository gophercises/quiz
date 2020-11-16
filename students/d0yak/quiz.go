package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// TODO: define timelimit flag
const timeInSeconds = 2

// A quiz game which asks the user basic math questions and print your score.
func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question','answer'")
	flag.Parse()

	csvFile, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprint("An error occured", err))
	}

	readFromCsv(csvFile)
	defer csvFile.Close()
}

// A global exit function errors out and takes an argument fmt.Sprint output
func exit(msg string) {
	fmt.Printf("\n") // better output
	fmt.Println(msg)
	os.Exit(1)
}

// Read from the CSV file and start the game
// TODO: this func should return a *Reader and pass that value to evaluateAnswers func
func readFromCsv(csvFile *os.File) {
	r := csv.NewReader(csvFile)
	i := 0 // problem number
	var input string
	var answers []bool

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Problem #%d: %s = ", i, record[0])
			timer := time.AfterFunc(time.Second*timeInSeconds, func() {
				exit(fmt.Sprintf("Time's Up! You've failed to answer in %d seconds. Try again!", timeInSeconds))
			})
			fmt.Scanln(&input)
			timer.Stop()
			// TODO: currently using a slice to append answers and compare the csv output 'answers' section in string type
			// it works for now but feels like it needs a professional touch, better approach.
			if input == record[1] {
				answers = append(answers, true)
			} else {
				answers = append(answers, false)
			}
			i += 1
		}
	}
	evaluateAnswers(answers)
}

// Calculate how many scores
func evaluateAnswers(answers []bool) {
	result := 0
	for _, value := range answers {
		if value == true {
			result += 1
		}
	}
	fmt.Printf("You scored %d out of %d\n", result, len(answers))
}
