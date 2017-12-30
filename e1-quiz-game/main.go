package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	problemsFilename := flag.String("c", "problems.csv", "The file containing the problems")
	timeLimit := flag.Int("l", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*problemsFilename)
	if err != nil {
		fmt.Printf("Failed to open %s\n", *problemsFilename)
		os.Exit(1)
	}

	csv := csv.NewReader(file)
	lines, err := csv.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read %s\n", *problemsFilename)
		os.Exit(1)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answers := make(chan string)

	correct := 0

	for _, line := range lines {
		question := strings.TrimSpace(line[0])
		answer := strings.TrimSpace(line[1])
		fmt.Printf("%s? ", question)

		go func() {
			var text string
			fmt.Scanf("%s\n", &text)
			answers <- text
		}()

		select {
		case <-timer.C:
			fmt.Printf("Times up!\n")
			fmt.Printf("You got %d correct output of %d\n", correct, len(lines))
			return
		case text := <-answers:
			if text == answer {
				correct++
			}
		}
	}

	fmt.Printf("You got %d correct output of %d\n", correct, len(lines))
}
