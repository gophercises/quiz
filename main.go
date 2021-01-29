package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
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

func main() {
	var problemFile = flag.String("problems-file", "problems.csv", "The problem csv file path.")
	var timeVar = flag.Duration("time", 30*time.Second, "Timer")
	var shuffle = flag.Bool("no-shuffle", true, "To disable shuffle of problems")
	flag.Parse()

	quiz, err := parseCSV(*problemFile)
	if err != nil {
		log.Fatal(err)
	}

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(quiz), func(i, j int) { quiz[i], quiz[j] = quiz[j], quiz[i] })
	}

	fmt.Printf("You have %v.  Good Luck!! Hit Enter to start\n", *timeVar)
	fmt.Scanln()
	timer := time.NewTimer(*timeVar)

	correct := takeQuiz(quiz, timer)

	fmt.Printf("\nTotal Questions: %d  Correct: %d\n\n", len(quiz), correct)
}

func parseCSV(csvFile string) ([]problem, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return []problem{}, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	problemsCSV, err := r.ReadAll()
	if err != nil {
		return []problem{}, err
	}

	return parseProblems(problemsCSV)
}

func parseProblems(csv [][]string) ([]problem, error) {
	if len(csv) == 0 {
		return []problem{}, errors.New("no problems in csv file")
	}

	problems := make([]problem, len(csv))
	for i, row := range csv {
		// make sure we have two columns in the row
		if len(row) != 2 {
			return []problem{}, fmt.Errorf("invalid csv structure on line %d", i+1)
		}
		problems[i] = problem{
			question: row[0],
			answer:   row[1],
		}
	}
	return problems, nil
}

func takeQuiz(quiz []problem, timer *time.Timer) int {
	var correct int
	for i, problem := range quiz {
		fmt.Printf("Question #%d: %s = ", i+1, problem.question)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case answer := <-answerCh:
			if strings.ToLower(strings.TrimSpace(problem.answer)) == strings.ToLower(strings.TrimSpace(answer)) {
				correct++
			}

		case <-timer.C:
			fmt.Println()
			fmt.Println("Time is up!!!")
			return correct
		}
	}
	return correct
}


