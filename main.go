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

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of question,answer")
	timer := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		a := make(chan string)
		var answer string
		go func() {
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			fmt.Scanf("%s\n", &answer)
			a <- answer
		}()

		timer := time.NewTimer(time.Duration(*timer) * time.Second)
		answeredInTime := make(chan bool)
		go func() {
			<-timer.C
			answeredInTime <- false
		}()
		go func() {
			<-a
			answeredInTime <- true
		}()
		if <-answeredInTime {
			if answer == p.a {
				correct += 1
			}
		} else {
			fmt.Printf("You scored %d out of %d", correct, len(problems))
			return
		}
	}

	fmt.Printf("You scored %d out of %d", correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
