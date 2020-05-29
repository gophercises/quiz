// https://github.com/gophercises/quiz
package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// loadQuiz load the questions from the CSV file
func loadQuiz(filename *string) ([][]string, error) {
	file, err := os.Open(*filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

// Display a question and wait for the answer from stdin
func ask(question string) (string, error) {
	fmt.Printf("%s?\n", question)

	scanner := bufio.NewScanner(os.Stdin)
	var s string
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		s = scanner.Text()
		break
	}
	return strings.Trim(s, " "), nil
}

// Display a simple welcome message and wait for the user to start
func intro(t time.Duration, nq int) error {
	fmt.Printf("You will be asked to answer %d questions in %v time.\n", nq, t)
	fmt.Println("Please enter any character to start the quiz...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		scanner.Text()
		break
	}
	return nil
}

// Run the quiz in a given time limit
func runQuiz(timer *time.Timer, records [][]string) (int, error) {
	nbResponses := 0
	for _, qa := range records {
		errCh := make(chan error)
		resCh := make(chan bool)

		go func() {
			q, a := qa[0], qa[1]
			resp, err := ask(q)
			if err != nil {
				errCh <- err
				return
			}
			resCh <- a == resp
		}()

		select {
		case <-timer.C:
			fmt.Println("runQuiz : done")
			return nbResponses, errors.New("Timeout reached!")
		case err := <-errCh:
			return nbResponses, err
		case res := <-resCh:
			if res {
				nbResponses++
			}
		}
	}

	return nbResponses, nil
}

// Utility function to log an error and exit the program
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var filename = flag.String("f", "problems.csv", "The CSV file to load")
	var timeLimit = flag.Duration("t", 10*time.Second, "The timeout of the quiz in seconds")
	flag.Parse()
	// Load the questions
	fmt.Println("Loading questions from file : ", *filename)
	records, err := loadQuiz(filename)
	fatal(err)
	// Display a welcome message to the player and wait for him to start
	err = intro(*timeLimit, len(records))
	fatal(err)
	//Process the Q/A
	timer := time.NewTimer(*timeLimit)
	nbResponses, err := runQuiz(timer, records)
	// Display the results
	fmt.Printf("Nb of correct answers : %d/%d\n", nbResponses, len(records))
	fatal(err)
}
