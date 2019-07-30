package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	rightAnswers int
	wrongAnswers int
	unanswered   int
	total        int
)

func getQuestions(file string) [][]string {
	// Open the csv file
	inputBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}

	// Convert the problems file from bytes to a string
	data := string(inputBytes)

	// First, we create a Reader type out of the data (string) and then we pass it to the csv.NewReader()
	// so that we can read the data from the csv file.
	r := csv.NewReader(strings.NewReader(data))
	r.Comma = ','

	// Read all of the records into memory
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}

func finalScore(rightAnswers int, wrongAnswers int, unanswered int, total int) {
	if unanswered != 0 {
		wrongAnswers = wrongAnswers + unanswered
	}
	fmt.Printf("\n --------- \n\nRight: %d\nWrong: %d\nTotal: %d\n", rightAnswers, wrongAnswers, total)
}

func askQuestions(records [][]string) {
	// This reads in inputs from the STDIN device to the NewReader memory
	reader := bufio.NewReader(os.Stdin)
	total = len(records)
	unanswered = total

	for _, record := range records {
		fmt.Printf("What is %s?\n", record[0])
		text, _ := reader.ReadString('\n')
		if strings.TrimRight(text, "\n") == record[1] {
			rightAnswers++
			fmt.Printf("Input: %sAnswer: %s\n", text, record[1])
		} else {
			wrongAnswers++
			fmt.Printf("Input: %sAnswer: %s\n", text, record[1])
		}
		unanswered--
	}

	finalScore(rightAnswers, wrongAnswers, unanswered, total)
}

func newTimer(seconds int) {
	limit := time.Second * time.Duration(seconds)
	timer := time.NewTimer(limit)
	go func() {
		<-timer.C
		finalScore(rightAnswers, wrongAnswers, unanswered, total)
		os.Exit(0)
	}()
}

func main() {
	file := flag.String("filename", "problems.csv", "Pass in the filename of the csv file.")
	seconds := flag.Int("time", 30, "This represents how much time a user will have to complete the quiz.")
	flag.Parse()

	questions := getQuestions(*file)
	newTimer(*seconds)
	askQuestions(questions)
}
