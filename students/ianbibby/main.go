/*
	The quiz accepts a CSV file and allows for a custom timeout and
	randomized shuffling of questions.

	The implementation is based on an actor model where the defined
	actions are:

		- AskQuestion
			- Prints the question.
			- Sends notification to listen for user input.
		- QuestionAnswered
			- Store the given answer.
			- Sends notification that an answer has been handled.
		- PrintSummary
			- Compares answers and prints the results.

	The timeout is achieved using a context, where a deadline is set
	whenever the timeout is > 0.

	The actor loop is terminated when either the last question has
	been answered, or the timeout has occured.
*/
package main

import (
	"context"
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
	ctx      context.Context
	ops      chan func() bool
	qas      []*qa
	answered map[*qa]string
}

// NewQuiz instatiates a new Quiz struct
func NewQuiz(ctx context.Context) *Quiz {
	return &Quiz{
		ctx:      ctx,
		ops:      make(chan func() bool),
		answered: map[*qa]string{},
	}
}

// LoadFromFile reads question data into an array
func (q *Quiz) LoadFromFile(filename string) {
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

// shuffle shuffles the question order
func (q *Quiz) shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(q.qas), func(i, j int) {
		q.qas[i], q.qas[j] = q.qas[j], q.qas[i]
	})
}

// ReadInput runs a loop that waits for input and blocks.
// To avoid unnecessary blocking, there are 2 conditions:
// Its pre-condition checks that a question is waiting to be answered.
// Its post-condition checks that the answer can be handled.
func (q *Quiz) ReadInput(promptCh <-chan *qa, donePromptCh chan struct{}) {
	var s string
	for {
		promptingQuestion, ok := <-promptCh // wait for a question to prompt for
		if !ok {
			return // No more questions to prompt for.
		}

		fmt.Scanln(&s)
		if q.ops == nil {
			return // Quiz has terminated whilst waiting for input.
		}
		q.ops <- q.QuestionAnswered(promptingQuestion, s, donePromptCh)
	}
}

// QuestionAnswered is an action that stores the given answer
// into the quiz map of answered questions.
// It then sends notification of completion over its done chan.
func (q *Quiz) QuestionAnswered(
	promptingQuestion *qa,
	answer string,
	donePromptCh chan<- struct{}) func() bool {

	return func() bool {
		q.answered[promptingQuestion] = answer
		donePromptCh <- struct{}{} // notify next question can be asked
		return false
	}
}

// AskQuestion is an action that prints the question to stdout.
// It then sends notification that the question is waiting to
// to be answered using the askCh chan.
func (q *Quiz) AskQuestion(question *qa, askCh chan<- *qa) func() bool {
	return func() bool {
		fmt.Printf("%s= ", question.question)
		askCh <- question
		return false
	}
}

// PrintSummary is an action that computes the score and prints
// out a summary.
// The return value of true indicates that this should be the
// last operation in the actor loop.
func (q *Quiz) PrintSummary() func() bool {
	return func() bool {
		var correct int
		for k, v := range q.answered {
			if k.answer == strings.Trim(v, " ") {
				correct++
			}
		}
		fmt.Printf("You got %d/%d correct!\n", correct, len(q.qas))
		return true
	}
}

// Run starts the main actor loop.
// It starts two interacting goroutines used for printing questions
// and waiting for answers via stdin.
func (q *Quiz) Run() {
	promptCh := make(chan *qa)
	donePromptCh := make(chan struct{})

	defer func() {
		close(q.ops)
		close(promptCh)
		close(donePromptCh)
	}()

	go q.ReadInput(promptCh, donePromptCh)

	go func() {
		for _, question := range q.qas {
			q.ops <- q.AskQuestion(question, promptCh)
			if _, ok := <-donePromptCh; !ok { // wait until question has been answered.
				return // timed out waiting for input.
			}
		}
		q.ops <- q.PrintSummary() // all questions have been answered.
	}()

loop:
	for {
		select {
		case op := <-q.ops:
			if done := op(); done {
				break loop
			}
		case <-q.ctx.Done():
			fmt.Println("Times up!")
			q.PrintSummary()() // Bypass the ops channel and execute op directly.
			break loop
		}
	}
}

func main() {
	const (
		defaultCSV     = "./problems.csv"
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

	quiz := NewQuiz(context.Background())
	quiz.LoadFromFile(csvArg)

	if shuffle {
		quiz.shuffle()
	}

	var cancel context.CancelFunc
	if timeoutArg > 0 {
		fmt.Println("Press enter to start...")
		fmt.Fscanf(os.Stdin, "%s")
		var timeoutCtx context.Context
		timeoutCtx, cancel = context.WithTimeout(quiz.ctx, time.Duration(timeoutArg)*time.Second)
		quiz.ctx = timeoutCtx
		defer cancel()
	}

	quiz.Run()
}
