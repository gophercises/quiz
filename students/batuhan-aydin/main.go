package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	csvUsage       = "a csv file in the format of 'question,answer' (default \"problems.csv\")"
	timeLimitUsage = "the time limit for the quiz in seconds (default 30)"
)

type record struct {
	q string // question
	r string // user's reply
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", csvUsage)
	timeLimit := flag.Int("time", 30, timeLimitUsage)
	flag.Parse()

	f := openFile(*csvFileName)

	problems := readFile(f)

	scanner := bufio.NewScanner(os.Stdin)
	rightAnswers := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, v := range problems {
		fmt.Printf("Problem #%d: %v = ", i+1, v.q)
		userReplyCh := make(chan string)
		go func() {
			scanner.Scan()
			userAnswer := strings.TrimSpace(scanner.Text())
			userReplyCh <- userAnswer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", rightAnswers, len(problems))
			return
		case answer := <-userReplyCh:
			if answer == v.r {
				rightAnswers++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", rightAnswers, len(problems))
}

func openFile(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalln("Could not open the file", err)
	}
	return f
}

func readFile(f *os.File) []record {
	reader := csv.NewReader(f)
	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error while reading file", err)
	}
	problems := make([]record, len(questions))
	for i, v := range questions {
		problems[i] = record{q: v[0], r: v[1]}
	}
	return problems

}
