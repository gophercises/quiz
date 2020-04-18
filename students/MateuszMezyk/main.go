package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type quizResult struct {
	correctAnswers      int
	totalQuestionsAsked int
	totalQuestions      int
}

func main() {

	filePath := flag.String("file", "./../../problems.csv", "You can pass different file containing quiz questions as a relative path")
	timeLimit := flag.Int("time-limit", 30, "Specify time limit for quiz in seconds. By default 30 seconds.")
	randomOrder := flag.Bool("random", false, "Set to'true' if you want questions to be presented in random order")
	flag.Parse()

	records := loadQuestionsAndAnswers(*filePath)
	quizResult := quizResult{totalQuestions: len(records)}

	var quizChan = make(chan bool)
	fmt.Println("Click enter to start.")
	var enterHit string
	fmt.Scanln(&enterHit)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	if *randomOrder {
		go askQuestionsInRandomOrder(records, &quizResult, &quizChan)
	} else {
		go askQuestions(records, &quizResult, &quizChan)
	}

	select {
	case <-quizChan:
		fmt.Println("You have responded to all of questions withing time limit. Congratulation!.")
		break
	case <-timer.C:
		fmt.Println("Time is over.")
		break
	}

	fmt.Printf("You have correctly answered to %d out of %d questions asked. Total number of questions in quiz is %d.", quizResult.correctAnswers, quizResult.totalQuestionsAsked, quizResult.totalQuestions)

}

func loadQuestionsAndAnswers(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Cannot open file", filePath, err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read file", filePath, err)
	}
	return records
}

func askQuestions(records [][]string, quizResult *quizResult, c *chan bool) {
	for _, row := range records {
		askSingleQuestion(row, quizResult, c)
	}
	*c <- true
}

func askQuestionsInRandomOrder(records [][]string, quizResult *quizResult, c *chan bool) {
	alreadyUsedNumbers := make(map[int]bool)
	var index int
	for i := 0; i < len(records); i++ {
		for {
			index = rand.Intn(len(records))
			if alreadyUsedNumbers[index] == false {
				alreadyUsedNumbers[index] = true
				break
			}
		}
		row := records[index]
		askSingleQuestion(row, quizResult, c)
	}
	*c <- true
}

func askSingleQuestion(row []string, quizResult *quizResult, c *chan bool) {
	fmt.Println("What is the result of following equation?", row[0])
	quizResult.totalQuestionsAsked++
	correctAnswer := row[1]
	var answer string
	fmt.Scanln(&answer)
	normalizeString(&correctAnswer)
	normalizeString(&answer)
	if answer == correctAnswer {
		quizResult.correctAnswers++
	}
}

func normalizeString(input *string) {
	*input = strings.TrimSpace(*input)
	*input = strings.ToLower(*input)
}
