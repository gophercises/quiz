package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"errors"
)

var problemFileName = flag.String("csv", "./problems.csv", "a csv file in the format 'quastion,answer'")
var limit = flag.Int("limit", 30, "the time limit for the quiz in seconds")

type recordtype struct {
	question string
	answer   string
}

// ReadStringWithLimitTime - function read string from reader with time limit
func  ReadStringWithLimitTime(limit int, reader *bufio.Reader) (string, error) {
	timer := time.NewTimer(time.Duration(limit) * time.Second).C
	doneChan := make(chan bool)
	answer, err := "", error(nil)
	go func() {
		answer, err = reader.ReadString('\n')
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

func main() {

	flag.Parse()

	problemFile, err := os.Open(*problemFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer problemFile.Close() // close CSV file

	readerProblem := csv.NewReader(problemFile)

	var problems []recordtype

	for {
		record, err := readerProblem.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			println(err)
		}
		problems = append(problems, recordtype{record[0], strings.ToLower(strings.Trim(record[1], "\n "))})
	}

	reader := bufio.NewReader(os.Stdin)

	successAnswerCount := 0
	for i := range problems {
		print("Problem #", i+1, ": ", problems[i].question, " = ")

		answer, err := ReadStringWithLimitTime(*limit, reader)
		if err != nil {
			println("Time expire!")			
			break
		}
		if strings.ToLower(strings.Trim(answer, "\n ")) != problems[i].answer {
			break
		} else {
			successAnswerCount++
		}
	}
	println("You scored", successAnswerCount, "out of", len(problems))

}
