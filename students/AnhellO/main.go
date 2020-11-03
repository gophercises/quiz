package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	flag.Parse()
	// Read CSV file
	quiz, err := readCSV(*fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Read user input
	results := doQuiz(quiz)

	// Print results
	fmt.Printf("Total Good: %d\nTotal Questions: %d\n", results.score(), len(results.submissions))
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

		quiz.questions = append(quiz.questions, Question{description: record[0], answer: record[1]})
	}

	return &quiz, nil
}

func doQuiz(quiz *Quiz) *Results {
	scanner := bufio.NewScanner(os.Stdin)
	score := Results{}
	fmt.Println("##### WELCOME TO THE QUIZ! #####")

	for i, question := range quiz.questions {
		// Read question i
		input := Input{response: "", good: false}
		fmt.Printf("Question %d: %s?\n", i, question.description)
		scanner.Scan()
		if scanner.Err() != nil {
			log.Fatalf("error reading question #%d", i)
			continue
		}
		input.response = strings.TrimSpace(scanner.Text())
		input.good = input.response == question.answer
		score.submissions = append(score.submissions, input)
	}

	return &score
}
