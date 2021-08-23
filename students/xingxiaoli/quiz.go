package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	filename := flag.String("file", "problems.csv", "Questions file")
	timelimit := flag.Int("time", 30, "Time limit of quiz in seconds")
	flag.Parse()

	renderTime := time.Second * time.Duration(*timelimit)
	timer := time.NewTimer(renderTime)

	fmt.Printf("Press 'Enter' to start the game")
	fmt.Scanln()

	quizs, err := readCSV(*filename)
	if err != nil {
		fmt.Println("read csv error: ", err)
		os.Exit(1)
	}

	answer := make(chan string)

	score := 0

floop:
	for _, qz := range quizs {
		go getAnswer(answer, qz)
		timer.Reset(renderTime)
	loop:
		for {
			select {
			case <-timer.C:
				fmt.Println("Timout")
				break floop
			case as := <-answer:
				if as == qz.answer {
					score += 1
				}
				break loop
			}
		}
	}

	fmt.Printf("Total Question: %d, Your Score: %d\n", len(quizs), score)
}

type Quiz struct {
	question, answer string
}

func getAnswer(ch chan string, qz *Quiz) {
	fmt.Printf("Question: %s, Answer?\n", qz.question)

	var input string
	fmt.Scanln(&input)
	ch <- input
}

func readCSV(filepath string) ([]*Quiz, error) {
	var quizs = make([]*Quiz, 0)

	fs, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	rd := csv.NewReader(fs)
	content, err := rd.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range content {
		qz := &Quiz{
			question: row[0],
			answer:   row[1],
		}
		quizs = append(quizs, qz)
	}
	return quizs, nil
}
