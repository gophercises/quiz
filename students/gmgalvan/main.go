package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

type StatusQuiz struct {
	Completed bool
	Response  string
}

func ReadCSV(file *os.File, limit int) []Quiz {
	reader := csv.NewReader(bufio.NewReader(file))
	var quiz []Quiz
	lineCounter := 0
	for {
		if lineCounter == limit {
			break
		}
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal("error")
		}
		quiz = append(quiz, Quiz{Question: line[0], Answer: line[1]})
		lineCounter++
	}
	return quiz
}

func StartQuiz(quizContent []Quiz, quizResult chan StatusQuiz) {
	positivesResponses := 0
	reader := bufio.NewReader(os.Stdin)
	noQuestions := len(quizContent)
	pResult := ""
	aResult := ""
	result := ""
	var Status StatusQuiz
	Status.Completed = false
	for _, question := range quizContent {
		fmt.Print(question.Question, ": ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == question.Answer {
			positivesResponses = positivesResponses + 1
		}
		pResult = strconv.FormatInt(int64(positivesResponses), 10)
		aResult = strconv.FormatInt(int64(noQuestions), 10)
		result = "You scored " + pResult + " out of " + aResult
		Status.Response = result
		quizResult <- Status
	}
	Status.Completed = true
	quizResult <- Status

}

func main() {
	var csvFile string
	var limitAsk int
	var limitAskTime time.Duration

	flag.StringVar(&csvFile, "csv", "problems.csv", "a csv file in the format of \"question answer\" (deafult \"problems.csv\")")
	flag.IntVar(&limitAsk, "limit", 10, "limit for the quiz in questions (default 10)")
	flag.DurationVar(&limitAskTime, "time", 30000*time.Millisecond, "limit for the quiz in seconds (default 30000)")
	flag.Parse()

	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	quizContent := ReadCSV(file, limitAsk)
	timer := time.NewTimer(limitAskTime)
	quizResult := make(chan StatusQuiz)
	go StartQuiz(quizContent, quizResult)

	endQuiz := false
	r := "You scored 0 out of 0"
	for {
		select {
		case <-timer.C:
			endQuiz = true
		case b := <-quizResult:
			r = b.Response
			if b.Completed {
				endQuiz = true
			}
		}
		if endQuiz {
			break
		}
	}
	fmt.Println("\n")
	fmt.Println(r)
}
