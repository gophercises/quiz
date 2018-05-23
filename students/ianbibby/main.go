package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// qa value holds a question/answer pair
type qa struct {
	question string
	answer   string
}

// Quiz stores the state of the quiz
type Quiz struct {
	qas          []*qa
	timeout      int
	done         chan bool
	question     chan *qa
	ask          chan *qa
	answered     chan bool
	correctCount int
	askedCount   int
}

// NewQuiz instatiates a new Quiz struct
func NewQuiz(timeout int) *Quiz {
	return &Quiz{
		timeout:  timeout,
		done:     make(chan bool),
		question: make(chan *qa),
		ask:      make(chan *qa),
		answered: make(chan bool),
	}
}

// timer starts a timer loop, for timeout control
func (q *Quiz) timer() {
	defer close(q.done)

	for {
		select {
		case <-time.After(time.Duration(q.timeout) * time.Second):
			fmt.Print("\tTime's up!\n")
			q.done <- true
			return
		case <-q.done:
			fmt.Println("Done")
			return
		}
	}
}

// readFile reads question data into an array
func (q *Quiz) readFile(filename string) {
	csvData, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer csvData.Close()

	reader := csv.NewReader(csvData)
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read from file: %v", err)
		}

		q.qas = append(q.qas, &qa{rec[0], rec[1]})
	}
}

// processQuestions loops through the questions array,
// sending each question to the ask channel, then
// waiting to be notified of an answer
func (q *Quiz) processQuestions() {
	defer close(q.ask)

	for _, question := range q.qas {
		q.ask <- question
		select {
		case <-q.answered:
			// Wait until the question has been answered before moving on
		case <-q.done:
			// Canceled, possibly due to a timeout
			return
		}
	}

	q.done <- true
}

// askQuestion ranges over the ask channel for questions to be asked,
// then waits for an answer over stdin, before sending an answered
// notification
func (q *Quiz) askQuestion() {
	defer close(q.answered)

	for qa := range q.ask {
		fmt.Printf("Question #%-3d %s= ", q.askedCount+1, qa.question)
		q.askedCount++

		var answer string
		fmt.Fscanln(os.Stdin, &answer)
		if strings.TrimSpace(answer) == qa.answer {
			q.correctCount++
		}
		q.answered <- true
	}
}

// printSummary prints a summary of the quiz play
func (q *Quiz) printSummary() {
	fmt.Printf("\nYou answered %d/%d questions correctly!\n", q.correctCount, q.askedCount)
}

// shuffle shuffles the question order
func (q *Quiz) shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(q.qas), func(i, j int) {
		q.qas[i], q.qas[j] = q.qas[j], q.qas[i]
	})
}

func main() {
	/*
		  Program description:
		  	The quiz accepts a CSV file and allows for a custom timeout and
			randomized order of asking.

			The implementation is channel-based:

			[question array] -> questions -> ask -> answered -> done

			A done channel allows graceful exit for when all questions have been
			asked, or a timeout has occured.
	*/
	const (
		defaultCSV     = "problems.csv"
		defaultTimeout = 30
	)
	var (
		csvArg     string
		timeoutArg int
		shuffle    bool
	)

	flag.StringVar(&csvArg, "csv", defaultCSV, "Path to a CSV problem file")
	flag.IntVar(&timeoutArg, "timeout", defaultTimeout, "Timeout value in seconds.  0 to disable")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the questions")
	flag.Parse()

	quiz := NewQuiz(timeoutArg)
	quiz.readFile(csvArg)

	if shuffle {
		quiz.shuffle()
	}

	if timeoutArg > 0 {
		fmt.Println("Press any key to start...")
		fmt.Fscanf(os.Stdin, "%s")

		go quiz.timer()
	}
	go quiz.processQuestions()
	go quiz.askQuestion()

	<-quiz.done
	quiz.printSummary()
}
