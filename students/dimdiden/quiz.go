package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// DEFAULTFILE is the quiz file expected to be load by default
const DEFAULTFILE = "problems.csv"

func main() {
	// Flag block
	file := flag.String("f", DEFAULTFILE, "specify the path to file")
	timeout := flag.Int("t", 0, "specify the number of seconds for timeout")
	flag.Parse()
	// Failed if the number of seconds is negative
	if *timeout < 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	// Open file
	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Run the main logic
	total, correct, err := run(f, *timeout)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Number of questions: %v\nNumber of correct answers: %v\n", total, correct)
}

// run is the main function to execute quiz app
func run(qinput io.Reader, timeout int) (total, correct int, err error) {
	// Reading csv file and parse it to the records var
	r := csv.NewReader(qinput)
	records, err := r.ReadAll()
	if err != nil {
		return total, correct, err
	}
	// Two channels. One for answers, another for timeout
	answerChan := make(chan string)
	timerChan := make(chan time.Time, 1)
	// Iterate over the records
	for _, record := range records {
		question, expected := record[0], record[1]
		fmt.Printf("Question: %v. Answer: ", question)
		total++
		// Listening for input in the separate goroutine
		go getAnswer(answerChan)
		// Run the timer in separate goroutine if the timeout is specified
		if timeout > 0 {
			go func() { timerChan <- <-time.After(time.Duration(timeout) * time.Second) }()
		}
		// Main select block
		select {
		case answer := <-answerChan:
			if answer == expected {
				correct++
			}
		case <-timerChan:
			fmt.Println()
			return total, correct, errors.New("Timeout reached!")
		}
	}
	return total, correct, nil
}

// getAnswer func is for listening input from user
func getAnswer(answerChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	answerChan <- strings.Replace(answer, "\n", "", -1)
}
