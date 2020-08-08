package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func readCSVFile(name string) ([][]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return problems
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	fileName := flag.String("f", "problems.csv", "file name")
	timeLimit := flag.Int("t", 30, "time limit")
	flag.Parse()

	lines, err := readCSVFile(*fileName)
	if err != nil {
		return err
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	numberOfCorrectAnswers := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored: %d/%d\n", numberOfCorrectAnswers, len(problems))
			return nil
		case answer := <-answerCh:
			if answer == p.a {
				numberOfCorrectAnswers++
			}
		}
	}

	fmt.Printf("You scored: %d/%d\n", numberOfCorrectAnswers, len(problems))

	return nil
}
