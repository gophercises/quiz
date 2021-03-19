package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	question string
	answer string
}

type Scoreboard struct {
	correct int
	total int
}

func main() {
	filename := "problems.csv"
	records, err := readCSV(filename)

	if err != nil {
		log.Fatal(err)
	}

	questions := convertToQuestions(records)

	scoreboard := startQuiz(questions,30)
	reportResult(scoreboard)
}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("CSV file does not exist")
	}

	return csv.NewReader(file).ReadAll()
}

func convertToQuestions(records [][]string) []Question {
	questions := make([]Question, len(records))
	for i, question := range records {
		questions[i] = Question{question: question[0], answer: question[1]}
	}
	return questions
}

func sanitize(str string) string {
	t := strings.Trim(str, "\n")
	t = strings.Trim(t, " ")
	return t
}

func getInput(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input <- sanitize(result)
	}
}

func startQuiz(questions []Question, timeLimit int) Scoreboard {
	scoreboard := Scoreboard{total: len(questions)}
	done := make(chan bool)

	go askQuestions(questions, &scoreboard, done)

	select {
	case <- done:
		return scoreboard
	case <- time.After(time.Duration(timeLimit) * time.Second):
		return scoreboard
	}
}

func askQuestions(questions []Question, scoreboard *Scoreboard, done chan<- bool) {
	input := make(chan string)

	go getInput(input)

	for i, question := range questions {
		fmt.Printf("Problem %d: %s: ", i + 1, question.question)

		response := <- input

		if response == question.answer {
			scoreboard.correct += 1
		}
	}

	done <- true
}

func reportResult(scoreboard Scoreboard) {
	fmt.Printf("Congratulations, you answered %d out of %d right", scoreboard.correct, scoreboard.total)
}
