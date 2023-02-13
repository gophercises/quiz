package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var fileFlag = flag.String("f", "problems.csv", "The file to read problems from")
	var timeFlag = flag.Int("t", 30, "Time in seconds until the quiz times out")

	flag.Parse()

	problemsFile, err := os.Open(*fileFlag)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(problemsFile)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Press enter to start the timed quiz:")

	fmt.Scanln()

	timeLimit := time.Duration(*timeFlag) * time.Second

	done := make(chan struct{})

	var correctAnswers int

	go func() {
		for _, record := range records {
			fmt.Println(record[0])

			var answer string

			fmt.Scanln(&answer)

			if answer == strings.TrimSpace(record[1]) {
				correctAnswers++
			}
		}

		close(done)
	}()

	select {
	case <-time.After(timeLimit):
	case <-done:
	}

	fmt.Println(correctAnswers, "out of", len(records), "problems correctly answered")
}
