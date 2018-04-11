package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"
	"path/filepath"
)

func main() {
	// Define program arguments (flags)
	var filePath string
	var random bool
	var timeout time.Duration

	flag.StringVar(&filePath, "file", "problems.csv", "Path to CSV file with problem definitions")
	flag.BoolVar(&random, "random", false, "Randomizes the questions")
	flag.DurationVar(&timeout, "timer", 30 * time.Second, "Time limit for answering a question")
	flag.Parse()

	// Get the questions from csv file
	questions, err := readQuestions(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Randomize questions if requested
	if random {
		questions = randomize(questions)
	}

	// Ask the questions
	correct, pct := ask(questions, timeout)
	fmt.Printf("\nRight answers: %d of %d total (%.f%%)\n", correct, len(questions), pct)
}

// readQuestions double checks the file path by making it absolute
// then opens a file reader and passes it along to the CSV reader
// The CSV reader reads the entire file and
func readQuestions(filePath string) ([][]string, error) {
	ff, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(ff)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	r.Comment = '#'
	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ask all questions from the questions set, and set a timer
// return number of correctly answered questions, and the calculated percentage
func ask(questions [][]string, timeout time.Duration) (int, float32) {
	fmt.Printf("You have %v to answer a question.\n", timeout)

	// For reading user input
	reader := bufio.NewReader(os.Stdin)
	correct := 0

	// Start the timer
	timer := time.NewTimer(timeout)

	// Go ask the questions in separate goroutine
	go func() {
		// Loop through the questions
		for _, q := range questions {
			fmt.Printf("%s? : ", q[0])
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if strings.EqualFold(input, q[1]) {
				correct++
			}
			// Reset the timer for next question
			timer.Stop()
			timer.Reset(timeout)
		}
		// Reset the timer to 0 to complete
		timer.Stop()
		timer.Reset(0)
	}()
	// Block and wait for timer
	<- timer.C
	pct := float32(correct) / float32(len(questions)) * 100
	return correct, pct
}

// Take the questions set, and randomize it using a permutation for randomized index
func randomize(src [][]string) [][]string {
	dest := make([][]string, len(src))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}