package main

import (
	"bufio"
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

type problem struct {
	question string
	answer   string
}

var (
	problemsFile = flag.String("problems", "problems.csv", "A CSV file containing problems and their solutions")
	quizTime     = flag.Int("time", 30, "The time in seconds this quiz will run")
	shuffle      = flag.Bool("shuffle", false, "Wheteher or not to shuffle the problems")
	osR          = bufio.NewReader(os.Stdin)
)

func readProblems(csvFile *os.File) []problem {
	csvR := csv.NewReader(csvFile)

	problems := make([]problem, 0)
	for {
		record, err := csvR.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] != "" && record[1] != "" {
			problems = append(problems, problem{record[0], strings.ToLower(record[1]})
		}
	}

	return problems
}

func shuffleProblems(problems []problem) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i1 := 0; i1 < len(problems); i1++ {
		i2 := r.Intn(len(problems))

		problem1 := problems[i1]
		problem2 := problems[i2]

		problems[i1] = problem2
		problems[i2] = problem1
	}
}

func askQuestion(q string, tries int8) string {
	fmt.Printf("\n%s: ", q)
	input, err := osR.ReadString('\n')
	if err != nil {
		if tries < 5 {
			log.Println("Your answer could not be processed, please try again")
			return askQuestion(q, tries+1)
		}
		log.Fatal("Something is wrong with this program, going to exit...")
	}
	return strings.ToLower(strings.TrimSpace(strings.TrimRight(input, "\n")))
}

func askQuestions(problems []problem, timer, correctAnswersChan, done chan interface{}) {
	index := 0
	for {
		select {
		case <-timer:
			fmt.Println("\nTime's up!")
			return

		default:
			if index >= len(problems) {
				close(done)
				return
			}

			problem := problems[index]
			answer := askQuestion(problem.question, 0)
			if answer == problem.answer {
				correctAnswersChan <- true
				fmt.Println("Correct!")
			} else {
				fmt.Println("False...")
			}
			index = index + 1
		}
	}
}

func main() {
	flag.Parse()

	if !strings.HasSuffix(*problemsFile, "csv") {
		log.Fatalf("Provided problems file '%s' is not a CSV file", *problemsFile)
	}

	csvFile, err := os.Open(*problemsFile)
	if err != nil {
		log.Fatalf("Could not open '%s'", *problemsFile)
	}

	problems := readProblems(csvFile)
	totalProblems := len(problems)
	if *shuffle {
		shuffleProblems(problems)
	}

	timer := make(chan interface{})
	correctAnswersChan := make(chan interface{})
	done := make(chan interface{})

	fmt.Println("Press ENTER to start the quiz...")
	osR.ReadString('\n')

	correctAnswers := 0
	go func() {
		for {
			select {
			case _ = <-correctAnswersChan:
				correctAnswers = correctAnswers + 1
			case <-done:
				return
			}
		}
	}()

	go askQuestions(problems, timer, correctAnswersChan, done)

	time.Sleep(time.Duration(*quizTime) * time.Second)
	close(timer)
	close(done)

	if totalProblems == correctAnswers {
		fmt.Println("\nCongratulations! You answered all questions correctly!")
	} else {
		fmt.Printf(
			"\nYou answered %d questions correctly but failed to do so for %d questions, try again",
			correctAnswers,
			totalProblems-correctAnswers,
		)
	}

	os.Exit(0)
}
