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

type recordtype struct {
	question string
	answer   string
}

// ReadStringWithLimitTime - function read string from reader with time limit
func ReadStringWithLimitTime(limit int) (string, error) {
	timer := time.NewTimer(time.Duration(limit) * time.Second).C
	doneChan := make(chan bool)
	answer, err := "", error(nil)
	go func() {
		fmt.Scanf("%s\n", &answer)
		doneChan <- true
	}()
	for {
		select {
		case <-timer:
			return "", errors.New("Timer expired")
		case <-doneChan:
			return answer, err
		}
	}
}

// ParseLines  - parse lines from array of array of string to array of recordtype
func ParseLines(lines [][]string) []recordtype {
	ret := make([]recordtype, len(lines))
	for i, line := range lines {
		ret[i] = recordtype{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {

	problemFileName := flag.String("csv", "./problems.csv", "a csv file in the format 'quastion,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	problemFile, err := os.Open(*problemFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *problemFileName))
	}

	defer problemFile.Close() // close CSV file

	readerProblem := csv.NewReader(problemFile)
	lines, err := readerProblem.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := ParseLines(lines)

	successAnswerCount := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s=", i+1, p.question)

		answer, err := ReadStringWithLimitTime(*limit)
		if err != nil {
			println("Time expire!")
			break
		}
		if strings.ToLower(strings.Trim(answer, "\n ")) == p.answer {
			successAnswerCount++
		}
	}
	println("You scored", successAnswerCount, "out of", len(problems))

}
