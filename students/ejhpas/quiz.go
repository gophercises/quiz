package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	questions, answers := readCSV("problems.csv")                                // Records for questions and answers, [0] and [1]
	var response byte                                                            // Response to begin timer/quiz
	var userAnswer string                                                        // User's answer to question
	var correct int                                                              // Correct answers
	var timeLimit = flag.Int("Time Limit", 30, "Set a time limit for the quiz.") // Change the integer for a different quiz time limit
	fmt.Printf("Press enter to start the timer and begin the quiz: ")
	fmt.Scanf("%c", &response)
	timer := NewTimer(*timeLimit, func() {
		fmt.Println("\nTime's up! You got", correct, "answer(s) correct out of", len(answers), "questions.")
		os.Exit(0)
	})
	defer timer.Stop()
	for i := 0; i < len(questions); i++ {
		fmt.Printf("%s: ", questions[i])
		fmt.Scanf("%s ", &userAnswer)
		userAnswer = strings.Trim(userAnswer, " ")
		userAnswer = strings.ToLower(userAnswer) // Lowercase answers for comparison to avoid case sensitivity
		if userAnswer == answers[i] {
			correct++
		}
	}
	fmt.Println("You got", correct, "answer(s) correct out of", len(answers), "questions.")
}

// Read CSV and separate records questions and answers
func readCSV(myFile string) ([]string, []string) {
	var questions []string
	var answers []string
	myCSV, err := os.Open(myFile)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	read := csv.NewReader(myCSV)
	for {
		record, err := read.Read()
		if err == io.EOF {
			return questions, answers
		} else if err != nil {
			log.Fatal("Error: ", err)
		}
		questions = append(questions, record[0])
		answers = append(answers, strings.ToLower(record[1])) // Lowercase answers for comparison to avoid case sensitivity
	}
}

// NewTimer creates a new timer that will send the current time on its channel after duration
func NewTimer(seconds int, action func()) *time.Timer {
	timer := time.NewTimer(time.Second * time.Duration(seconds))
	go func() {
		<-timer.C
		action()
	}()
	return timer
}
