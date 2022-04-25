package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	question string
	answer string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Float64("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	_ = timeLimit
	
	file, err := os.Open(*csvFileName)
	defer file.Close()
	if  err != nil {
		exit(fmt.Sprintf("Error while opening csv file: %s\n", *csvFileName))
	}

	fileContent := csv.NewReader(file)
	lines, err := fileContent.ReadAll()
	if err != nil {
		exit(err.Error())
	}
	problems := parseLines(lines)
	var correct int
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, problem := range problems {
		question := problem.question
		answer, _ := strconv.Atoi(problem.answer)
		answerChan := make(chan int)

		fmt.Printf("Problem %d: %s = ", i+1, question)

		go func() {
			scanner  := bufio.NewScanner(os.Stdin);
			scanner.Scan();
			input , err := strconv.Atoi(scanner.Text());
			if err != nil {
				fmt.Println()
				return;
			}
			answerChan <- input
		}()

		select {
		case _, ok:= <-timer.C:
			if ok {
				fmt.Printf("\nYou scored %d out of %d.", correct, len(problems))
				return
			}
		case input, ok := <-answerChan:
			if ok && answer == input {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.", correct, len(problems))

}

func parseLines(problems [][]string) []problem {
	ret := make([]problem, len(problems))
	for i, line := range problems {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}