package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type question struct {
	problem string
	result  string
}

type quiz struct {
	questions []question
	score     int
}

func (q *quiz) ask() {
	q.score = 0
	for _, question := range q.questions {
		fmt.Println(question.problem)
		var answer string
		fmt.Scanln(&answer)
		if answer == question.result {
			q.score++
		}
	}
}

func readCSV(p string) []question {
	qfile, err := os.Open(p)
	if err != nil {
		log.Fatalf("could not open file: %v\n", err)
	}
	defer qfile.Close()
	csvrd := csv.NewReader(qfile)
	records, err := csvrd.ReadAll()
	if err != nil {
		log.Fatalf("could not parse .csv: %v\n", err)
	}
	var questions []question
	for _, record := range records {
		q := question{problem: record[0], result: record[1]}
		questions = append(questions, q)
	}
	return questions
}

var (
	qCSV    string
	timeout int
)

func init() {
	flag.StringVar(&qCSV, "quiz", "", "A .csv file with questions and answers.")
	flag.IntVar(&timeout, "timeout", 30, "The time limit for answering questions.")
	flag.Parse()
}

func main() {
	if qCSV == "" {
		log.Fatalln("a .csv file must be provided through the -quiz flag")
	}
	questions := readCSV(qCSV)
	qz := quiz{questions: questions, score: 0}
	timeoutCh := time.After(time.Duration(timeout) * time.Second)
	resultCh := make(chan quiz)
	go func() {
		qz.ask()
		resultCh <- qz
	}()
	select {
	case <-resultCh:
	case <-timeoutCh:
		fmt.Println("Timeout!")
	}
	fmt.Println("Score: ", qz.score)
}
