package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Question describes a single question from the quiz
type Question struct {
	description string
	answer      string
}

// Quiz describes a CSV quiz
type Quiz struct {
	questions []Question
}

// Input describes an user response for a question
type Input struct {
	response string
	good     bool
}

// Results represents the whole user submissions to a quiz
type Results struct {
	submissions []Input
}

func (r Results) score() int {
	count := 0
	for _, result := range r.submissions {
		if result.good {
			count++
		}
	}

	return count
}

func main() {
	fileName := flag.String("file", "problems.csv", "the CSV filename")
	timer := flag.Int("timer", 30, "the amount of time expected for the quiz")
	flag.Parse()

	// Ask for confirmation before starting the quiz
	shouldStartQuiz()
	fmt.Printf("You have %d seconds to finalize the quiz!\n", *timer)

	results := Results{}
	inputChan := make(chan Input)

	// Read CSV file
	quiz, err := readCSV(*fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Read user input
OuterLoop:
	for i, question := range quiz.questions {
		go func() {
			input := doQuiz(i, question)
			inputChan <- input
		}()
		select {
		case input := <-inputChan:
			results.submissions = append(results.submissions, input)
		case <-time.After(time.Duration(*timer) * time.Second):
			fmt.Println("Quiz time is over!")
			break OuterLoop
		}
	}

	// Print results
	fmt.Printf(
		"Total Good: %d\nTotal Questions Answered: %d\nTotal Questions on Quiz: %d\n",
		results.score(),
		len(results.submissions),
		len(quiz.questions),
	)
}

func readCSV(fileName string) (*Quiz, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading csv file: %s", err)
	}

	quiz := Quiz{}
	reader := csv.NewReader(strings.NewReader(string(data)))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading CSV row: %s\n", err)
		}

		quiz.questions = append(quiz.questions, Question{
			description: normalizeSpaces(record[0]),
			answer:      normalizeSpaces(record[1]),
		})
	}

	// Randomize quiz questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(quiz.questions),
		func(i, j int) {
			quiz.questions[i], quiz.questions[j] = quiz.questions[j], quiz.questions[i]
		},
	)

	return &quiz, nil
}

func shouldStartQuiz() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Do you want to start the quiz? [Y/n for exiting]...")
		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatal("error reading confirmation")
			continue
		}

		if strings.TrimSpace(scanner.Text()) == "Y" {
			fmt.Println("##### WELCOME TO THE QUIZ! #####")
			break
		}

		if strings.TrimSpace(scanner.Text()) == "n" {
			os.Exit(0)
		}
	}
}

func doQuiz(i int, question Question) Input {
	// Read question i
	scanner := bufio.NewScanner(os.Stdin)
	input := Input{}
	fmt.Printf("Question %d: %s?\n", i+1, question.description)
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatalf("error reading question #%d", i)
		return input
	}

	input.response = normalizeSpaces(scanner.Text())
	input.good = strings.EqualFold(input.response, question.answer)

	return input
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}
