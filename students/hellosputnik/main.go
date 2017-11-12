package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

type Quiz struct {
	problems []Problem
	score    int
}

type Settings struct {
	filename  *string
	timeLimit *int
}

func main() {
	quiz := Quiz{}
	settings := Settings{}

	settings.filename = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	settings.timeLimit = flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// Get the flags (if any).
	flag.Parse()

	// Create a file handle for the file.
	file, err := os.Open(*settings.filename)

	// If there was an error opening the file, exit.
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Create a buffered reader to read the file input.
	fin := bufio.NewScanner(file)

	// Read the problems from the comma-separated values file.
	for fin.Scan() {
		line := strings.Split(fin.Text(), ",")
		problem := Problem{question: line[0], answer: line[1]}

		quiz.problems = append(quiz.problems, problem)
	}

	// Create a timer to enforce a time limit.
	timer := time.NewTimer(time.Second * time.Duration(*settings.timeLimit))
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\nYou scored %d out of %d.", quiz.score, len(quiz.problems))
	}()

	// Quiz the user.
	for i, problem := range quiz.problems {
		fmt.Printf("Problem #%d: %s = ", (i + 1), problem.question)

		var input string
		fmt.Scan(&input)

		if input == problem.answer {
			quiz.score++
		}
	}

	// Print the user's results.
	fmt.Printf("You scored %d out of %d.", quiz.score, len(quiz.problems))
}
