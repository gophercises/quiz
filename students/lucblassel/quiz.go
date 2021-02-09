package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("filename", "problems.csv", "file path to the questions file")
	timeAmount := flag.Int("timeAmount", 30, "number of seconds to put on the time (optional)")
	flag.Parse()

	problems, err := readRecords(*filename)
	if err != nil {
		log.Fatal("error reading csv: ", err)
	}

	var ready string

	fmt.Printf("You will have %d seconds to complete the quizz.\nPress [ENTER] when you are ready to start.\n\n", *timeAmount)
	fmt.Scanf("%s", &ready)

	correct := timeQuestions(problems, *timeAmount)

	fmt.Printf("\nYou got %d / %d problems correct.", correct, len(problems))
}

type Problem struct {
	question string
	answer   string
}

func readRecords(path string) ([]Problem, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(csvFile)
	records := make([]Problem, 0)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(
			records,
			Problem{
				question: record[0],
				answer:   strings.TrimSpace(record[1]),
			})
	}
}

func askQuestions(problems []Problem, output chan int) {
	var userAnswer string
	for i, problem := range problems {

		fmt.Printf("#%d: %s  ", i+1, problem.question)
		fmt.Scanf("%s", &userAnswer)
		if userAnswer == problem.answer {
			output <- 1
		}
	}
	close(output)
}

func timeQuestions(problems []Problem, timeAmount int) int {
	var correct int
	timer := time.NewTimer(time.Duration(timeAmount) * time.Second)
	answers := make(chan int)

	go askQuestions(problems, answers)

	for {
		select {
		case <-timer.C:
			fmt.Println("\n\nTime's up.")
			return correct
		case _, ok := <-answers:
			if !ok {
				return correct
			}
			correct++
		}
	}
}
