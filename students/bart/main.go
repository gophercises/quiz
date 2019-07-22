package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func fatalError(message string, err error) {
	if err != nil {
		log.Fatalln(message, ":", err)
	}
}

type question struct {
	question string
	answer   string
}

type quiz struct {
	answered          int
	answeredCorrectly int
	questions         []question
}

// loadQuiz loads all questions into memory, assumed to be safe as
// the instructions state that the quiz will be < 100 questions
func loadQuiz(filePath string) *quiz {
	csvFile, err := os.Open(filePath)
	fatalError("Error opening quiz CSV file", err)
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var quiz quiz
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		fatalError("Error parsing CSV", err)
		question := question{line[0], line[1]}
		quiz.questions = append(quiz.questions, question)
	}
	return &quiz
}

// run prints questions, check answers, and records total
// answered and total correct
func (quiz *quiz) run() {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
quizLoop:
	for _, question := range quiz.questions {
		fmt.Println(question.question)
		answerCh := make(chan string)
		go func() {
			scanner.Scan()
			answer := scanner.Text()
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			break quizLoop
		case answer := <-answerCh:
			if answer == question.answer {
				quiz.answeredCorrectly++
			}
			quiz.answered++
		}
	}
	return
}

// report prints summary of quiz performance
func (quiz *quiz) report() {
	fmt.Printf(
		"You answered %v questions out of a total of %v and got %v correct",
		quiz.answered,
		len(quiz.questions),
		quiz.answeredCorrectly,
	)
}

var (
	scanner     = bufio.NewScanner(os.Stdin)
	filePathPtr = flag.String("file", "./problems.csv", "Path to csv file containing quiz.")
	timeLimit   = flag.Int64("time-limit", 30, "Set the total time in seconds allowed for the quiz.")
)

func main() {
	quiz := loadQuiz(*filePathPtr)
	quiz.run()
	quiz.report()
}
