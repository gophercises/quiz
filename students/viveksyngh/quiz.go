package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"strings"
	"log"
	"flag"
	"time"
	"math/rand"
) 

type Question struct {
	question string
	answer string
}

func getQuestions(filePath string) ([]Question) {
	file, err := os.Open(filePath)
	if(err != nil){
		log.Fatal("Failed to open file.")
	}
	reader := csv.NewReader(file)
	questionList, err := reader.ReadAll()
	if(err != nil) {
		log.Fatal("Failed to parse CSV file.")
	}
	questions := make([]Question, 0)
	
	for _, question := range questionList {
			questions = append(questions, Question{strings.TrimSpace(question[0]), 
													strings.TrimSpace(question[1])})
	}
	return questions
}

func Quiz(questions []Question, timer *time.Timer) (score int){

	for i, question := range questions {
		fmt.Printf("Problem #%d %s : ", i + 1, question.question)
		answerChannel := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			answerChannel <- userAnswer
		}()

		select {
			case <-timer.C:
				fmt.Println("\nTimeout")
				return score
			case userAnswer := <- answerChannel:
				if(strings.TrimSpace(userAnswer) == question.answer) {
					score += 1
				}
		}
	}
	return score
}

func randomize(questions []Question) []Question{
	n := len(questions)
	for i := n-1; i>0; i-- {
		j := rand.Intn(i)
		temp := questions[i]
		questions[i] = questions[j]
		questions[j] = temp
	}
	return questions
}

var csvPath string
var timeout int
var shuffle bool

func init() {
	flag.StringVar(&csvPath, "csv", "problems.csv", "a CSV file in format of 'question,answer'")
	flag.IntVar(&timeout, "limit", 30, "The time limit of the quiz in seconds")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the questions (default 'false')")
}

func main() {
	flag.Parse()
	fmt.Print("Hit Enter to start the timer:")
	questions := getQuestions(csvPath)
	if(shuffle) {
		questions = randomize(questions)
	}
	fmt.Scanln()
	timer := time.NewTimer(time.Second * time.Duration(timeout))
	score := Quiz(questions, timer)
	fmt.Printf("Your scored %d out of %d\n", score, len(questions))
}
