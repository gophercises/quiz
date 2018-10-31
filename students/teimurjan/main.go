package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Problem object
type Problem struct {
	Question string
	Answer   string
}

// ValidateAnswer validates answer for the problem
func (p *Problem) ValidateAnswer(answer string) bool {
	return p.Answer == answer
}

// Quiz object
type Quiz struct {
	Problems []Problem
	Score    int
}

func main() {
	var file string
	flag.StringVar(&file, "file", "problems.csv", "--file=path/to/problems/file")
	var quizTime int
	flag.IntVar(&quizTime, "time", 30, "--time=15")
	flag.Parse()

	WaitForStart()
	quiz := Quiz{
		Problems: ParseProblemsFrom(file),
		Score:    0,
	}
	go func() {
		<-time.After(time.Duration(quizTime) * time.Second)
		ShowTimeIsUpMessage()
		ShowFinalMessage(quiz.Score, len(quiz.Problems))
		os.Exit(0)
	}()
	RunQuiz(&quiz)
	ShowFinalMessage(quiz.Score, len(quiz.Problems))
}

// WaitForStart makes the program wait for a user to press Enter button
func WaitForStart() {
	fmt.Print("Press 'Enter' to start the quiz.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// ParseProblemsFrom parses problems from file by the provided path
func ParseProblemsFrom(pathToFile string) []Problem {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal("File does not exists")
	}
	reader := csv.NewReader(bufio.NewReader(file))
	var problems []Problem
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}
	return problems
}

// RunQuiz starts the quiz
func RunQuiz(q *Quiz) {
	reader := bufio.NewReader(os.Stdin)
	for _, problem := range q.Problems {
		AskQuestion(&problem)
		answer := ReadLine(reader)
		if problem.ValidateAnswer(answer) {
			q.Score++
		}
	}
}

// AskQuestion asks a problem's question
func AskQuestion(p *Problem) {
	fmt.Print(p.Question + " ")
}

// ReadLine read a line using buifo.Reader
func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

// ShowTimeIsUpMessage shows time is up message
func ShowTimeIsUpMessage() {
	fmt.Println("\rTime is up!")
}

// ShowFinalMessage shows final message with the correctness statistics
func ShowFinalMessage(correctAnswersCount int, problemsCount int) {
	fmt.Printf("\rThere is/are %d correct answers given for %d problems.\n", correctAnswersCount, problemsCount)
}
